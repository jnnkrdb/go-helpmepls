package synchronization

import (
	"io"
	"os"
)

var logWriter io.Writer = os.Stdout

func SetLogger(w io.Writer) {
	logWriter = w
}
