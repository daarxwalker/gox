package gox

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestRenderElements(t *testing.T) {
	t.Run(
		"render div", func(t *testing.T) {
			assert.Equal(t, "<div></div>", Render(Div()))
		},
	)
	t.Run(
		"render button with attributes and text", func(t *testing.T) {
			assert.Equal(
				t,
				`<button type="submit">Test</button>`,
				Render(Button(Type("submit"), Text("Test"))),
			)
		},
	)
	t.Run(
		"render page", func(t *testing.T) {
			page :=
				Fragment(
					Html(
						Head(),
						Body(),
					),
				)
			assert.Equal(
				t,
				`<!DOCTYPE html><html><head></head><body></body></html>`,
				Render(page),
			)
		},
	)
	t.Run(
		"render nested", func(t *testing.T) {
			nested :=
				Fragment(
					Div(
						Class("test class"),
						Div(
							Text("test text"),
						),
					),
				)
			assert.Equal(
				t,
				`<div class="test class"><div>test text</div></div>`,
				Render(nested),
			)
		},
	)
}
