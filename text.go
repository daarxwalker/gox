package gox

func Text(value any) Node {
	return node{
		nodeType: nodeText,
		value:    value,
	}
}
