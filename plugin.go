package gox

import (
	"fmt"
	"strings"
)

// CLSX plugin
// Conditional classes rendering
type CLSX map[any]bool

func (c CLSX) Node() Node {
	return Class(c.String())
}

func (c CLSX) Merge(items ...CLSX) CLSX {
	for _, item := range items {
		for k, v := range item {
			c[k] = v
		}
	}
	return c
}

func (c CLSX) String() string {
	result := make([]string, 0)
	for class, use := range c {
		if !use {
			continue
		}
		switch v := class.(type) {
		case string:
			result = append(result, v)
		case fmt.Stringer:
			result = append(result, v.String())
		default:
			result = append(result, fmt.Sprint(v))
		}
	}
	return strings.Join(result, " ")
}
