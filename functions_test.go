package gox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctions(t *testing.T) {
	t.Run(
		"func render", func(t *testing.T) {
			styleElementValue := ".h1{background:blue;}"
			styleAttributeValue := "background:blue;"
			assert.Equal(
				t,
				fmt.Sprintf(`style="%s"<div></div>Test<style>%s</style>`, styleAttributeValue, styleElementValue),
				Render(
					Style(Attribute(), Text(styleAttributeValue)),
					Div(),
					Text("Test"),
					Style(Element(), Text(styleElementValue)),
				),
			)
		},
	)
	t.Run(
		"func write", func(t *testing.T) {
			w := new(strings.Builder)
			err := Write(w, Div(Text("test")))
			assert.Equal(t, nil, err)
			assert.Equal(t, `<div>test</div>`, w.String())
		},
	)
}
