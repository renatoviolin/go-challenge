// https://github.com/dn365/gin-zerolog/blob/master/gin_zerolog.go
package logger

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ginHands struct {
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
}

var Logger zerolog.Logger

func InitLog() {
	fileLogger, _ := os.OpenFile("app.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	writer := io.MultiWriter(os.Stdout, fileLogger)
	Logger = zerolog.New(writer).With().Timestamp().Logger()
}

func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "request"
		}
		cData := &ginHands{
			SerName:    "http",
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			MsgStr:     msg,
		}

		httpLogWrite(cData)
	}
}

func httpLogWrite(data *ginHands) {
	switch {
	case data.StatusCode >= 400 && data.StatusCode < 500:
		{
			Logger.Warn().Str("from", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("ip", data.ClientIP).Msg(data.MsgStr)
		}
	case data.StatusCode >= 500:
		{
			Logger.Error().Str("from", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("ip", data.ClientIP).Msg(data.MsgStr)
		}
	default:
		Logger.Info().Str("from", data.SerName).Str("method", data.Method).Str("path", data.Path).Dur("resp_time", data.Latency).Int("status", data.StatusCode).Str("ip", data.ClientIP).Msg(data.MsgStr)
	}
}
func LogError(from string, msg string) {
	Logger.Error().Str("from", from).Str("message", msg).Send()
}

func LogInfo(from string, msg string) {
	Logger.Info().Str("from", from).Str("message", msg).Send()
}

func LogFatal(from string, msg string) {
	Logger.Fatal().Str("from", from).Str("message", msg).Send()
}
