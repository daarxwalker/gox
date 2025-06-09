package gox

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestProcessor(t *testing.T) {
	t.Run(
		"filter attributes and children from nodes", func(t *testing.T) {
			nodes := []Node{
				Div(),
				Class("test"),
				Style(Text("test")),
				Text("test"),
				Fragment(
					H1(Text("test")),
					H2(Text("test")),
					H3(Text("test")),
				),
			}
			attributes, children := processNodes(nodes)
			assert.Equal(t, 2, len(attributes))
			assert.Equal(t, 5, len(children))
		},
	)
}

func BenchmarkProcessor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nodes := []Node{
			Div(
				Div(Text("Div 1")),
				Div(Text("Div 2")),
				Div(Text("Div 3")),
				Div(Text("Div 4")),
				Div(Text("Div 5")),
			),
		}
		processNodes(nodes)
	}
}
