package caravan

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// PrintSuccess print green text
var PrintSuccess = color.New(color.FgGreen).PrintlnFunc()

// PrintWarning print yellow text
var PrintWarning = color.New(color.Bold, color.FgYellow).PrintlnFunc()

// PrintError print red text
var PrintError = color.New(color.Bold, color.FgRed).PrintlnFunc()

// PrintNotice print normal text
var PrintNotice = fmt.Println

// PrintLog print text with time and other info
func PrintLog(text ...interface{}) {
	ts := fmt.Sprintf("[%s]", getTime())
	s := make([]interface{}, len(text)+1)
	s[0] = ts
	for i, t := range text {
		s[i+1] = t
	}
	PrintNotice(s...)
}

// WarningSound warning sound for error case
func WarningSound() {
	fmt.Print("\a")
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
