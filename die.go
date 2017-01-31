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
