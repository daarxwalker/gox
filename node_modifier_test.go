package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifyNodeType(t *testing.T) {
	t.Run(
		"modify default element to attribute", func(t *testing.T) {
			value := "test"
			title := Title(Attribute(), Text(value))
			assert.Equal(t, fmt.Sprintf(`title="%s"`, value), Render(title))
		},
	)
	t.Run(
		"modify default attribute to element", func(t *testing.T) {
			value := ".test{display:block;}"
			style := Style(Element(), Raw(value))
			assert.Equal(t, fmt.Sprintf(`<style>%s</style>`, value), Render(style))
		},
	)
}
