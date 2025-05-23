package gox

func Text[T any](value T) Node {
	return node{
		nodeType: nodeText,
		value:    value,
	}
}
