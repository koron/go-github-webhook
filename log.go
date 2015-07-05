package webhook

import (
	"log"
	"os"
)

var (
	Logger *log.Logger = log.New(os.Stderr, "", log.LstdFlags)
)

func logErr(err error) {
	if Logger == nil {
		return
	}
	Logger.Printf("webhoook error: %s", err)
	return
}
