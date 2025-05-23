package htmx

import "github.com/daarxwalker/gox"

func Indicator(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-indicator")(value...)
}
