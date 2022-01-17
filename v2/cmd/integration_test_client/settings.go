package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	client "github.com/projectdiscovery/nuclei/v2/pkg/web/api/client"
	orderedmap "github.com/wk8/go-ordered-map"
)

var settingsTestCases *orderedmap.OrderedMap

func init() {
	settingsTestCases = orderedmap.New()
	settingsTestCases.Set("add", AddSettings{})
	settingsTestCases.Set("get", GetSettings{})
	settingsTestCases.Set("update", UpdateSettings{})

}

type GetSettings struct{}

func (t GetSettings) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "settings", "get")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type AddSettings struct{}

func (t AddSettings) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "settings", "add", "-filepath", "./test-data/template.yaml", "-name", "default", "-type", "internal")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type UpdateSettings struct{}

func (t UpdateSettings) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "settings", "get")
	if err != nil {
		return err
	}
	items := []client.GetSettingsResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err = RunNucleiClientWithResult(cmd, "settings", "update", "-filepath", "./test-data/template.yaml", "-name",fmt.Sprint(item.Name), "-type", "int")
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}

	}
	return nil
}
