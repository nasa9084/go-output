package output_test

import (
	"flag"
	"fmt"

	"github.com/nasa9084/go-output"
)

func ExampleWriter_stdout() {
	// use new flag set for testing
	flagSet := flag.NewFlagSet("ExampleCommand", flag.ContinueOnError)
	output.UseFlagSet(flagSet)

	// parse without argument, then stdout will be used
	flag.Parse()

	w, _ := output.Writer()
	defer w.Close() // you can close this even for stdout

	fmt.Fprintln(w, "Hello")

	// Output:
	// Hello
}
