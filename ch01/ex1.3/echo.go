package main

import (
    "os"
    "strings"
)

func Echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func Echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
}

func Echo3(args []string) {
    strings.Join(args[1:], " ")
}

func main() {
    Echo1(os.Args)
    Echo2(os.Args)
    Echo3(os.Args)
}

