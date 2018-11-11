package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	assert.True(t, isDir("."))
	assert.False(t, isDir("watch.go"))
}
