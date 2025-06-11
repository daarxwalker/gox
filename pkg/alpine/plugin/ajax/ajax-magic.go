package ajax

import (
	"encoding/json"
	"strings"
)

type ajaxOptions struct {
	Target  string   `json:"target,omitzero"`
	Body    any      `json:"body,omitzero"`
	Method  string   `json:"method"`
	Targets []string `json:"targets,omitzero"`
	Headers Map      `json:"headers,omitzero"`
	Focus   bool     `json:"focus,omitzero"`
	Sync    bool     `json:"sync,omitzero"`
}

type Option func(*ajaxOptions)

func (o ajaxOptions) Bytes() []byte {
	bytes, err := json.Marshal(o)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func (o ajaxOptions) String() string {
	bytes := o.Bytes()
	if len(bytes) == 0 {
		return "{}"
	}
	return string(bytes)
}

func (o ajaxOptions) EscapedString() string {
	bytes := o.Bytes()
	if len(bytes) == 0 {
		return "{}"
	}
	return strings.ReplaceAll(string(bytes), `"`, `'`)
}

func REQUEST(url string, options ...Option) string {
	return createAjax("", url, options...)
}

func GET(url string, options ...Option) string {
	return createAjax("get", url, options...)
}

func POST(url string, options ...Option) string {
	return createAjax("post", url, options...)
}

func PUT(url string, options ...Option) string {
	return createAjax("put", url, options...)
}

func PATCH(url string, options ...Option) string {
	return createAjax("patch", url, options...)
}

func DELETE(url string, options ...Option) string {
	return createAjax("delete", url, options...)
}

func WithMethod(method string) Option {
	return func(options *ajaxOptions) {
		options.Method = method
	}
}

func WithBody(body any) Option {
	return func(options *ajaxOptions) {
		options.Body = body
	}
}

func WithTarget(target string) Option {
	return func(options *ajaxOptions) {
		options.Target = target
	}
}

func WithTargets(targets ...string) Option {
	return func(options *ajaxOptions) {
		options.Targets = targets
	}
}

func WithHeaders(headers Map) Option {
	return func(options *ajaxOptions) {
		options.Headers = headers
	}
}

func WithFocus(focus bool) Option {
	return func(options *ajaxOptions) {
		options.Focus = focus
	}
}

func WithSync(sync bool) Option {
	return func(options *ajaxOptions) {
		options.Sync = sync
	}
}

func createAjax(method, url string, options ...Option) string {
	opts := &ajaxOptions{Method: method}
	for _, option := range options {
		option(opts)
	}
	return "$ajax('" + url + "'," + opts.EscapedString() + ")"
}
