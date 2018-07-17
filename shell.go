package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func split_command_input(cmds string)  {
    separators := []rune{' '}
        f := func(r rune) bool {
            for _, cmds := range separators {
	        if r == cmds {
		   return true
		}
	    }
	  return false
	}
	cmd_array := strings.FieldsFunc(cmds,f)
	cmd_array = cmd_array[:len(cmd_array)]
	fmt.Printf("%s", cmd_array)
}


func main() {

    cwd := make([]string, 4096)
    cwd[0] = "%"
    host, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    buffer := bufio.NewReader(os.Stdin)

    for {
	dir, err := os.Getwd()
	if err != nil {
	    panic(err)
	}
	cwd[0] = dir
        fmt.Printf("%s@%s %s %s" , os.Getenv("LOGNAME"), host,  cwd[0], "> ")
        line, _ := buffer.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
        if (len(line) == 1) {
            break
        }
    split_command_input(line)
    }
}
