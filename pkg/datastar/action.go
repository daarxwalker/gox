package datastar

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	CONTENT_TYPE      = "contentType"
	INCLUDE_LOCAL     = "includeLocal"
	SELECTOR          = "selector"
	HEADERS           = "headers"
	OPEN_WHEN_HIDDEN  = "openWhenHidden"
	RETRY_INTERVAL    = "retryInterval"
	RETRY_SCALER      = "retryScaler"
	RETRY_MAX_WAIT_MS = "retryMaxWaitMs"
	RETRY_MAX_COUNT   = "retryMaxCount"
	ABORT             = "abort"
)

func GET(endpoint string, config ...Map) string {
	builder := new(strings.Builder)
	builder.WriteString("@get('")
	builder.WriteString(endpoint)
	if len(config) > 0 {
		builder.WriteString("',")
		marshaledConfig, marshalErr := json.Marshal(config[0])
		if marshalErr == nil {
			builder.WriteString(strings.ReplaceAll(string(marshaledConfig), `"`, `'`))
		}
		builder.WriteString(")")
		return builder.String()
	}
	builder.WriteString("')")
	return builder.String()
}

func POST(endpoint string, config ...Map) string {
	builder := new(strings.Builder)
	builder.WriteString("@post('")
	builder.WriteString(endpoint)
	if len(config) > 0 {
		builder.WriteString("',")
		marshaledConfig, marshalErr := json.Marshal(config[0])
		if marshalErr == nil {
			builder.WriteString(strings.ReplaceAll(string(marshaledConfig), `"`, `'`))
		}
		builder.WriteString(")")
		return builder.String()
	}
	builder.WriteString("')")
	return builder.String()
}

func PUT(endpoint string, config ...Map) string {
	builder := new(strings.Builder)
	builder.WriteString("@put('")
	builder.WriteString(endpoint)
	if len(config) > 0 {
		builder.WriteString("',")
		marshaledConfig, marshalErr := json.Marshal(config[0])
		if marshalErr == nil {
			builder.WriteString(strings.ReplaceAll(string(marshaledConfig), `"`, `'`))
		}
		builder.WriteString(")")
		return builder.String()
	}
	builder.WriteString("')")
	return builder.String()
}

func PATCH(endpoint string, config ...Map) string {
	builder := new(strings.Builder)
	builder.WriteString("@patch('")
	builder.WriteString(endpoint)
	if len(config) > 0 {
		builder.WriteString("',")
		marshaledConfig, marshalErr := json.Marshal(config[0])
		if marshalErr == nil {
			builder.WriteString(strings.ReplaceAll(string(marshaledConfig), `"`, `'`))
		}
		builder.WriteString(")")
		return builder.String()
	}
	builder.WriteString("')")
	return builder.String()
}

func DELETE(endpoint string, config ...Map) string {
	builder := new(strings.Builder)
	builder.WriteString("@delete('")
	builder.WriteString(endpoint)
	if len(config) > 0 {
		builder.WriteString("',")
		marshaledConfig, marshalErr := json.Marshal(config[0])
		if marshalErr == nil {
			builder.WriteString(strings.ReplaceAll(string(marshaledConfig), `"`, `'`))
		}
		builder.WriteString(")")
		return builder.String()
	}
	builder.WriteString("')")
	return builder.String()
}

func Clipboard(value string) string {
	builder := new(strings.Builder)
	builder.WriteString("@clipboard('")
	builder.WriteString(value)
	builder.WriteString("')")
	return builder.String()
}

func SetAll(value string, condition any) string {
	builder := new(strings.Builder)
	builder.WriteString("@setAll('")
	builder.WriteString(value)
	builder.WriteString("',")
	switch c := condition.(type) {
	case string:
		builder.WriteString(c)
	case bool:
		if c {
			builder.WriteString("true")
		} else {
			builder.WriteString("false")
		}
	default:
		builder.WriteString(fmt.Sprint(c))
	}
	builder.WriteString(")")
	return builder.String()
}

func ToggleAll(value string) string {
	builder := new(strings.Builder)
	builder.WriteString("@toggleAll('")
	builder.WriteString(value)
	builder.WriteString("')")
	return builder.String()
}
