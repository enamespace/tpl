package options

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

func (o *InsecureServerOptions) Validate() []error {
	errs := []error{}
	return errs
}
