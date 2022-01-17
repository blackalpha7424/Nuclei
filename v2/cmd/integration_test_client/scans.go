package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	client "github.com/projectdiscovery/nuclei/v2/pkg/web/api/client"
	orderedmap "github.com/wk8/go-ordered-map"
)

var scansTestCases *orderedmap.OrderedMap

func init() {
	scansTestCases = orderedmap.New()
	scansTestCases.Set("add", AddScan{})
	scansTestCases.Set("get", GetScan{})
	//scansTestCases.Set("execute", ExecuteScan{})
	scansTestCases.Set("progress", ProgressScan{})
	scansTestCases.Set("matches", MatchesScan{})
	scansTestCases.Set("update", UpdateScan{})
	scansTestCases.Set("errors", ErrorsScan{})
	scansTestCases.Set("delete", DeleteScan{})

}

type GetScan struct{}

func (t GetScan) Execute(cmd exec.Cmd) error {

	result, err := RunNucleiClientWithResult(cmd, "scans", "get")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type AddScan struct{}

func (t AddScan) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "scans", "add", "-name", "integration-test-scan", "-run", "-targets", "example.com", "-templates", "workflows/zimbra-workflow.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type ProgressScan struct{}

func (t ProgressScan) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "scans", "progress")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type UpdateScan struct{}

func (t UpdateScan) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "scans", "get", "-search", "integration-test-scan")
	if err != nil {
		return err
	}
	items := []client.GetScansResponse{}
	json.Unmarshal([]byte(result), &items)
	if err := NotEmptyArray(len(items)); err != nil {
		return err
	}
	for _, item := range items {
		result, err = RunNucleiClientWithResult(cmd, "scans", "update", "-id", fmt.Sprint(item.ID), "-stop")
		if err != nil {
			return err
		}
		if len(result) == 0 {
			return errors.New("empty data return from server")
		}

	}
	return nil
}

type DeleteScan struct{}

func (t DeleteScan) Execute(cmd exec.Cmd) error {
	if result, err := RunNucleiClientWithResult(cmd, "scans", "get", "-search", "integration-test-scan"); err == nil {
		items := []client.GetScansResponse{}
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
	}
	return nil
}

type ExecuteScan struct{}

func (t ExecuteScan) Execute(cmd exec.Cmd) error {
	if result, err := RunNucleiClientWithResult(cmd, "scans", "get", "-search", "integration-test-scan"); err == nil {
		items := []client.GetScansResponse{}
		json.Unmarshal([]byte(result), &items)
		if err := NotEmptyArray(len(items)); err != nil {
			return err
		}
		for _, item := range items {
			result, err := RunNucleiClientWithResult(cmd, "scans", "execute", "-id", fmt.Sprint(item.ID))
			if err != nil {
				return err
			}
			if len(result) == 0 {
				return errors.New("empty data return from server")
			}

		}
	}
	return nil
}

type MatchesScan struct{}

func (t MatchesScan) Execute(cmd exec.Cmd) error {
	if result, err := RunNucleiClientWithResult(cmd, "scans", "get", "-search", "integration-test-scan"); err == nil {
		items := []client.GetScansResponse{}
		json.Unmarshal([]byte(result), &items)
		if err := NotEmptyArray(len(items)); err != nil {
			return err
		}
		for _, item := range items {
			result, err := RunNucleiClientWithResult(cmd, "scans", "matches", "-id", fmt.Sprint(item.ID))
			if err != nil {
				return err
			}
			if len(result) == 0 {
				return errors.New("empty data return from server")
			}
		}
	}
	return nil
}

type ErrorsScan struct{}

func (t ErrorsScan) Execute(cmd exec.Cmd) error {
	if result, err := RunNucleiClientWithResult(cmd, "scans", "get", "-search", "integration-test-scan"); err == nil {
		items := []client.GetScansResponse{}
		json.Unmarshal([]byte(result), &items)
		if err := NotEmptyArray(len(items)); err != nil {
			return err
		}
		for _, item := range items {
			result, err := RunNucleiClientWithResult(cmd, "scans", "errors", "-id", fmt.Sprint(item.ID))
			if err != nil {
				return err
			}
			if len(result) == 0 {
				return errors.New("empty data return from server")
			}
		}
	}
	return nil
}
