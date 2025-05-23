package gox

func Attribute() Node {
	return node{
		nodeType: nodeModifier,
		value:    nodeAttribute,
	}
}

func Element() Node {
	return node{
		nodeType: nodeModifier,
		value:    nodeElement,
	}
}
