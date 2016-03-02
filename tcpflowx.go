package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	// "time"
)

func main() {
	var err error
	var cmd *exec.Cmd
	var cmdArgs []string

	// tcpflow -c -g -FT -X /dev/null -r eventid2.pcap
	cmdName := strings.TrimSpace("tcpflow")
	// note: it does NOT work to put all args into a single string!
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-c`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-g`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-FT`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-X`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`/dev/null`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-r`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`/Users/chrissmith/go/src/github.com/cleesmith/pcap_dump/eventid2.pcap`))
	// this works too, coz working dir is the dir we execute in:
	// cmdArgs = append(cmdArgs, strings.TrimSpace(`eventid2.pcap`))

	fmt.Printf("Executing command:\n\t'%v' with %d args: '%s'\n", cmdName, len(cmdArgs), cmdArgs)

	// https://gobyexample.com/variadic-functions
	cmd = exec.Command(cmdName, cmdArgs...)
	// fmt.Printf("cmd=%T=%#v\n", cmd, cmd)
	// fmt.Printf("\tworking dir: '%v'\n", cmd.Dir)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: cmd.Start: err:\n%v\n", err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error: cmd.Start: err:\n%v\n", err)
		return
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdoutPipe)
	for in.Scan() {
		fmt.Printf("line=%#v\n", in.Text())
	}
	if err := in.Err(); err != nil {
		fmt.Printf("Error: err:\n%v\n", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: cmd.Wait: err:\n%v\n", err)
		return
	}

}
