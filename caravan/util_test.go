package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowConf(t *testing.T) {
	output := captureOutput(func() { ShowConf(nil) })
	assert.Empty(t, output)
	output = captureOutput(func() { ShowConf(&DefaultConf) })
	assert.Contains(t, "debug", output)
}
