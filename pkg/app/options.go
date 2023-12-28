package app

import "github.com/enamespace/tpl/pkg/cli"

type CliOptions interface {
	Flags() cli.NamedFlagSets
	Validate() []error
}
