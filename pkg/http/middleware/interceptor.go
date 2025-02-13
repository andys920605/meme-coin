package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/pkg/errors"
	"github.com/andys920605/meme-coin/pkg/http/crypto"
	"github.com/andys920605/meme-coin/pkg/http/gcontext"
	"github.com/andys920605/meme-coin/pkg/http/template_response"
)

type Interceptor struct {
	crypto *crypto.Crypto
}

func NewInterceptor() *Interceptor {
	return &Interceptor{
		crypto: crypto.NewCrypto(),
	}
}

func (g *Interceptor) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		gcontext.SetSource(c, gcontext.ParseSource(c))

		defer func() {
			if err := recover(); err != nil {
				g.setErrorContext(c,
					fmt.Errorf("panic: %s", err).Error(),
					errors.InternalServerPanic.Code(),
					string(debug.Stack()),
				)
				g.errorResponse(c, errors.InternalServerPanic)
				return
			}
		}()

		if err := g.decodeRequestBody(c); err != nil {
			g.setErrorContext(c, err.Error(), errors.InternalServerError.Code(), errors.StackTracer(err))
			g.errorResponse(c, errors.InternalServerError)
			return
		}

		if c.ContentType() == gin.MIMEJSON {
			buf, _ := c.GetRawData()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(buf))

			var body interface{}
			_ = json.Unmarshal(buf, &body)
			bodyBytes, _ := json.Marshal(body)
			c.Set(gcontext.ContextKeyReqBody, string(bodyBytes))
		}

		c.Next()

		if len(c.Errors) == 0 {
			status := c.GetInt(gcontext.ContextKeyRespStatus)
			resp, _ := c.Get(gcontext.ContextKeyRespBody)
			g.response(c, status, resp)
			return
		}

		err := c.Errors[0].Err
		customError := errors.CauseCustomError(err)
		// If the error is not a custom error, it will be handled by the default error handler.
		if customError.IsEmpty() {
			g.setErrorContext(c, err.Error(), errors.InternalServerError.Code(), errors.StackTracer(err))
			g.errorResponse(c, errors.InternalServerError)
			return
		}

		g.setErrorContext(c, err.Error(), customError.Code(), errors.StackTracer(err))
		g.errorResponse(c, customError)
	}
}

func (g *Interceptor) decodeRequestBody(c *gin.Context) error {
	if gcontext.GetSource(c) != gcontext.SourceApp {
		return nil
	}

	if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete ||
		c.Request.Method == http.MethodOptions {
		return nil
	}

	userId, ok := gcontext.GetUserIdString(c)
	if !ok {
		return errors.MissAuthorization.New("missing user id")
	}

	buf, err := c.GetRawData()
	if err != nil {
		return errors.InternalServerError.Wrap(err, "get raw data")
	}

	if len(buf) == 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(buf))
		return nil
	}

	key, err := g.crypto.GenerateKey(userId)
	if err != nil {
		return errors.InternalServerError.Wrap(err, "generate chacha20 key")
	}

	decodedReq, err := g.crypto.Decrypt(string(buf), key)
	if err != nil {
		return errors.InternalServerError.Wrap(err, "decrypt request")
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(decodedReq))
	return nil
}

func (g *Interceptor) setErrorContext(c *gin.Context, err string, code int, stackTrace any) {
	c.Set(gcontext.ContextKeyError, err)
	c.Set(gcontext.ContextKeyCode, code)
	c.Set(gcontext.ContextKeyStackTrace, stackTrace)
}

func (g *Interceptor) response(c *gin.Context, status int, resp any) {
	userId, ok := gcontext.GetUserIdString(c)
	if gcontext.GetSource(c) == gcontext.SourceApp && ok {
		key, _ := g.crypto.GenerateKey(userId)
		encryptedByte, _ := g.crypto.Encrypt(resp, key)
		c.String(status, "%s", encryptedByte)
		return
	}

	c.JSON(status, resp)
}

func (g *Interceptor) errorResponse(c *gin.Context, err errors.CustomError) {
	resp := template_response.Error(err.Code(), err.Message())
	g.response(c, err.Status().ToHTTPStatus(), resp)
}
