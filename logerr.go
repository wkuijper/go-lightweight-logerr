package logerr

import (
	"fmt"
	"os"
)

func F(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "ERROR: ")
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr, "")
}
