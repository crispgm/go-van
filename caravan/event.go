package caravan

import (
	"errors"
	"os/exec"
	"strings"
)

// EventType of a event
type EventType int8

// Event hooks
const (
	HookOnInit EventType = iota
	HookOnChange
	HookOnDeploy
	HookOnError
)

// Event body
type Event struct {
	EventType EventType
	Event     string
	Path      string
	Filename  string
}

// EventCtrl for event scheduling
type EventCtrl struct {
	conf  Conf
	event chan Event
}

// ErrNoCommand ...
var ErrNoCommand = errors.New("No command attached")

// NewEventCtrl creates an EventCtrl
func NewEventCtrl(conf *Conf) *EventCtrl {
	return &EventCtrl{
		conf:  *conf,
		event: make(chan Event),
	}
}

// NewEmptyEvent creates an empty event with only event type
func NewEmptyEvent(et EventType) *Event {
	return &Event{
		EventType: et,
	}
}

// NewEvent creates an event
func NewEvent(et EventType, ev string, path string, filename string) *Event {
	return &Event{
		EventType: et,
		Event:     ev,
		Path:      path,
		Filename:  filename,
	}
}

// FireEvent fires a event
func (ec EventCtrl) FireEvent(event *Event) {
	go func() {
		ec.event <- *event
	}()
}

// EventLoop starts event loop
func (ec *EventCtrl) EventLoop() {
	go func() {
		if ec.conf.Debug {
			PrintNotice("SYSTEM Start event loop")
		}
		for {
			select {
			case e := <-ec.event:
				if ec.conf.Debug {
					PrintNotice("SYSTEM Handling event:", getHookName(e.EventType), "on", e.Filename)
				}
				outputs, err := ec.runEventHook(e)
				if err == ErrNoCommand {
					return
				}
				if err != nil {
					PrintError("SYSTEM Run hook error", err)
				} else {
					PrintSuccess("SYSTEM Invoke", getHookName(e.EventType), "-> output:", strings.Trim(strings.Join(outputs, "\n"), "\n\t "))
				}
			}
		}
	}()
}

func (ec EventCtrl) runEventHook(event Event) ([]string, error) {
	var outputs []string
	var err error
	switch event.EventType {
	case HookOnInit:
		outputs, err = runCommands(ec.conf.OnInit)
	case HookOnChange:
		outputs, err = runCommands(ec.conf.OnChange)
	case HookOnDeploy:
		outputs, err = runCommands(ec.conf.OnDeploy)
	case HookOnError:
		outputs, err = runCommands(ec.conf.OnError)
	default:
		PrintError("Illegal event type", event.EventType)
	}
	return outputs, err
}

func runCommands(commands []string) ([]string, error) {
	if len(commands) == 0 {
		return nil, ErrNoCommand
	}
	var outputs []string
	for _, command := range commands {
		realCommand := strings.Split(command, " ")
		var cmd *exec.Cmd
		if len(realCommand) == 1 {
			cmd = exec.Command(realCommand[0])
		} else {
			cmd = exec.Command(realCommand[0], realCommand[1:]...)
		}
		output, err := cmd.CombinedOutput()
		outputs = append(outputs, string(output))
		if err != nil {
			return outputs, err
		}
	}
	return outputs, nil
}

func getHookName(et EventType) string {
	var hookName string
	switch et {
	case HookOnInit:
		hookName = "OnInit"
	case HookOnChange:
		hookName = "OnChange"
	case HookOnDeploy:
		hookName = "OnDeploy"
	case HookOnError:
		hookName = "OnError"
	default:
		hookName = ""
	}
	return hookName
}
