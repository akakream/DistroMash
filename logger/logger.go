package logger

import (
	"io"
	"log"
	"os"

	"github.com/akakream/DistroMash/pkg/utils"
	kitlog "github.com/go-kit/log"
)

var logger kitlog.Logger

func Init() {
	var (
		logout io.Writer
		err    error
	)

	if utils.IsEnvDev() {
		logout = os.Stderr
	} else {
		logpath := os.Getenv("LOG_PATH")
		if _, err := os.Stat(logpath); err != nil {
			_, err := os.Create(logpath)
			if err != nil {
				log.Fatal(err)
			}
		}
		logout, err = os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(logout))
	logger = kitlog.With(logger, kitlog.DefaultTimestampUTC, "caller", kitlog.Caller(4))
}
