package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventWithType(t *testing.T) {
	e := NewEmptyEvent(HookOnError)
	assert.Equal(t, e.EventType, HookOnError)
}
