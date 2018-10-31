package van

import (
	"log"

	"github.com/rjeczalik/notify"
)

// HandleFunc of notify
type HandleFunc func(notify.EventInfo) error

// Watch a path
func Watch(path string, handleFunc HandleFunc) {
	c := make(chan notify.EventInfo, 1)

	for {
		if err := notify.Watch(path, c, notify.All); err != nil {
			log.Fatal(err)
		}
		defer notify.Stop(c)

		// Block until an event is received.
		ei := <-c
		log.Println("Got event:", ei)
		err := handleFunc(ei)
		log.Println("Handle event:", err)
	}
}
