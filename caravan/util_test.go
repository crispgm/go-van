package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowConf(t *testing.T) {
	output := CaptureOutput(func() { ShowConf(nil) })
	assert.Empty(t, output)
	output = CaptureOutput(func() { ShowConf(&DefaultConf) })
	assert.Contains(t, output, "debug")
}
