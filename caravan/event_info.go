package caravan

import (
	"github.com/rjeczalik/notify"
)

// DeployOnceEI implements notify.EventInfo interface
// which is passed to events of DeployOnce
type DeployOnceEI struct {
	SourcePath string
}

// Event returns event
func (do DeployOnceEI) Event() notify.Event {
	return notify.Event(0)
}

// Path returns event path
func (do DeployOnceEI) Path() string {
	return do.SourcePath
}

// Sys return sys as nil
func (do DeployOnceEI) Sys() interface{} {
	return nil
}
