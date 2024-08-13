package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Server interface {
	ListenAndServe(ctx context.Context, wg *sync.WaitGroup)
	AddHandler(method string, pattern string, handler func(c *gin.Context))
}

type GinServer struct {
	Engine     *gin.Engine
	Port       int64
	CtxTimeout time.Duration
	BasePath   string
}

func NewGinServer(port int64, timeout time.Duration) *GinServer {
	return &GinServer{
		Engine:     gin.New(),
		Port:       port,
		CtxTimeout: timeout,
	}
}

func (s *GinServer) SetBasePath(path string) *GinServer {
	s.BasePath = path
	log.Println("base path was set:", s.BasePath)
	return s
}

func (s *GinServer) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	gin.Recovery()

	s.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", s.Port),
		Handler:      s.Engine,
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicln(err)
		}
		<-ctx.Done()
		if err := server.Shutdown(ctx); err != nil {
			log.Panicln(err)
		}
	}()
}

func (s *GinServer) AddHandler(method string, pattern string, handler func(c *gin.Context)) {
	path := s.BasePath + pattern
	switch strings.ToUpper(method) {
	case http.MethodGet:
		s.Engine.GET(path, handler)
	case http.MethodPost:
		s.Engine.POST(path, handler)
	case http.MethodPut:
		s.Engine.PUT(path, handler)
	case http.MethodDelete:
		s.Engine.DELETE(path, handler)
	default:
		panic("invalid server method on added handler")
	}
}
