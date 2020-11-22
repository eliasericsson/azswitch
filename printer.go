package main

import (
	"fmt"
	"io"
	"github.com/fatih/color"
)

var (
	ActiveItemColor = color.New(color.FgGreen, color.Bold)
	SuccessColor = color.New(color.FgGreen)
	NonselectColor = color.New(color.FgWhite)
	WarningColor = color.New(color.FgYellow, color.Bold)
	ErrorColor = color.New(color.FgRed, color.Bold)
)

func init()  {
	SuccessColor.EnableColor()
	NonselectColor.EnableColor()
	WarningColor.EnableColor()
	ErrorColor.EnableColor()
}

func Success(w io.Writer, format string, args ...interface{}) error {
	_, err := fmt.Fprintf(w, SuccessColor.Sprint("✔ ")+fmt.Sprintf(format+"\n", args...))
	return err
}

func Nonselect(w io.Writer, format string, args ...interface{}) error {
	_, err := fmt.Fprintf(w, NonselectColor.Sprint("🗴 ")+fmt.Sprintf(format+"\n", args...))
	return err
}

func Warning(w io.Writer, format string, args ...interface{}) error {
	_, err := fmt.Fprintf(w, WarningColor.Sprint("warning: ")+format+"\n", args...)
	return err
}
	
func Error(w io.Writer, format string, args ...interface{}) error {
	_, err := fmt.Fprintf(w, ErrorColor.Sprint("error: ")+format+"\n", args...)
	return err
}