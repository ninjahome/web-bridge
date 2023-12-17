package util

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"sync"
	"time"
)

var _logInstance *zerolog.Logger
var logOnce sync.Once
var logLevel = "debug"

func LogInst() *zerolog.Logger {
	logOnce.Do(func() {

		writer := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})
		out := zerolog.ConsoleWriter{Out: writer}
		out.TimeFormat = time.StampMilli
		logLvl, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			logLvl = zerolog.DebugLevel
		}
		logger := zerolog.New(out).
			Level(logLvl).
			With().
			Caller().
			Timestamp().
			Logger()
		_logInstance = &logger
	})

	return _logInstance
}

func SetLogLevel(ll string) {
	logLevel = ll
	logLvl, _ := zerolog.ParseLevel(logLevel)
	zerolog.SetGlobalLevel(logLvl)
}
