package gox

func processNodes(nodes []Node) ([]node, []node) {
	attributes := make([]node, 0)
	children := make([]node, 0)
	for _, item := range nodes {
		if item == nil {
			continue
		}
		n, ok := item.(node)
		if !ok {
			componentNode := item.Node()
			if componentNode == nil {
				continue
			}
			n, ok = componentNode.(node)
			if !ok {
				continue
			}
		}
		switch n.nodeType {
		case nodeFragment:
			for _, an := range n.attributes {
				attributes = append(attributes, an)
			}
			for _, chn := range n.children {
				children = append(children, chn)
			}
		case nodeAttribute:
			attributes = append(attributes, n)
		case nodeElement, nodeText, nodeRaw:
			if n.name == elementHtml {
				children = append(
					children, node{
						nodeType: nodeRaw,
						value:    doctypeHtml,
					},
				)
			}
			children = append(children, n)
		}
	}
	return attributes, children
}
