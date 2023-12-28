package internal

import (
	"context"
	"errors"

	"net/http"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/enamespace/tpl/internal/config"
)

type Server struct {
	cfg *config.Config
	*gin.Engine
	insecureServer *http.Server
}

func Run(cfg *config.Config) error {
	server, err := createServer(cfg)
	if err != nil {
		return err
	}

	return server.Run()
}

func (s *Server) Run() error {

	s.insecureServer = &http.Server{
		Addr:    s.cfg.InsecureServing.BindAddress,
		Handler: s,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Printf("Start to listening the incoming requests on http address: %s", s.cfg.InsecureServing.BindAddress)

		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err.Error())

			return err
		}

		log.Printf("Server on %s stopped", s.cfg.InsecureServing.BindAddress)
		return nil
	})

	eg.Go(func() error {
		log.Println("Start to https")
		return nil
	})

	eg.Go(func() error {
		log.Println("Start to grpc")
		return nil
	})

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())

	}

	return nil
}

func createServer(cfg *config.Config) (*Server, error) {

	server := &Server{}
	initRouter(server.Engine)

	return nil, nil
}
