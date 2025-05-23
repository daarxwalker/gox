package gox

import (
	"strings"
)

type Node interface {
	Node() Node
}

type node struct {
	nodeType   int
	name       string
	value      any
	attributes []node
	children   []node
}

const (
	nodeElement = iota
	nodeAttribute
	nodeText
	nodeFragment
	nodeRaw
	nodeModifier
)

func (n node) Node() Node {
	return n
}

func createElement(name string, nodes ...Node) Node {
	attributes, children := processNodes(nodes)
	return node{
		nodeType:   nodeElement,
		name:       name,
		attributes: attributes,
		children:   children,
	}
}

func createAttribute[T any](name string, value ...T) Node {
	n := node{
		nodeType: nodeAttribute,
		name:     name,
	}
	if len(value) > 0 {
		n.value = value[0]
	}
	return n
}

func createShared(name string, nodeType int, nodes ...Node) Node {
	if len(nodes) > 0 {
		modifier := findNodeWithType(nodeModifier, nodes...)
		if modifier.nodeType == nodeModifier && modifier.value != nil {
			nodeType = modifier.value.(int)
			nodes = removeNodesWithType(nodeModifier, nodes...)
		}
	}
	if nodeType == nodeAttribute {
		if len(nodes) == 0 {
			return createAttribute[any](name)
		}
		values := make([]string, 0)
		for _, item := range nodes {
			n, ok := item.(node)
			if !ok {
				continue
			}
			switch n.nodeType {
			case nodeFragment:
				values = append(values, getStringNodesValues(n.children...)...)
			default:
				switch v := n.value.(type) {
				case string:
					values = append(values, v)
				}
			}
		}
		if len(values) == 0 {
			return createAttribute[any](name)
		}
		if len(values) == 1 {
			return createAttribute(name, values[0])
		}
		return createAttribute(name, strings.Join(values, " "))
	}
	return createElement(name, nodes...)
}
