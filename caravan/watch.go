package caravan

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/rjeczalik/notify"
)

// HandleFunc of notify
type HandleFunc func(notify.EventInfo) error

var errTestBreak = errors.New("test break")

// Watch a path
func Watch(conf Conf, handleFunc HandleFunc) {
	c := make(chan notify.EventInfo, 1)
	var logger *LogFormat
	if len(conf.LogFormat) > 0 {
		logger = NewLogger(&conf.LogFormat)
	} else {
		logger = NewLogger(nil)
	}

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
		PrintNotice(logger.Log(ei.Event().String(), ei.Path(), getFileName(ei.Path())))
		err = handleFunc(ei)
		if err != nil {
			PrintError("Handle event error:", err)
		}
		if err == errTestBreak {
			break
		}
	}
}

func isDir(realPath string) bool {
	fi, err := os.Stat(realPath)
	if err != nil {
		PrintError(err)
		return false
	}
	if mode := fi.Mode(); mode.IsDir() {
		return true
	}
	return false
}

func getFileName(path string) string {
	items := strings.Split(path, string(os.PathSeparator))
	if len(items) <= 1 {
		return path
	}
	return items[len(items)-1]
}
