package gox

func Append(n Node, nodes ...Node) Node {
	attributes, children := processNodes(nodes)
	updatedNode := n.(node)
	updatedNode.attributes = append(updatedNode.attributes, attributes...)
	updatedNode.children = append(updatedNode.children, children...)
	return updatedNode
}
