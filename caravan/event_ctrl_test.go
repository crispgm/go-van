package caravan_test

import (
	"testing"
	"time"

	"github.com/crispgm/go-van/caravan"
	"github.com/stretchr/testify/assert"
)

func TestFireEvent(t *testing.T) {
	conf := &caravan.Conf{
		Debug:  true,
		OnInit: []string{"ls"},
	}
	output := caravan.CaptureOutput(func() {
		eventCtrl := caravan.NewEventCtrl(conf)
		eventCtrl.EventLoop()
		eventCtrl.FireEvent(caravan.NewEmptyEvent(caravan.HookOnInit))
		time.Sleep(1000 * time.Millisecond)
	})
	assert.Contains(t, output, "SYSTEM Handling event")
}
