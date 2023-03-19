package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	outputStdout := false
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if outputStdout {
		SetOutput(os.Stdout)
	} else {
		outputFile, err := os.Create("logs.txt")
		if err != nil {
			fmt.Println("Switched logging to Stdout because of log file open error")
			SetOutput(os.Stdout)
			return
		}
		SetOutput(outputFile)
	}
}

func SetOutput(out io.Writer) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        out,
		TimeFormat: time.RFC1123,
		NoColor:    !(out == os.Stdout || out == os.Stderr),
	})
}

func Debug(msg string) {
	_, filename, line, _ := runtime.Caller(1)
	log.Debug().Msg(fmt.Sprintf("%s:%d: %s", filepath.Base(filename), line, msg))
}

func Info(msg string) {
	_, filename, line, _ := runtime.Caller(1)
	log.Info().Msg(fmt.Sprintf("%s:%d: %s", filepath.Base(filename), line, msg))
}

func InfoWithoutCaller(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	_, filename, line, _ := runtime.Caller(1)
	log.Warn().Msg(fmt.Sprintf("%s:%d: %s", filepath.Base(filename), line, msg))
}

func Error(err error) {
	log.Error().Err(err).Msg("")
}