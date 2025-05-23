package gox

import "io"

func Render(nodes ...Node) string {
	attributes, children := processNodes(nodes)
	return nodeRenderer{
		node: node{
			nodeType:   nodeFragment,
			attributes: attributes,
			children:   children,
		},
	}.render()
}

func Write(w io.Writer, nodes ...Node) error {
	_, err := w.Write([]byte(Render(nodes...)))
	return err
}

func MustWrite(w io.Writer, nodes ...Node) {
	if err := Write(w, nodes...); err != nil {
		panic(err)
	}
}
