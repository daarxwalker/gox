package datastar

import (
	"strings"
	
	"github.com/daarxwalker/gox"
)

func CustomValidity(value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-custom-validity")
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(value)
}

func Intersect(value string, modifiers ...Modifier) gox.Node {
	return On(INTERSECT, value, modifiers...)
}

func Interval(value string, modifiers ...Modifier) gox.Node {
	return On(INTERVAL, value, modifiers...)
}

func AnySignalChange(value string, modifiers ...Modifier) gox.Node {
	return On(SIGNAL_CHANGE, value, modifiers...)
}

func SignalChange(name string, value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString(SIGNAL_CHANGE)
	builder.WriteString("-")
	builder.WriteString(name)
	return On(builder.String(), value, modifiers...)
}

func Persist(signals []string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-persist")
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(strings.Join(signals, " "))
}

func ReplaceURL(value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-replace-url")
	attributeName := applyModifiers(builder.String(), modifiers...)
	if !strings.HasPrefix(value, "`") && !strings.HasSuffix(value, "`") {
		valueBuilder := new(strings.Builder)
		valueBuilder.WriteString("`")
		valueBuilder.WriteString(value)
		valueBuilder.WriteString("`")
		value = valueBuilder.String()
	}
	return gox.CreateAttribute[string](attributeName)(value)
}

func ScrollIntoView(modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-scroll-into-view")
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)()
}

func ViewTransition(value string, modifiers ...Modifier) gox.Node {
	builder := new(strings.Builder)
	builder.WriteString("data-view-transition")
	attributeName := applyModifiers(builder.String(), modifiers...)
	return gox.CreateAttribute[string](attributeName)(value)
}
