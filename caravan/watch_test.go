package caravan

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/rjeczalik/notify"
	"github.com/stretchr/testify/assert"
)

func TestWatchFile(t *testing.T) {
	conf := Conf{
		Source:      "../fixtures/testsrc",
		Destination: "../fixtures/testoutput",
	}
	done := make(chan int)
	go Watch(conf, func(ei notify.EventInfo) error {
		defer close(done)
		assert.Contains(t, ei.Path(), "fixtures/testsrc/created_by_watch.test")
		return errTestBreak
	})
	time.Sleep(100 * time.Millisecond)
	fn := fmt.Sprintf("%s/created_by_watch.test", conf.Source)
	os.Remove(fn)
	os.Create(fn)
	<-done
}

func TestIsDir(t *testing.T) {
	assert.True(t, isDir("."))
	assert.False(t, isDir("./aaa"))
	assert.False(t, isDir("watch.go"))
}

func TestGetFileName(t *testing.T) {
	assert.Equal(t, "fn", getFileName("/path/to/fn"))
	assert.Equal(t, "fn", getFileName("to/../fn"))
	assert.Equal(t, "fn", getFileName("../fn"))
	assert.Equal(t, "fn", getFileName("to/fn"))
	assert.Equal(t, "fn", getFileName("fn"))
	assert.Equal(t, "", getFileName(""))
}
