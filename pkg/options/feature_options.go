package options

import "github.com/spf13/pflag"

type FeatureOptions struct {
	EnableProfiling bool `json:"profiling"      mapstructure:"profiling"`
	EnableMetrics   bool `json:"enable-metrics" mapstructure:"enable-metrics"`
}

func NewFeatureOptions() *FeatureOptions {
	return &FeatureOptions{
		EnableMetrics:   false,
		EnableProfiling: false,
	}
}

func (o *FeatureOptions) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&o.EnableMetrics, "feature.enable-metrics", o.EnableMetrics,
		"enable server metric on /metrics")
	fs.BoolVar(&o.EnableProfiling, "feature.enable-profiling", o.EnableProfiling,
		"enable server profiling on host:port/debug/pprof/",
	)
}

func (o *FeatureOptions) Validate() []error {
	return []error{}
}
