package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	t.Run(
		"render text", func(t *testing.T) {
			assert.Equal(t, "test", Render(Text("test")))
		},
	)
	t.Run(
		"render text if value is int", func(t *testing.T) {
			assert.Equal(t, "1", Render(Text(1)))
		},
	)
}
