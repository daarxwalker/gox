package datastar

import (
	"strings"
	
	"github.com/daarxwalker/gox"
)

func Indicator(name string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-indicator-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)()
}
