package ajax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daarxwalker/gox"
)

func TestDirective(t *testing.T) {
	t.Run(
		"render merge with transition", func(t *testing.T) {
			assert.Equal(
				t,
				`x-merge.transition="replace"`,
				gox.Render(Merge(REPLACE, TRANSITION)),
			)
		},
	)
	t.Run(
		"render targets", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target="test1 test2"`,
				gox.Render(Targets([]string{"test1", "test2"})),
			)
		},
	)
	t.Run(
		"render autofocus", func(t *testing.T) {
			assert.Equal(
				t,
				`x-autofocus`,
				gox.Render(AutoFocus()),
			)
		},
	)
	t.Run(
		"render headers", func(t *testing.T) {
			assert.Equal(
				t,
				`x-headers="{'test1':1,'test2':true}"`,
				gox.Render(Headers(Map{"test1": 1, "test2": true})),
			)
		},
	)
	t.Run(
		"render sync", func(t *testing.T) {
			assert.Equal(
				t,
				`x-sync`,
				gox.Render(Sync()),
			)
		},
	)
}
