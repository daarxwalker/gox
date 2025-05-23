package gox

func Fragment(nodes ...Node) Node {
	attributes, children := processNodes(nodes)
	return node{
		nodeType:   nodeFragment,
		children:   children,
		attributes: attributes,
	}
}
