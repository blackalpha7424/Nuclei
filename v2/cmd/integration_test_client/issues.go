package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	client "github.com/projectdiscovery/nuclei/v2/pkg/web/api/client"
	orderedmap "github.com/wk8/go-ordered-map"
)

var issuesTestCases *orderedmap.OrderedMap

func init() {
	issuesTestCases = orderedmap.New()
	issuesTestCases.Set("get", GetIssues{})
	issuesTestCases.Set("update", UpdateIssues{})
	issuesTestCases.Set("delete", DeleteIssues{})

}

type GetIssues struct{}

func (t GetIssues) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "issues", "get")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type UpdateIssues struct{}

func (t UpdateIssues) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "issues", "get", "-search", "integration-test-issues")
	if err != nil {
		return err
	}
	items := []client.GetIssuesResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err = RunNucleiClientWithResult(cmd, "issues", "update", "-id", fmt.Sprint(item.ID))
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}

	}
	return nil
}

type DeleteIssues struct{}

func (t DeleteIssues) Execute(cmd exec.Cmd) error {
	 result, err := RunNucleiClientWithResult(cmd, "issues", "get", "-search", "integration-test-issues")
	 if err != nil {
		return err
	}
		items := []client.GetIssuesResponse{}
		json.Unmarshal([]byte(result), &items)
		if err := NotEmptyArray(len(items)); err != nil {
			return err
		}
		for _, item := range items {
			result, err := RunNucleiClientWithResult(cmd, "scans", "delete", "-id", fmt.Sprint(item.ID))
			if err != nil {
				return err
			}
			if len(result) == 0 {
				return errors.New("empty data return from server")
			}
		}
	return nil
}
