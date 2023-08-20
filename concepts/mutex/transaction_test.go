package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Transaction(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	transaction()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$ 400400") {
		t.Error("Expected $ 400400,but got something else")
	}

}
