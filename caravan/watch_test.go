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
	os.Remove(fn)
}

func TestWatchFileWithArgs(t *testing.T) {
	conf := Conf{
		Source:      "../fixtures/testsrc",
		Destination: "../fixtures/testoutput",
		ExtraArgs:   []string{"--exclude=created_by_watch1"},
	}
	fn1 := fmt.Sprintf("%s/created_by_watch1.test", conf.Source)
	os.Remove(fn1)
	fn2 := fmt.Sprintf("%s/created_by_watch2.test", conf.Source)
	os.Remove(fn2)

	done := make(chan int)
	go Watch(conf, func(ei notify.EventInfo) error {
		defer close(done)
		assert.NotContains(t, ei.Path(), "fixtures/testsrc/created_by_watch1.test")
		assert.Contains(t, ei.Path(), "fixtures/testsrc/created_by_watch2.test")
		return errTestBreak
	})
	time.Sleep(100 * time.Millisecond)

	os.Create(fn1)
	os.Create(fn2)
	<-done
	os.Remove(fn1)
	os.Remove(fn2)
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
