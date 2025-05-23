package htmx

import "github.com/daarxwalker/gox"

const (
	atrributePrefix = "hx"
)

func Disable() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-disable")("true")
}

func DisabledElt() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-disabled-elt")()
}

func Ext(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-ext")(value...)
}

func Preserve() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-preserve")("true")
}
