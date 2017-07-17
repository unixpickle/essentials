package essentials

import (
	"fmt"
	"os"
)

// Die prints the arguments to standard error in a style
// like the one used by fmt.Println, then exits with an
// error status code.
func Die(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

// Must dies with the error if it is non-nil.
// If the error is nil, Must is a no-op.
func Must(err error) {
	if err != nil {
		Die(err)
	}
}
