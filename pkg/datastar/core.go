package datastar

import (
	"fmt"
	"strings"
	
	"github.com/daarxwalker/gox"
)

func Signals(value any, modifiers ...Modifier) gox.Node {
	return gox.CreateAttribute[string](applyModifiers("data-signals", modifiers...))(createAttributeValue(value))
}

func Signal(name string, value any, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-signals-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(fmt.Sprint(value))
}

func Computed(name string, value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-computed-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(value)
}

func StarIgnore() gox.Node {
	return gox.CreateAttribute[string]("data-star-ignore")()
}
