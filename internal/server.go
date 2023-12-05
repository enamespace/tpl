package internal

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/enamespace/tpl/internal/config"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	cfg *config.Config
}

func Run(cfg *config.Config) error {
	server, err := New(cfg)
	if err != nil {
		return err
	}

	return server.Run()
}

func New(cfg *config.Config) (*Server, error) {
	if err := cfg.Validate(); err != nil {
		return nil, errors.New("err")
	}

	return &Server{cfg: cfg}, nil

}

func (srv *Server) PrepareRun() error {
	var eg errgroup.Group

	eg.Go(func() error {
		log.Println("Start to init 1")
		return nil
	})

	eg.Go(func() error {
		log.Println("Start to init 2")
		return nil
	})

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())

	}

	return nil
}

func (srv *Server) Run() error {

	return nil
}
