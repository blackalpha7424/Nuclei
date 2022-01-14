package main

import (
	"errors"
	"os/exec"
	orderedmap "github.com/wk8/go-ordered-map"
)

var templatesTestCases *orderedmap.OrderedMap

func init() {
	templatesTestCases = orderedmap.New()
	templatesTestCases.Set("add", addTemplate{})
	templatesTestCases.Set("get", getTemplate{})
	templatesTestCases.Set("update", updateTemplate{})
	templatesTestCases.Set("raw",rawTemplate{})
	templatesTestCases.Set("execute",executeTemplate{})
	templatesTestCases.Set("delete", deleteTemplate{})

}
type getTemplate struct{}

func (t getTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "get")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type addTemplate struct{}

func (t addTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "add", "-filepath", "./test-data/template.yaml", "-folder", "nuclei-templates", "-path", "test-data/template.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type updateTemplate struct{}

func (t updateTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "update", "-filepath", "./test-data/template.yaml", "-path", "test-data/template.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type deleteTemplate struct{}

func (t deleteTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "delete", "-path", "test-data/template.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type rawTemplate struct{}

func (t rawTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "raw", "-path", "technologies/ibm/ibm-http-server.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}

type executeTemplate struct{}

func (executeTemplate) Execute(cmd exec.Cmd) error {
	result, err := RunNucleiClientWithResult(cmd, "templates", "execute", "-target", "example.com", "-path", "technologies/ibm/ibm-http-server.yaml")
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("empty data return from server")
	}
	return nil
}
