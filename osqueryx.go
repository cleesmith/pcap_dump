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

	cmdName := strings.TrimSpace("osqueryi")
	cmdArgs = append(cmdArgs, strings.TrimSpace("--json"))
	// cmdArgs = append(cmdArgs, strings.TrimSpace(`select * from time;`))
	// cmdArgs = append(cmdArgs, strings.TrimSpace(`select * from etc_protocols where number = 1;`)) // icmp
	// cmdArgs = append(cmdArgs, strings.TrimSpace(`select * from etc_protocols where number = 6;`)) // tcp
	cmdArgs = append(cmdArgs, strings.TrimSpace(`select * from etc_protocols where name = 'tcp';`)) // tcp
	fmt.Printf("Executing command:\n'%v' with args=%T: '%v'\n", cmdName, cmdArgs, cmdArgs)
	// https://gobyexample.com/variadic-functions
	cmd = exec.Command(cmdName, cmdArgs...)
	// fmt.Printf("cmd=%T=%#v\n", cmd, cmd)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		fmt.Printf("cmd.Start: err:\n%v\n", err)
		fmt.Printf("stdout:\n%v\n", stdout.String())
		fmt.Printf("stderr:\n%v\n", stderr.String())
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("cmd.Wait: err:\n%v\n", err)
		fmt.Printf("stdout:\n%v\n", stdout.String())
		fmt.Printf("stderr:\n%v\n", stderr.String())
		return
	}
	fmt.Printf("stdout:\n%v\n", stdout.String())
	fmt.Printf("stderr:\n%v\n", stderr.String())
}
