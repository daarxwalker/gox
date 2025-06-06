package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributeEvent(t *testing.T) {
	t.Run(
		"render onclick", func(t *testing.T) {
			assert.Equal(t, `onclick="test.showModal()"`, Render(OnClick("test.showModal()")))
		},
	)
}
