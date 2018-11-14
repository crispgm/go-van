package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExclusion(t *testing.T) {
	var res bool
	f := NewFilter([]string{".git", ".svn", "/node_modules"})
	res, _ = f.Exclude(".git")
	assert.Equal(t, true, res)
	res, _ = f.Exclude(".git")
	assert.Equal(t, true, res)
	res, _ = f.Exclude(".git1")
	assert.Equal(t, true, res)
	res, _ = f.Exclude(".g1t")
	assert.Equal(t, false, res)
}

func TestExclusionRegexErr(t *testing.T) {
	f := NewFilter([]string{"+"})
	res, err := f.Exclude(".git")
	assert.Equal(t, false, res)
	assert.Error(t, err)
}
