package ajax

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/daarxwalker/gox"
)

func TestModifier(t *testing.T) {
	t.Run(
		"render target with status code", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target.4xx="test-form"`,
				gox.Render(Target("test-form", "4xx")),
			)
		},
	)
	t.Run(
		"render target with error", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target.error="test-form"`,
				gox.Render(Target("test-form", ERROR)),
			)
		},
	)
	t.Run(
		"render value with alias", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target="test:test-alias"`,
				gox.Render(Target("test", Alias("test-alias"))),
			)
		},
	)
	t.Run(
		"render target with special target", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target.away="_top"`,
				gox.Render(Target(TOP, AWAY)),
			)
		},
	)
	t.Run(
		"render dynamic target", func(t *testing.T) {
			assert.Equal(
				t,
				`x-target:dynamic="'test_'+test.id"`,
				gox.Render(Target("'test_'+test.id", DYNAMIC)),
			)
		},
	)
}
