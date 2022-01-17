package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	client "github.com/projectdiscovery/nuclei/v2/pkg/web/api/client"
	orderedmap "github.com/wk8/go-ordered-map"
)

var targetsTestCases *orderedmap.OrderedMap

func init() {
	targetsTestCases = orderedmap.New()
	targetsTestCases.Set("add", addTargets{})
	targetsTestCases.Set("get", getTargets{})
	targetsTestCases.Set("update", updateTargets{})
	targetsTestCases.Set("contents", contentsTargets{})
	targetsTestCases.Set("delete", deleteTargets{})

}

type getTargets struct{}

func (t getTargets) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "targets", "get")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type addTargets struct{}

func (t addTargets) Execute(cmd exec.Cmd) error {
	// /result, err := RunNucleiClientWithResult(cmd, "targets", "add", "-filepath", "./test-data/contents.txt", "-name", "nuclei-templates", "-path", "httptarget")
	result, err := RunNucleiClientWithResult(cmd, "targets", "add", "-filepath", "./test-data/contents.txt", "-name", "./test-data/contents.txt", "-path", "./test-data/contents.txt")
	if err != nil {
		return err
	}
	return NotEmpty(result)
}

type updateTargets struct{}

func (t updateTargets) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "targets", "get")
	if err != nil {
		return err
	}
	items := []client.GetTargetsResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err = RunNucleiClientWithResult(cmd, "targets", "update", "-filepath", "./test-data/contents.txt", "-id", fmt.Sprint(item.ID))
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}

	}
	return nil
}

type deleteTargets struct{}

func (t deleteTargets) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "targets", "get")
	if err != nil {
		return err
	}
	items := []client.GetTargetsResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err := RunNucleiClientWithResult(cmd, "targets", "delete", "-id", fmt.Sprint(item.ID))
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}
	}
	return nil
}

type contentsTargets struct{}

func (t contentsTargets) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "targets", "get")
	if err != nil {
		return err
	}
	items := []client.GetTargetsResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err := RunNucleiClientWithResult(cmd, "targets", "contents", "-id", fmt.Sprint(item.ID))
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}
	}
	return nil
}
