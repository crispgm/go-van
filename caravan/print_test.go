package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintLog(t *testing.T) {
	output := CaptureOutput(func() { PrintLog("test log") })
	assert.Regexp(t, `^\[[0-9]{2}:[0-9]{2}:[0-9]{2}\] test log\n$`, output)
}

func TestWarningSound(t *testing.T) {
	output := CaptureOutput(func() { WarningSound() })
	assert.Equal(t, "\a", output)
}
