package mainapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type mainApiServer struct {
	httpServ *http.Server
}

func NewMainApiServer() *mainApiServer {
	return &mainApiServer{}
}

func (s *mainApiServer) RunMainApiServer(port string, handler *gin.Engine) error {
	s.httpServ = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	return s.httpServ.ListenAndServe()
}

func (s *mainApiServer) Shatdown(ctx context.Context) error {
	return s.httpServ.Shutdown(ctx)
}
