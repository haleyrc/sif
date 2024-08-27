package blog

import (
	"context"
	"io"

	"github.com/haleyrc/sif/templates/blog/components"
	"github.com/haleyrc/sif/templates/blog/pages"
)

type Template struct {
	siteData pages.SiteData
}

func New(
	author string,
	copyrightStartYear int,
	description, title string,
	menu []pages.MenuItem,
) (*Template, error) {
	t := &Template{
		siteData: pages.SiteData{
			Author:             author,
			CopyrightStartYear: copyrightStartYear,
			Description:        description,
			Menu:               menu,
			Title:              title,
		},
	}
	return t, nil
}

type ArchiveViewModel struct {
	Name        string
	PostsByYear *components.PostIndexArgs
}

func (t *Template) Archive(ctx context.Context, w io.Writer, vm ArchiveViewModel) error {
	return pages.Archive(t.siteData, vm.Name, vm.PostsByYear).Render(ctx, w)
}

type IndexViewModel struct {
	Name            string
	LatestPost      *components.PostArgs
	MostRecentPosts *components.PostListArgs
}

func (t *Template) Index(ctx context.Context, w io.Writer, vm IndexViewModel) error {
	return pages.Index(t.siteData, vm.Name, vm.LatestPost, vm.MostRecentPosts).Render(ctx, w)
}

type PageViewModel struct {
	Page *components.PageArgs
}

func (t *Template) Page(ctx context.Context, w io.Writer, vm PageViewModel) error {
	return pages.Page(t.siteData, vm.Page).Render(ctx, w)
}

type PostViewModel struct {
	Post *components.PostArgs
}

func (t *Template) Post(ctx context.Context, w io.Writer, vm PostViewModel) error {
	return pages.Post(t.siteData, vm.Post).Render(ctx, w)
}

type PostsForTagViewModel struct {
	Posts *components.PostListArgs
	Tag   string
}

func (t *Template) PostsForTag(ctx context.Context, w io.Writer, vm PostsForTagViewModel) error {
	return pages.PostsForTag(t.siteData, vm.Tag, vm.Posts).Render(ctx, w)
}

type PostsForYearViewModel struct {
	Posts *components.PostListArgs
	Year  string
}

func (t *Template) PostsForYear(ctx context.Context, w io.Writer, vm PostsForYearViewModel) error {
	return pages.PostsForYear(t.siteData, vm.Year, vm.Posts).Render(ctx, w)
}

type TagsViewModel struct {
	Name       string
	PostsByTag *components.PostIndexArgs
}

func (t *Template) Tags(ctx context.Context, w io.Writer, vm TagsViewModel) error {
	return pages.Archive(t.siteData, vm.Name, vm.PostsByTag).Render(ctx, w)
}
