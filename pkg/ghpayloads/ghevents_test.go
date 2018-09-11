package ghpayloads

import (
	"io/ioutil"
	"testing"
)

// https://hackernoon.com/advanced-testing-in-go-tutorial-28b89d3a813
// https://blog.alexellis.io/golang-writing-unit-tests/

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/test.json")
	if err != nil {
		t.Fatal("could not open test data file")
	}
	// I don't like this since it does not render each assetion in `go test` output
	if string(data) != "testx" {
		t.Fail()
	}
	if string(data) != "testx" {
		t.Errorf("contents of file where not as expected: %s", data)
	}
}