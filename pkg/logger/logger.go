package logger

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger(filename string, maxsize int, maxbackup int, compress bool, level string) {
	logger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize, // megabytes
		MaxBackups: maxbackup,
		// MaxAge: 28, //days
		Compress: compress, // disabled by default
	}
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(lvl)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// use lumberjack to write to implement rotation.
	log.SetOutput(logger)
}

func WithTrace(ctx *gin.Context) *log.Entry {
	fields := log.Fields{}
	if len(ctx.GetString("X-Trace-ID")) > 0 {
		fields["trace_id"] = ctx.GetString("X-Trace-ID")
	}
	if len(ctx.GetString("X-Span-ID")) > 0 {
		fields["span_id"] = ctx.GetString("X-Span-ID")
	}
	return log.WithFields(fields)
}
