package gontp

import (
	"github.com/beevik/ntp"
	"io"
	"os"
)

func Run() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, _ = io.WriteString(os.Stderr, error.Error(err))
		os.Exit(1)
	}
	_, err = io.WriteString(os.Stdout, time.String())
	if err != nil {
		os.Exit(1)
	}
}
