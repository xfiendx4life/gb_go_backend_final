package logger

import (
	"io"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(level zapcore.Level, filelog string) *zap.SugaredLogger {
	var output io.Writer
	var encoder zapcore.Encoder
	var err error
	// choosing file or stderr
	if filelog != "" {
		output, err = os.Create(filelog)
		if err != nil {
			log.Printf("can't create logger file")
		} // we are going to use file as log output
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) // using json for file
	} else {
		output = os.Stderr
		encoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) // using simple console

	}
	writeSyncer := zapcore.AddSync(output)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	return zap.New(core).Sugar()
}
