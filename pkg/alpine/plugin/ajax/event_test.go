package ajax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daarxwalker/gox"
)

func TestEvent(t *testing.T) {
	t.Run(
		"render ajax before event", func(t *testing.T) {
			assert.Equal(
				t,
				`@ajax:before="console.log('test')"`,
				gox.Render(OnBefore("console.log('test')")),
			)
		},
	)
}
