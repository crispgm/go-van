package van

import "github.com/fatih/color"

// PrintSuccess print green text
var PrintSuccess = color.New(color.Bold, color.FgGreen).PrintlnFunc()

// PrintWarning print yellow text
var PrintWarning = color.New(color.Bold, color.FgYellow).PrintlnFunc()

// PrintError print red text
var PrintError = color.New(color.Bold, color.FgRed).PrintlnFunc()

// PrintNotice print normal text
var PrintNotice = color.New(color.FgWhite).PrintlnFunc()
