package datastar

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daarxwalker/gox"
)

func TestModifier(t *testing.T) {
	t.Run(
		"render delay", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-load__delay.1s="test"`,
				gox.Render(On(LOAD, "test", WithDelay("1s"))),
			)
		},
	)
}
