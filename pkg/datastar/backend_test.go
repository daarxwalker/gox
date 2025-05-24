package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarxwalker/gox"
)

func TestBackend(t *testing.T) {
	t.Run(
		"render indicator", func(t *testing.T) {
			assert.Equal(
				t,
				`data-indicator-test`,
				gox.Render(Indicator("test")),
			)
		},
	)
}
