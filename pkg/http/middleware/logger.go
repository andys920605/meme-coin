package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/pkg/http/gcontext"
	"github.com/andys920605/meme-coin/pkg/logging"
	"github.com/andys920605/meme-coin/pkg/trace"
)

func NewLoggerHandler(logger *logging.Logging) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		traceId := trace.GetTraceIDFromContext(c.Request.Context())

		fields := logging.Fields{
			"endpoint":   c.Request.URL.Path,
			"method":     c.Request.Method,
			"user_agent": c.Request.UserAgent(),
			"client_ip":  c.ClientIP(),
			"trace_id":   traceId,
		}

		if len(c.Request.URL.RawQuery) > 0 {
			fields["raw_query"] = c.Request.URL.RawQuery
		}

		if c.GetHeader("Authorization") != "" {
			fields["authorization"] = c.GetHeader("Authorization")
		}

		c.Next()

		if body, ok := c.Get(gcontext.ContextKeyReqBody); ok {
			fields["request_body"] = body
		}

		if userId, ok := gcontext.GetUserIdString(c); ok {
			fields["user_id"] = userId
		}

		fields["elapsed_time"] = time.Since(startTime).Milliseconds()
		fields["status_code"] = c.Writer.Status()
		if c.Writer.Status() >= 200 && c.Writer.Status() < 300 {
			if c.FullPath() == "/healthz" || c.Request.Method == http.MethodOptions {
				return
			}

			logger.WithFields(fields).Infof(
				"[HTTP Server] %s %s request succeed",
				c.Request.Method,
				c.Request.URL.Path,
			)
			return
		}

		fields["error"], _ = c.Get(gcontext.ContextKeyError)
		fields["code"], _ = c.Get(gcontext.ContextKeyCode)
		fields["stack_trace"], _ = c.Get(gcontext.ContextKeyStackTrace)
		if c.Writer.Status() >= 400 && c.Writer.Status() < 500 {
			logger.WithFields(fields).Warningf(
				"[HTTP Server] %s %s request failed",
				c.Request.Method,
				c.Request.URL.Path,
			)
			return
		}

		logger.WithFields(fields).Errorf(
			"[HTTP Server] %s %s request failed",
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
