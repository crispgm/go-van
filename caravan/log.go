package caravan

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// LogFormat handles user-level log
// Default: [%t] Event %e %p
type LogFormat struct {
	format     string
	timestamp  int64
	timeString string
	logBuffer  string
}

var supportedLogFormat = []string{
	"%t", // time in string
	"%T", // timestamp
	"%e", // event address
	"%p", // full path
	"%f", // file name
}

var defaultLogFormat = "[%t] Event %e %p"

// NewLogger create a LogFormat unit
func NewLogger(format *string) *LogFormat {
	if format == nil {
		return &LogFormat{
			format: defaultLogFormat,
		}
	}

	return &LogFormat{
		format: *format,
	}
}

func (lf *LogFormat) setTime() {
	t := time.Now()
	lf.timestamp = t.Unix()
	lf.timeString = fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

// Log input data as string
func (lf LogFormat) Log(eventID, path, filename string) string {
	re := regexp.MustCompile(`(\%t|\%T|\%e|\%p|\%f)+`)
	results := re.FindAllIndex([]byte(lf.format), -1)
	lf.logBuffer = lf.format

	lf.setTime()
	for _, item := range results {
		token := lf.format[item[0]:item[1]]
		switch token {
		case "%t":
			lf.logBuffer = strings.Replace(lf.logBuffer, token, fmt.Sprintf("%d", lf.timestamp), 1)
			break
		case "%T":
			lf.logBuffer = strings.Replace(lf.logBuffer, token, lf.timeString, 1)
			break
		case "%e":
			lf.logBuffer = strings.Replace(lf.logBuffer, token, eventID, 1)
			break
		case "%p":
			lf.logBuffer = strings.Replace(lf.logBuffer, token, path, 1)
			break
		case "%f":
			lf.logBuffer = strings.Replace(lf.logBuffer, token, filename, 1)
			break
		default:
			break
		}
	}
	return lf.logBuffer
}
