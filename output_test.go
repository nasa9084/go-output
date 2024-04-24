package output_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nasa9084/go-output"
)

func TestWriter_file(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "go-output-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	writeTo := filepath.Join(tempDir, "testwriter_file")

	flagSet := flag.NewFlagSet("TestWriter_file", 0)
	output.UseFlagSet(flagSet)
	flagSet.Parse([]string{"-o", writeTo})

	w, err := output.Writer()
	if err != nil {
		t.Fatal(err)
	}

	want := "Hello"
	fmt.Fprintf(w, want)

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	b, err := os.ReadFile(writeTo)
	if err != nil {
		t.Fatal(err)
	}

	if got := string(b); got != want {
		t.Fatalf("%s != %s", got, want)
	}
}
