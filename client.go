package live

import (
	`fmt`
	`strings`

	`github.com/class100/core`
	`github.com/storezhang/gox`
)

type (
	ApiPath string

	Client interface {
		liver
	}

	httpClient struct {
		c *core.HttpSignatureClient

		options options
	}
)

// newHttpClient 创建默认客户端
func NewClient(options ...Option) (client Client, err error) {
	var hsc *core.HttpSignatureClient

	appliedOptions := defaultOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	if "" == strings.TrimSpace(appliedOptions.Endpoint) {
		err = ErrMustSetEndpoint

		return
	}

	if hsc, err = core.NewHttpSignatureClient(appliedOptions.clientOptions...); nil != err {
		return
	}

	client = &httpClient{
		c:       hsc,
		options: appliedOptions,
	}

	return
}

func (client *httpClient) requestApi(
	path ApiPath,
	method core.HttpMethod,
	data interface{},
	rsp interface{},
	params ...gox.HttpParameter,
) (err error) {
	var url string

	url = fmt.Sprintf("%s/api/%s", client.options.Endpoint, path)

	if rsp == nil {
		rsp = new(interface{})
	}

	return client.c.RequestApi(url, method, data, rsp, params...)
}
