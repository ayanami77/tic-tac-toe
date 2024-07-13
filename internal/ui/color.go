package ui

import "github.com/fatih/color"

func BluePrintf(format string, a ...interface{}) {
	blue := color.New(color.FgBlue, color.Bold)
	blue.Printf(format, a...)
}

func RedPrintf(format string, a ...interface{}) {
	red := color.New(color.FgRed, color.Bold)
	red.Printf(format, a...)
}

func GreenPrintf(format string, a ...interface{}) {
	green := color.New(color.FgGreen, color.Bold)
	green.Printf(format, a...)
}
