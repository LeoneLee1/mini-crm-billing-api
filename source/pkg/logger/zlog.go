package logger

import (
	"fmt"
	"io"
	"mini-crm-billing-api/source/services/constant"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(pretty bool) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", file, line)
	}

	var writer io.Writer = os.Stdout
	if pretty {
		writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}

	Log = zerolog.New(writer).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)
}

// helper
func Debug() *zerolog.Event { return Log.Debug() }
func Info() *zerolog.Event  { return Log.Info() }
func Warn() *zerolog.Event  { return Log.Warn() }
func Error() *zerolog.Event { return Log.Error() }

func GinZLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		dur := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		ip := c.ClientIP()

		var ev *zerolog.Event
		switch {
		case status >= 500:
			ev = Log.Error()
		case status >= 400:
			ev = Log.Warn()
		case status >= 300:
			ev = Log.Debug()
		default:
			ev = Log.Info()
		}

		ev.Str("method", method).
			Str("path", path).
			Str("client_ip", ip).
			Str("user_id", c.GetString("id")).
			Str("log_activity", c.GetString("log_id")).
			Str(constant.RequestIDKey, c.GetString(constant.RequestIDKey)).
			Int("status", status).
			Dur("latency", dur).
			Msg("http request")
	}
}
