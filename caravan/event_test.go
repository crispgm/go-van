package caravan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventWithType(t *testing.T) {
	e := NewEmptyEvent(HookOnError)
	assert.Equal(t, e.EventType, HookOnError)
}

func TestNewEvent(t *testing.T) {
	e := NewEvent(HookOnChange, "", "/path/to/some/place", GetFileName("/path/to/some/place"))
	assert.Equal(t, HookOnChange, e.EventType)
	assert.Equal(t, "", e.Event)
	assert.Equal(t, "/path/to/some/place", e.Path)
	assert.Equal(t, "place", e.Filename)
}

func TestRunCommands(t *testing.T) {
	_, err := runCommands([]string{})
	assert.EqualError(t, ErrNoCommand, err.Error())
}
