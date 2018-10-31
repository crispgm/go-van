package van

import (
	"fmt"
	"time"

	"github.com/rjeczalik/notify"
)

// HandleFunc of notify
type HandleFunc func(notify.EventInfo) error

// Watch a path
func Watch(conf Conf, handleFunc HandleFunc) {
	c := make(chan notify.EventInfo, 1)

	for {
		if err := notify.Watch(conf.Source, c, notify.All); err != nil {
			panic(err)
		}
		defer notify.Stop(c)

		// Block until an event is received.
		ei := <-c
		PrintNotice(getTime(), "Event", ei.Event().String, ei.Path())
		err := handleFunc(ei)
		if err != nil {
			PrintError("Handle event error:", err)
		}
	}
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
