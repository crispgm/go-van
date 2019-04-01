package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogFormat(t *testing.T) {
	fmt := "%t %T %f %e %p"
	lf := NewLogger(&fmt)
	assert.Contains(t, lf.Log("eventXXX", "/path/to/file", "file"), "file eventXXX /path/to/file")
}
