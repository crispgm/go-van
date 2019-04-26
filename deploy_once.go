package main

import (
	"github.com/rjeczalik/notify"
)

type deployOnceEI struct {
	SourcePath string
}

func (do deployOnceEI) Event() notify.Event {
	return 0
}

func (do deployOnceEI) Path() string {
	return do.SourcePath
}

func (do deployOnceEI) Sys() interface{} {
	return nil
}
