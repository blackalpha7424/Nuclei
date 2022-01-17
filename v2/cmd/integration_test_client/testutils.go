package main

import (
	"fmt"
	"os/exec"
)

type TestCase interface {
	Execute(exec.Cmd) error
}

func RunNucleiClientWithResult(cmd exec.Cmd, extra ...string) (string, error) {
	cmd.Args = append(cmd.Args, extra...)
	fmt.Println(cmd.String())
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
