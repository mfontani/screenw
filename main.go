package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	width, _, err := terminal.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	if width > 999 {
		panic("too wide!")
	}
	var lines [3]string
	for n := 1; n <= width; n++ {
		if n%100 == 0 {
			lines[0] += fmt.Sprintf("%d", n/100%100)
		} else {
			lines[0] += " "
		}
		if n%10 == 0 {
			lines[1] += fmt.Sprintf("%d", n/10%10)
		} else {
			lines[1] += " "
		}
		lines[2] += fmt.Sprintf("%d", n%10)
	}
	_end := fmt.Sprintf("=%d", width)
	_fmtfirst := fmt.Sprintf("%%-%ds%4s", width-4, _end)
	_fmt := fmt.Sprintf("%%-%ds", width-1)
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			lines[i] = fmt.Sprintf(_fmtfirst, strings.TrimRight(lines[i], " "))
		} else {
			lines[i] = fmt.Sprintf(_fmt, lines[i])
		}
		fmt.Printf("%s\n", lines[i])
	}
}
