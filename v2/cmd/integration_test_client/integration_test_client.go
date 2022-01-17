package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/logrusorgru/aurora"
	orderedmap "github.com/wk8/go-ordered-map"
)

var (
	debug     = os.Getenv("DEBUG") == "true"
	success   = aurora.Green("[✓]").String()
	failed    = aurora.Red("[✘]").String()
	serverURL = flag.String("url", "", "Base URL of the Nuclei Server")
	errored   = false
)

func main() {
	flag.Parse()
	tests := orderedmap.New()
    tests.Set("templates", templatesTestCases)
	tests.Set("targets", targetsTestCases)
	tests.Set("scans", scansTestCases)
	tests.Set("settings",  settingsTestCases)
	tests.Set("issues",    issuesTestCases)

	fmt.Println("nuclei server is running on:", *serverURL)
	cmd := exec.Command("./nuclei-client")
	cmd.Args = append(cmd.Args, "-no-color", "-url", *serverURL, "-username", "user", "-password", "pass")
	if debug {
		cmd.Stderr = os.Stderr
	}
	for pair := tests.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("Running test cases for \"%s\" endpoint\n", aurora.Blue(pair.Key))
		child := pair.Value.(*orderedmap.OrderedMap)
		for it := child.Oldest(); it != nil; it = it.Next() {
			fmt.Println(it.Key)
			if err := it.Value.(TestCase).Execute(*cmd); err != nil {
				fmt.Fprintf(os.Stderr, "%s Test \"%s\" failed: %s\n", failed, it.Key, err)
				errored = true
			} else {
				fmt.Printf("%s Test \"%s\" passed!\n", success, it.Key)
			}
		}
	}
	if errored {
		os.Exit(1)
	}
}
