package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/projectdiscovery/nuclei/v2/pkg/testutils"
)

func getTemplatePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "nuclei-templates")
}

var templatesPathTestCases = map[string]testutils.TestCase{
	//cwd
	"./dns/cname-fingerprint.yaml": &cwdTemplateTest{},
	//relative path
	"dns/cname-fingerprint.yaml": &relativePathTemplateTest{},
	//absolute path
	fmt.Sprintf("%v/dns/nameserver-fingerprint.yaml", getTemplatePath()): &absolutePathTemplateTest{},
}

type cwdTemplateTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *cwdTemplateTest) Execute(filePath string) error {
	var routerErr error

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug)
	if err != nil {
		return err
	}
	if routerErr != nil {
		return routerErr
	}
	return expectResultsCount(results, 1)
}

type relativePathTemplateTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *relativePathTemplateTest) Execute(filePath string) error {
	var routerErr error

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug)
	if err != nil {
		return err
	}
	if routerErr != nil {
		return routerErr
	}
	return expectResultsCount(results, 1)
}

type absolutePathTemplateTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *absolutePathTemplateTest) Execute(filePath string) error {
	var routerErr error

	if _, err := os.Stat(filePath); err == nil {
		// path/to/whatever exists
		fmt.Println("------------EXIST-------------")

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		fmt.Println("------------NO EXIST-------------")

	}
	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "hackerone.com", debug)
	if err != nil {
		return err
	}
	if routerErr != nil {
		return routerErr
	}
	return expectResultsCount(results, 1)
}
