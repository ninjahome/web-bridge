package util

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"io"
	"os"
	"sync"
	"time"
)

var _logInstance *zerolog.Logger
var logOnce sync.Once
var logLevel = "debug"

func LogInst() *zerolog.Logger {
	logOnce.Do(func() {

		file, err := os.OpenFile("game.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		writer := diode.NewWriter(os.Stderr, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger Dropped %d messages", missed)
		})

		multi := io.MultiWriter(writer, file)
		out := zerolog.ConsoleWriter{Out: multi}
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
	logLvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		fmt.Println("set log level err:", err)
		return
	}
	zerolog.SetGlobalLevel(logLvl)
	fmt.Println("set log level success:", ll)
}
