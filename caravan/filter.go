package caravan

import (
	"regexp"
)

// Filter for exclusion
type Filter struct {
	patterns []string
	cache    map[string]bool
}

// NewFilter creates a filter
func NewFilter(patterns []string) *Filter {
	return &Filter{
		patterns: patterns,
		cache:    make(map[string]bool, 0),
	}
}

// Exclude ...
func (f *Filter) Exclude(path string) (bool, error) {
	var err error
	if _, ok := f.cache[path]; ok {
		return true, err
	}
	for _, pattern := range f.patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}
		if r.MatchString(path) {
			return true, err
		}
	}
	return false, err
}
