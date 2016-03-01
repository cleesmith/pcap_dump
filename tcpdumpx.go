package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var err error
	var cmd *exec.Cmd
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var cmdArgs []string

	// **********************************************
	// remember the exec.Command executes as though
	// it did a "cd /" (root dir) and then runs
	// **********************************************

	// cmdName := strings.TrimSpace("pwd")
	// cmdArgs = append(cmdArgs, strings.TrimSpace(`-L`))

	// tcpdump -n -tttt -r eventid2.pcap
	cmdName := strings.TrimSpace("tcpdump")
	cmdArgs = append(cmdArgs, strings.TrimSpace("-n"))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-tttt`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`-r`))
	cmdArgs = append(cmdArgs, strings.TrimSpace(`/Users/chrissmith/go/src/github.com/cleesmith/pcap_dump/eventid2.pcap`))
	// note: it does NOT work to put all args into a single string!

	fmt.Printf("Executing command:\n\t'%v' with %d args: '%s'\n", cmdName, len(cmdArgs), cmdArgs)

	// https://gobyexample.com/variadic-functions
	cmd = exec.Command(cmdName, cmdArgs...)
	// cmd = exec.Command(cmdName) // no args
	// fmt.Printf("cmd=%T=%#v\n", cmd, cmd)
	fmt.Printf("\tworking dir: '%v'\n", cmd.Dir)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error: cmd.Start: err:\n%v\n", err)
		fmt.Printf("stdout:\n%v\n", stdout.String())
		fmt.Printf("stderr:\n%v\n", stderr.String())
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: cmd.Wait: err:\n%v\n", err)
		fmt.Printf("stdout:\n%v\n", stdout.String())
		fmt.Printf("stderr:\n%v\n", stderr.String())
		return
	}
	fmt.Printf("stdout:\n%v\n", stdout.String())
	fmt.Printf("stderr:\n%v\n", stderr.String())
}
