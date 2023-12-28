package options

import (
	"github.com/enamespace/tpl/pkg/cli"
	"github.com/enamespace/tpl/pkg/options"
)

type Options struct {
	InsecureServing *options.InsecureServerOptions `json:"insecure" mapstructure:"insecure"`
	FeatureOptions  *options.FeatureOptions        `json:"feature"  mapstructure:"feature"`
	MySQLOptions    *options.MySQLOptions          `json:"mysql"    mapstructure:"mysql"`
}

func (o *Options) Flags() (fss cli.NamedFlagSets) {

	o.InsecureServing.AddFlags(fss.FlagSet("insecure"))
	o.FeatureOptions.AddFlags(fss.FlagSet("feature"))
	o.MySQLOptions.AddFlags(fss.FlagSet("mysql"))

	return fss
}

func (o *Options) Validate() []error {
	errs := []error{}

	errs = append(errs, o.InsecureServing.Validate()...)
	errs = append(errs, o.FeatureOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)

	return errs
}

func NewOptions() *Options {
	return &Options{
		InsecureServing: options.NewInsecureServingOptions(),
		FeatureOptions:  options.NewFeatureOptions(),
		MySQLOptions:    options.NewMySQLOptions(),
	}
}
