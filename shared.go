package gox

import (
	"fmt"
	"strings"
)

const (
	sharedCite     = "cite"
	sharedClipPath = "clipPath"
	sharedData     = "data"
	sharedFilter   = "filter"
	sharedForm     = "form"
	sharedLabel    = "label"
	sharedMask     = "mask"
	sharedPath     = "path"
	sharedPattern  = "pattern"
	sharedSlot     = "slot"
	sharedSpan     = "span"
	sharedStyle    = "style"
	sharedTitle    = "title"
)

func Cite(nodes ...Node) Node {
	return createShared(sharedCite, nodeElement, nodes...)
}

func ClipPath(nodes ...Node) Node {
	return createShared(sharedClipPath, nodeAttribute, nodes...)
}

func Data(nodes ...Node) Node {
	builder := new(strings.Builder)
	builder.WriteString(sharedData)
	nameNode := findNodeWithName(attributeName, nodes...)
	name := fmt.Sprint(nameNode.value)
	if nameNode.value != nil && len(name) > 0 {
		builder.WriteString("-")
		builder.WriteString(name)
	}
	nodes = removeNodesWithName(attributeName, nodes...)
	return createShared(builder.String(), nodeAttribute, nodes...)
}

func Filter(nodes ...Node) Node {
	return createShared(sharedFilter, nodeAttribute, nodes...)
}

func Form(nodes ...Node) Node {
	return createShared(sharedForm, nodeElement, nodes...)
}

func Label(nodes ...Node) Node {
	return createShared(sharedLabel, nodeElement, nodes...)
}

func Mask(nodes ...Node) Node {
	return createShared(sharedMask, nodeElement, nodes...)
}

func Path(nodes ...Node) Node {
	return createShared(sharedPath, nodeElement, nodes...)
}

func Pattern(nodes ...Node) Node {
	return createShared(sharedPattern, nodeElement, nodes...)
}

func Slot(nodes ...Node) Node {
	return createShared(sharedSlot, nodeElement, nodes...)
}

func Span(nodes ...Node) Node {
	return createShared(sharedSpan, nodeElement, nodes...)
}

func Style(nodes ...Node) Node {
	return createShared(sharedStyle, nodeAttribute, nodes...)
}

func Title(nodes ...Node) Node {
	return createShared(sharedTitle, nodeElement, nodes...)
}
