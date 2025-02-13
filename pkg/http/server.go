package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/andys920605/meme-coin/pkg/http/middleware"
	"github.com/andys920605/meme-coin/pkg/logging"
)

type Router interface {
	Register(engine *gin.Engine)
}

type Server struct {
	ginEngine           *gin.Engine
	logger              *logging.Logging
	connectionCloseChan chan func()
}

func NewServer(logger *logging.Logging, serviceName string) *Server {
	gin.SetMode(gin.ReleaseMode)
	ginEngine := gin.New()

	// default middleware
	ginEngine.Use(middleware.NewTraceHandler(serviceName))
	ginEngine.Use(middleware.NewLoggerHandler(logger))

	return &Server{
		ginEngine:           ginEngine,
		logger:              logger,
		connectionCloseChan: make(chan func(), 1),
	}
}

// RegisterDefaultCORS registers CORS middleware with default config.
func (s *Server) RegisterDefaultCORS() {
	s.ginEngine.Use(cors.Default())
}

// RegisterCORS registers CORS middleware with custom config.
func (s *Server) RegisterCORS(customConfig cors.Config) {
	s.ginEngine.Use(cors.New(customConfig))
}

// RegisterRouter registers all the routes for the API.
func (s *Server) RegisterRouter(router Router) {
	router.Register(s.ginEngine)
}

// Run runs the HTTP server.
func (s *Server) Run(port string) {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: s.ginEngine,
	}

	go func() {
		sigint := make(chan os.Signal, 2)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		s.logger.Infof("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			s.logger.Emergencyf("http server shutdown error: %s", err)
		}
		close(s.connectionCloseChan)
	}()

	s.logger.Infof("run http server %s address success", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Errorf("run http server %s address error: %s", httpServer.Addr, err)
	}

	shutdown := <-s.connectionCloseChan
	if shutdown != nil {
		shutdown()
	}
}

// SetShutdownHandler sets the shutdown handler.
func (s *Server) SetShutdownHandler(shutdown func()) {
	s.connectionCloseChan <- shutdown
}
