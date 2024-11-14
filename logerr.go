package logerr

import (
	"io"
)

type LogSink interface {
	Log(level string, format string, args ...any)
}

func F(format string, args ...any) {
	// DISABLED
}

func FF(out io.Writer, format string, args ...any) {
	// DISABLED
}

func SF(sink LogSink, format string, args ...any) {
	// DISABLED
}
