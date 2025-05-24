package datastar

import (
	"fmt"
	"strings"
	
	"github.com/daarxwalker/gox"
)

func Attrs(value any, modifiers ...Modifier) gox.Node {
	return gox.CreateAttribute[string](applyModifiers("data-attr", modifiers...))(createAttributeValue(value))
}

func Attr(name string, value any, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-attr-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(fmt.Sprint(value))
}

func Bind(name string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-bind-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)()
}

func Classes(value any, modifiers ...Modifier) gox.Node {
	return gox.CreateAttribute[string](applyModifiers("data-class", modifiers...))(createAttributeValue(value))
}

func Class(name string, value any, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-class-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(fmt.Sprint(value))
}

func On(name string, value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-on-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(value)
}

func Ref(name string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-ref-")
	builder.WriteString(name)
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)()
}

func Show(condition string) gox.Node {
	return gox.CreateAttribute[string]("data-show")(condition)
}

func Text(value string) gox.Node {
	return gox.CreateAttribute[string]("data-text")(value)
}
