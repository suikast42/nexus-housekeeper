package logging

import (
	"context"
	"github.com/rs/zerolog"
	zeroLogger "github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"time"
)

func LoggerWithContext(ctx context.Context) *zerolog.Logger {
	logger := Logger()
	withContext := logger.WithContext(ctx)
	return zerolog.Ctx(withContext)
}

func Logger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	//multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout)
	//
	//logger := zerolog.New(multi).With().Timestamp().Logger()
	return zeroLogger.Output(consoleWriter)
}
