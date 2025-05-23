package gox

import "strings"

func Raw(html ...string) Node {
	return node{
		nodeType: nodeRaw,
		value:    strings.Join(html, "\n"),
	}
}
