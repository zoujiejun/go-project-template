package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-project-template/internal/pkg/logger"
	"io"
	"time"
)

type bodyWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (b *bodyWriter) Write(body []byte) (int, error) {
	b.Body.Write(body)
	return b.ResponseWriter.Write(body)
}

type RequestLogging struct {
	logger *logger.Logger
}

func NewRequestLogging(logger *logger.Logger) *RequestLogging {
	return &RequestLogging{logger: logger}
}

func (rl *RequestLogging) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		blw := &bodyWriter{
			Body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = blw

		requestBody, _ := io.ReadAll(ctx.Request.Body)
		ctx.Request.Body = io.NopCloser(bytes.NewReader(requestBody))

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)
		response := blw.Body.String()

		kvs := []interface{}{
			"status", ctx.Writer.Status(),
			"method", ctx.Request.Method,
			"path", path,
			"query", query,
			"ip", ctx.ClientIP(),
			"userAgent", ctx.Request.UserAgent(),
			"latency", latency.Milliseconds(),
			"request", string(requestBody),
			"response", response,
		}

		if len(ctx.Errors) > 0 {
			for idx, e := range ctx.Errors {
				kvs = append(kvs, fmt.Sprintf("error_%d", idx), e.Error(), fmt.Sprintf("stact_%d", idx), fmt.Sprintf("%+v", e.Err))
			}
		}

		rl.logger.Info("request log", kvs...)
	}
}
