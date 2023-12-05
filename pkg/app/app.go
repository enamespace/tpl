package app

import (
	"log"

	"github.com/spf13/viper"
)

type Option func(*App)

type App struct {
	Cfg interface{}
}

func WithConfig(cfg interface{}) Option {
	return func(a *App) {
		a.Cfg = cfg
	}
}

func (a *App) build() {
	if err := viper.Unmarshal(a.Cfg); err != nil {
		log.Fatal(err)
	}
}

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}

	a.build()
	return a
}
