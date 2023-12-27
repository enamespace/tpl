package server

import (
	"github.com/enamespace/tpl/pkg/options"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type GenericGinServer struct {
	*gin.Engine
	GenericServerOptions *options.GenericServerOptions
	Mode                 string
}

func New(o *options.GenericServerOptions, mode string) *GenericGinServer {
	gin.SetMode(mode)
	s := &GenericGinServer{
		Engine:               gin.New(),
		GenericServerOptions: o,
		Mode:                 mode,
	}

	if o.EnableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}
	if o.EnableProfiling {
		pprof.Register(s.Engine)
	}

	return s
}
