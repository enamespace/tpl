package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/enamespace/tpl/pkg/errors"
)

type Option func(*App)
type RunFunc func(basename string) error

type App struct {
	basename    string
	description string
	options     CliOptions
	runFunc     RunFunc
	cmd         *cobra.Command
}

func WithRunFunc(f RunFunc) Option {
	return func(a *App) {
		a.runFunc = f
	}
}

func WithBasename(basename string) Option {
	return func(a *App) {
		a.basename = basename
	}
}

func WithCliOptions(cliopts CliOptions) Option {
	return func(a *App) {
		a.options = cliopts
	}
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           a.basename,
		Long:          a.description,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true
	a.options.Flags()
	// a.options.

	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {

	if a.options != nil {
		if err := a.applyOptions(); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) applyOptions() error {
	if errs := a.options.Validate(); len(errs) > 0 {
		return errors.NewAggregate(errs)
	}

	return nil
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}

	a.buildCommand()
	return a
}
