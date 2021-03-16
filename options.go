package live

import (
	`github.com/class100/core`
)

type (
	// Option 配置选项
	Option func(*options)

	options struct {
		clientOptions []core.Option

		// Endpoint 服务端点
		Endpoint string
	}
)

func defaultOptions() options {
	return options{}
}

// WithEndpoint 配置服务端点
func WithEndpoint(endpoint string) Option {
	return func(options *options) {
		options.Endpoint = endpoint
	}
}
