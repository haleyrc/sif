package blog

import (
	"slices"
	"strings"
	"time"
)

type About struct {
	Content string
}

type Metadata struct {
	Manifest string
}

type Blog struct {
	Metadata Metadata
	Name     string
	Author   string
	Since    string
	About    About
	Menu     Menu
	Posts    Posts
}

type Posts struct {
	All    []Post
	ByTag  map[string][]Post
	ByYear map[string][]Post
	Latest Post
}

func (posts Posts) Get(slug string) (Post, bool) {
	idx := slices.IndexFunc(posts.All, func(p Post) bool {
		return p.Slug == slug
	})
	if idx == -1 {
		return Post{}, false
	}
	return posts.All[idx], true
}

type Link struct {
	Text string
	Path string
}

type Menu struct {
	Links []Link
}

type Post struct {
	Title     string
	Slug      string
	Timestamp time.Time
	Tags      []string
	Content   string
}

func slugify(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ToLower(s)
	return s
}
