package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitCommandLineInput(cmds string) []string {
	separators := []rune{' '}
	f := func(r rune) bool {
		for _, sep := range separators {
			if r == sep {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(cmds, f)
}

func main() {
	cwd := make([]string, 4096)
	cwd[0] = "%"
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd[0], err = os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s@%s %s %s", os.Getenv("LOGNAME"), host, cwd[0], "> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimSuffix(line, "\n")
		if len(line) == 0 {
			break
		}

		args := splitCommandLineInput(line)
		_ = args
		// Do something with the args
	}
}
