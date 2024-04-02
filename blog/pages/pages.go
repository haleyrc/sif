package pages

import (
	"slices"
	"sort"

	"golang.org/x/exp/maps"

	"github.com/haleyrc/sif/blog"
)

func keys(m map[string][]blog.Post, reverse bool) []string {
	keys := maps.Keys(m)
	sort.Strings(keys)
	if reverse {
		slices.Reverse(keys)
	}
	return keys
}
