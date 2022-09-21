package helper

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lizhanfei/easygolib/log"
	"go.uber.org/zap/zapcore"
	"user/conf"
	"user/env"
	"strconv"
	"time"
)

var Logger log.Zlog

func InitLog() {
	config := log.LoggerConfig{
		Stdout:   true,
		ZapLevel: zapcore.InfoLevel, //最低日志等级
		Path:     env.LogPath,
		LogName:  env.AppName,
		Log2File: false,
		Module:   env.AppName,
	}
	switch conf.LogConf.Level {
	case "debug":
		config.ZapLevel = zapcore.DebugLevel
	case "info":
		config.ZapLevel = zapcore.InfoLevel
	case "error":
		config.ZapLevel = zapcore.ErrorLevel
	}
	Logger = log.NewZLog(getRequestId, getUriPath, config, nil)
}

func getRequestId(ctx context.Context) string {
	var makeRequestId = func() string {
		usec := uint64(time.Now().UnixNano())
		return strconv.FormatUint(usec&0x7FFFFFFF|0x80000000, 10)
	}
	var requestID string
	if c, ok := ctx.(*gin.Context); ok && c != nil {
		requestID, _ = c.Value(env.ContextKeyRequestID).(string)
		if "" != requestID {
			return requestID
		}
		if c.Request != nil && c.Request.Header != nil {
			requestID = c.Request.Header.Get(env.TraceHeaderKey)
		}
		if requestID == "" {
			requestID = makeRequestId()
		}
		c.Set(env.ContextKeyRequestID, requestID)
		return requestID
	}
	return makeRequestId()
}

func getUriPath(ctx context.Context) string {
	if c, ok := ctx.(*gin.Context); ok && c != nil {
		uri, _ := c.Value(env.UriNow).(string)
		if "" != uri {
			return uri
		}
		uri = c.Request.URL.Path
		c.Set(env.UriNow, uri)
		return uri
	}
	return ""
}
