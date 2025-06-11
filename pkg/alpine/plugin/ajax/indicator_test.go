package ajax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daarxwalker/gox"
)

func TestIndicator(t *testing.T) {
	t.Run(
		"render aria busy", func(t *testing.T) {
			assert.Equal(
				t,
				`aria-busy="true"`,
				gox.Render(Indicator(true)),
			)
		},
	)
}
