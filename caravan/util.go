package caravan

import (
	"bytes"
	"io"
	"os"
)

// ShowConf prints conf
func ShowConf(conf *Conf) {
	if conf == nil {
		return
	}
	PrintNotice("=>", "debug:", conf.Debug)
	PrintNotice("=>", "once:", conf.Once)
	PrintNotice("=>", "source:", conf.Source)
	PrintNotice("=>", "destination:", conf.Destination)
	PrintNotice("=>", "deploy_mode:", conf.Mode)
	PrintNotice("=>", "incremental:", conf.Incremental)
	PrintNotice("=>", "exclude:", conf.Exclude)
}

// CaptureOutput ...
func CaptureOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
