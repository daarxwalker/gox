package gox

import (
	"fmt"
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
	builder := new(strings.Builder)
	switch n.nodeType {
	case nodeFragment:
		if len(n.attributes) > 0 {
			n.renderAttributes(builder)
		}
		if len(n.children) > 0 {
			n.renderChildren(builder)
		}
	case nodeAttribute:
		n.renderAttribute(builder)
	case nodeText:
		builder.WriteString(n.renderText())
	default:
		n.renderElement(builder)
	}
	return builder.String()
}

func (n nodeRenderer) renderElement(builder *strings.Builder) {
	builder.WriteString("<")
	builder.WriteString(n.name)
	_, isVoid := voidElementsMap[n.name]
	if len(n.attributes) > 0 {
		builder.WriteString(" ")
		n.renderAttributes(builder)
	}
	if isVoid {
		builder.WriteString(" />")
	}
	if !isVoid {
		builder.WriteString(">")
		if len(n.children) > 0 {
			n.renderChildren(builder)
		}
		builder.WriteString("</")
		builder.WriteString(n.name)
		builder.WriteString(">")
	}
}

func (n nodeRenderer) renderAttribute(builder *strings.Builder) {
	if n.value == nil {
		builder.WriteString(n.name)
		return
	}
	value := n.createStringValue()
	if strings.HasPrefix(value, openBracket) && strings.HasSuffix(value, closeBracket) {
		value = strings.ReplaceAll(value, `"`, `'`)
	}
	builder.WriteString(n.name)
	builder.WriteString(`="`)
	builder.WriteString(value)
	builder.WriteString(`"`)
}

func (n nodeRenderer) renderAttributes(builder *strings.Builder) {
	size := len(n.attributes)
	if size == 0 {
		return
	}
	for i, a := range n.attributes {
		nodeRenderer{a}.renderAttribute(builder)
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
}

func (n nodeRenderer) renderChildren(builder *strings.Builder) {
	for _, ch := range n.children {
		switch ch.nodeType {
		case nodeRaw:
			builder.WriteString(nodeRenderer{ch}.renderRaw())
		case nodeElement:
			nodeRenderer{ch}.renderElement(builder)
		case nodeAttribute:
			nodeRenderer{ch}.renderAttribute(builder)
		case nodeText:
			builder.WriteString(nodeRenderer{ch}.renderText())
		}
	}
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
