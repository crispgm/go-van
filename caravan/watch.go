package caravan

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rjeczalik/notify"
)

// HandleFunc of notify
type HandleFunc func(notify.EventInfo) error

var errTestBreak = errors.New("test break")

// Watch a path
func Watch(conf Conf, handleFunc HandleFunc) {
	c := make(chan notify.EventInfo, 1)

	for {
		realPath, err := filepath.Abs(conf.Source)
		if err != nil {
			panic(err)
		}
		if isDir(realPath) {
			realPath += "/..."
		}
		if err := notify.Watch(realPath, c, notify.All); err != nil {
			panic(err)
		}
		defer notify.Stop(c)

		// Block until an event is received.
		ei := <-c
		PrintNotice(getTime(), "Event", ei.Event().String, ei.Path())
		err = handleFunc(ei)
		if err != nil {
			PrintError("Handle event error:", err)
		}
		if err == errTestBreak {
			break
		}
	}
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

func isDir(realPath string) bool {
	fi, err := os.Stat(realPath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if mode := fi.Mode(); mode.IsDir() {
		return true
	}
	return false
}
