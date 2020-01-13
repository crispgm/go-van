package caravan_test

import (
	"testing"
	"time"

	"github.com/crispgm/go-van/caravan"
	"github.com/stretchr/testify/assert"
)

func TestFireAllEvents(t *testing.T) {
	conf := &caravan.Conf{
		Debug:    true,
		OnInit:   []string{"ls", "echo \"i love van\""},
		OnChange: []string{"echo changed"},
		OnError:  []string{"echo error"},
		OnDeploy: []string{"echo deployed"},
	}
	output := caravan.CaptureOutput(func() {
		eventCtrl := caravan.NewEventCtrl(conf)
		eventCtrl.EventLoop()
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnInit))
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnChange))
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnDeploy))
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnError))
		time.Sleep(1000 * time.Millisecond)
	})
	assert.Contains(t, output, "SYSTEM Handling event: OnInit on")
	assert.Contains(t, output, "SYSTEM Handling event: OnChange on")
	assert.Contains(t, output, "SYSTEM Handling event: OnDeploy on")
	assert.Contains(t, output, "SYSTEM Handling event: OnError on")
}

func TestFirePartialEvents(t *testing.T) {
	conf := &caravan.Conf{
		Debug:    true,
		OnChange: []string{"echo changed"},
		OnDeploy: []string{"echo deployed"},
	}
	output := caravan.CaptureOutput(func() {
		eventCtrl := caravan.NewEventCtrl(conf)
		eventCtrl.EventLoop()
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnInit))
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnDeploy))
		time.Sleep(1000 * time.Millisecond)
	})
	assert.Contains(t, output, "SYSTEM Handling event: OnInit on")
	assert.NotContains(t, output, "SYSTEM Handling event: OnChange on")
	assert.Contains(t, output, "SYSTEM Handling event: OnDeploy on")
	assert.NotContains(t, output, "SYSTEM Handling event: OnError on")
}
