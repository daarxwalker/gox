package gox

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type nodeRenderer struct {
	node
}

const (
	openBracket  = "{"
	closeBracket = "}"
)

func (n nodeRenderer) render() string {
	switch n.nodeType {
	case nodeFragment:
		builder := new(strings.Builder)
		if len(n.attributes) > 0 {
			builder.WriteString(n.renderAttributes())
		}
		if len(n.children) > 0 {
			builder.WriteString(n.renderChildren())
		}
		return builder.String()
	case nodeText:
		return n.renderText()
	default:
		return n.renderElement()
	}
}

func (n nodeRenderer) renderElement() string {
	builder := new(strings.Builder)
	builder.WriteString("<")
	builder.WriteString(n.name)
	isVoid := slices.Index(voidElements, n.name) > 0
	if len(n.attributes) > 0 {
		builder.WriteString(" ")
		builder.WriteString(n.renderAttributes())
	}
	if isVoid {
		builder.WriteString(" />")
	}
	if !isVoid {
		builder.WriteString(">")
		if len(n.children) > 0 {
			builder.WriteString(n.renderChildren())
		}
		builder.WriteString("</")
		builder.WriteString(n.name)
		builder.WriteString(">")
	}
	return builder.String()
}

func (n nodeRenderer) renderAttribute() string {
	if n.value == nil {
		return n.name
	}
	builder := new(strings.Builder)
	value := fmt.Sprint(n.value)
	if strings.HasPrefix(value, openBracket) && strings.HasSuffix(value, closeBracket) {
		value = strings.ReplaceAll(value, `"`, `'`)
	}
	builder.WriteString(n.name)
	builder.WriteString(`="`)
	builder.WriteString(value)
	builder.WriteString(`"`)
	return builder.String()
}

func (n nodeRenderer) renderAttributes() string {
	size := len(n.attributes)
	if size == 0 {
		return ""
	}
	builder := new(strings.Builder)
	for i, a := range n.attributes {
		builder.WriteString(nodeRenderer{a}.renderAttribute())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func (n nodeRenderer) renderChildren() string {
	builder := new(strings.Builder)
	for _, ch := range n.children {
		switch ch.nodeType {
		case nodeRaw:
			builder.WriteString(nodeRenderer{ch}.renderRaw())
		case nodeElement:
			builder.WriteString(nodeRenderer{ch}.renderElement())
		case nodeAttribute:
			builder.WriteString(nodeRenderer{ch}.renderAttribute())
		case nodeText:
			builder.WriteString(nodeRenderer{ch}.renderText())
		}
	}
	return builder.String()
}

func (n nodeRenderer) renderText() string {
	return n.createStringValue()
}

func (n nodeRenderer) renderRaw() string {
	return n.createStringValue()
}

func (n nodeRenderer) createStringValue() string {
	switch v := n.value.(type) {
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	case error:
		return v.Error()
	default:
		return fmt.Sprint(v)
	}
}
