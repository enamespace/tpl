package internal

import (
	"github.com/enamespace/tpl/internal/config"
	"github.com/enamespace/tpl/internal/options"
	"github.com/enamespace/tpl/pkg/app"
)

func NewApp(basename string) *app.App {
	opts := options.NewOptions()

	return app.NewApp(
		app.WithBasename(basename),
		app.WithCliOptions(opts),
		app.WithRunFunc(run(opts)),
	)
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		cfg := config.CreateConfigFromOptions(opts)

		return Run(cfg)
	}
}
