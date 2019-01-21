package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWarningSound(t *testing.T) {
	output := captureOutput(func() { WarningSound() })
	assert.Equal(t, "\a", output)
}
