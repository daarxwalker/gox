package htmx

import "github.com/daarxwalker/gox"

func Boost() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-boost")("true")
}

func Disinherit(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-disinherit")(value...)
}
