// Test for checkMethodExist
package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"

	calc "github.com/zahaar/grcp-calculator/gen"
)

func TestCheckMethodExist(t *testing.T) {

	value := "add"
	actual, expected := checkMethodExist(&value), calc.MathMethod_ADD
	if actual != expected {
		t.Errorf("Should return {ADD} Math Method ")
	}

	value = "Add" // Testing against UpperCase
	actual, expected = checkMethodExist(&value), calc.MathMethod_ADD
	if actual != expected {
		t.Errorf("Should return {ADD} Math Method ")
	}

}

// Test that checkMethodExist crashes in case incorrect method is supplied
// ref: https://blog.antoine-augusti.fr/2015/12/testing-an-os-exit-scenario-in-golang/
func TestCheckMethodExistCrashes(t *testing.T) {
	// Only run the failing part when a specific env variable is set
	if os.Getenv("BE_CRASHER") == "1" {
		value := "grail"
		checkMethodExist(&value)
		return
	}

	// Start the actual test in a different subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestCheckMethodExistCrashes")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	stdout, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Check that the log fatal message is what we expected
	gotBytes, _ := ioutil.ReadAll(stdout)
	got := string(gotBytes)
	expected := "16:17:45 Method: {grail}, is not present in: add, sub, mul, div"
	if !strings.HasSuffix(got[8:len(got)-1], expected[8:]) {
		t.Fatalf("Unexpected log message. Got %s but should contain %s", got[8:len(got)-1], expected[8:])
	}

	// Check that the program exited
	err := cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1", err)
	}
}
