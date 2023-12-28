package options

import "github.com/spf13/pflag"

type InsecureServerOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
}

func NewInsecureServingOptions() *InsecureServerOptions {
	return &InsecureServerOptions{
		BindAddress: "127.0.0.1",
		BindPort:    8080,
	}
}

func (o *InsecureServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.BindAddress, "insecure.bind-address", o.BindAddress,
		"IP address for insecure")
	fs.IntVar(&o.BindPort, "insecure.bind-port", o.BindPort,
		"Listening port for insecure")
}

func (o *InsecureServerOptions) Validate() []error {
	errs := []error{}
	return errs
}
