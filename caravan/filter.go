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
	if v, ok := f.cache[path]; ok {
		return v, err
	}
	for _, pattern := range f.patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return false, err
		}
		v := r.MatchString(path)
		if v {
			f.cache[path] = true
			return true, err
		}
	}
	f.cache[path] = false
	return false, err
}
