package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"

	"github.com/haleyrc/sif/blog"
	"github.com/haleyrc/sif/blog/pages"
)

//go:embed aboutBody.html
var aboutBody string

//go:embed postBody.html
var postBody string

//go:embed site.webmanifest
var manifest []byte

func main() {
	tz, _ := time.LoadLocation("America/New_York")

	b := blog.Blog{
		Metadata: blog.Metadata{
			Manifest: "/assets/site.webmanifest",
		},
		Name:   "Blog",
		Author: "Ryan Haley",
		Since:  "2023",
		About: blog.About{
			Content: aboutBody,
		},
		Menu: blog.Menu{
			Links: []blog.Link{
				{Text: "archive", Path: "/archive"},
				{Text: "tags", Path: "/tags"},
				{Text: "about", Path: "/about"},
			},
		},
		Posts: blog.Posts{
			All: []blog.Post{
				{
					Title:     "Article title",
					Slug:      "20240401-article-title",
					Timestamp: time.Date(2024, 4, 1, 9, 56, 0, 0, tz),
					Tags:      []string{"Go", "React"},
					Content:   postBody,
				},
				{
					Title:     "A recent post",
					Slug:      "20240315-a-recent-post",
					Timestamp: time.Date(2024, 3, 15, 9, 56, 0, 0, tz),
					Tags:      []string{"React"},
				},
				{
					Title:     "Another recent post",
					Slug:      "20240314-another-recent-post",
					Timestamp: time.Date(2024, 3, 14, 9, 56, 0, 0, tz),
					Tags:      []string{"Go"},
				},
				{
					Title:     "A less recent post",
					Slug:      "20240313-a-less-recent-post",
					Timestamp: time.Date(2024, 3, 13, 9, 56, 0, 0, tz),
				},
				{
					Title:     "A really not recent post",
					Slug:      "20240312-a-really-not-recent-post",
					Timestamp: time.Date(2024, 3, 12, 9, 56, 0, 0, tz),
					Tags:      []string{"React"},
				},
				{
					Title:     "A post from a previous year",
					Slug:      "20230312-a-really-not-recent-post",
					Timestamp: time.Date(2023, 3, 12, 9, 56, 0, 0, tz),
				},
			},
			ByTag: map[string][]blog.Post{
				"Go": []blog.Post{
					{
						Title:     "Article title",
						Slug:      "20240401-article-title",
						Timestamp: time.Date(2024, 4, 1, 9, 56, 0, 0, tz),
						Tags:      []string{"Go", "React"},
						Content:   postBody,
					},
					{
						Title:     "Another recent post",
						Slug:      "20240314-another-recent-post",
						Timestamp: time.Date(2024, 3, 14, 9, 56, 0, 0, tz),
						Tags:      []string{"Go"},
					},
				},
				"React": []blog.Post{
					{
						Title:     "Article title",
						Slug:      "20240401-article-title",
						Timestamp: time.Date(2024, 4, 1, 9, 56, 0, 0, tz),
						Tags:      []string{"Go", "React"},
						Content:   postBody,
					},
					{
						Title:     "A recent post",
						Slug:      "20240315-a-recent-post",
						Timestamp: time.Date(2024, 3, 15, 9, 56, 0, 0, tz),
						Tags:      []string{"React"},
					},
					{
						Title:     "A really not recent post",
						Slug:      "20240312-a-really-not-recent-post",
						Timestamp: time.Date(2024, 3, 12, 9, 56, 0, 0, tz),
						Tags:      []string{"React"},
					},
				},
			},
			ByYear: map[string][]blog.Post{
				"2024": {
					{
						Title:     "Article title",
						Slug:      "20240401-article-title",
						Timestamp: time.Date(2024, 4, 1, 9, 56, 0, 0, tz),
						Tags:      []string{"Go", "React"},
						Content:   postBody,
					},
					{
						Title:     "A recent post",
						Slug:      "20240315-a-recent-post",
						Timestamp: time.Date(2024, 3, 15, 9, 56, 0, 0, tz),
						Tags:      []string{"React"},
					},
					{
						Title:     "Another recent post",
						Slug:      "20240314-another-recent-post",
						Timestamp: time.Date(2024, 3, 14, 9, 56, 0, 0, tz),
						Tags:      []string{"Go"},
					},
					{
						Title:     "A less recent post",
						Slug:      "20240313-a-less-recent-post",
						Timestamp: time.Date(2024, 3, 13, 9, 56, 0, 0, tz),
					},
					{
						Title:     "A really not recent post",
						Slug:      "20240312-a-really-not-recent-post",
						Timestamp: time.Date(2024, 3, 12, 9, 56, 0, 0, tz),
						Tags:      []string{"React"},
					},
				},
				"2023": {
					{
						Title:     "A post from a previous year",
						Slug:      "20230312-a-really-not-recent-post",
						Timestamp: time.Date(2023, 3, 12, 9, 56, 0, 0, tz),
					},
				},
			},
			Latest: blog.Post{
				Title:     "Article title",
				Slug:      "20240401-article-title",
				Timestamp: time.Date(2024, 4, 1, 9, 56, 0, 0, tz),
				Tags:      []string{"Go", "React"},
				Content:   postBody,
			},
		},
	}

	http.HandleFunc("/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {
		post, ok := b.Posts.Get(r.PathValue("slug"))
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		templ.Handler(pages.Post(b, post)).ServeHTTP(w, r)
	})

	http.HandleFunc("/tags/{tag}", func(w http.ResponseWriter, r *http.Request) {
		tag := r.PathValue("tag")
		templ.Handler(pages.PostsForTag(b, tag, b.Posts.ByTag[tag])).ServeHTTP(w, r)
	})
	http.Handle("/tags", templ.Handler(pages.Tags(b)))

	http.HandleFunc("/archive/{year}", func(w http.ResponseWriter, r *http.Request) {
		year := r.PathValue("year")
		templ.Handler(pages.PostsForYear(b, year, b.Posts.ByYear[year])).ServeHTTP(w, r)
	})
	http.Handle("/archive", templ.Handler(pages.Archive(b)))

	http.HandleFunc("/assets/site.webmanifest", func(w http.ResponseWriter, r *http.Request) {
		w.Write(manifest)
	})

	http.Handle("/about", templ.Handler(pages.About(b)))
	http.Handle("/", templ.Handler(pages.Index(b)))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
