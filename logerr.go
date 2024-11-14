package logerr

import (
	"io"
)

type LogSink interface {
	Log(level string, format string, args ...any)
}

func F(format string, args ...any) {
	fmt.Fprintf(os.Stdout, "ERROR: ")
	fmt.Fprintf(os.Stdout, format, args...)
	fmt.Fprintln(os.Stdout, "")
}

func FF(out io.Writer, format string, args ...any) {
	fmt.Fprintf(out, "ERROR: ")
	fmt.Fprintf(out, format, args...)
	fmt.Fprintln(out, "")
}

func SF(sink LogSink, format string, args ...any) {
	sink.Log("ERROR", format, args)
}
