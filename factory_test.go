package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	t.Run(
		"create and render custom attribute", func(t *testing.T) {
			customAttr := CreateAttribute[string]("test-attribute")
			assert.Equal(t, `test-attribute="test"`, Render(customAttr("test")))
		},
	)
	t.Run(
		"create and render custom element", func(t *testing.T) {
			customEl := CreateElement("test-element")
			assert.Equal(t, "<test-element></test-element>", Render(customEl()))
		},
	)
}
