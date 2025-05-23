package gox

func Range[T any](s []T, fn func(item T, index int) Node) Node {
	children := make([]node, 0)
	for i, val := range s {
		item := fn(val, i)
		n, ok := item.(node)
		if !ok {
			continue
		}
		if n.nodeType == nodeFragment {
			children = append(children, n.children...)
			continue
		}
		children = append(children, n)
	}
	return node{
		nodeType: nodeFragment,
		children: children,
	}
}

func MapRange[K comparable, V any](m map[K]V, fn func(key K, value V) Node) Node {
	children := make([]node, 0)
	for k, v := range m {
		item := fn(k, v)
		n, ok := item.(node)
		if !ok {
			continue
		}
		if n.nodeType == nodeFragment {
			children = append(children, n.children...)
			continue
		}
		children = append(children, n)
	}
	return node{
		nodeType: nodeFragment,
		children: children,
	}
}

func If(condition bool, nodes ...Node) Node {
	if !condition {
		return node{
			nodeType: nodeFragment,
		}
	}
	attributes, children := processNodes(nodes)
	return node{
		nodeType:   nodeFragment,
		children:   children,
		attributes: attributes,
	}
}
