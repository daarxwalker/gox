package gox

import "strings"

func Raw(html ...string) Node {
	if len(html) == 0 {
		return Fragment()
	}
	builder := new(strings.Builder)
	for _, h := range html {
		builder.WriteString(h)
	}
	return &node{
		nodeType: nodeRaw,
		value:    builder.String(),
	}
}
