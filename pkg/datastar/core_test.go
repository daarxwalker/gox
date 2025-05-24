package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarxwalker/gox"
)

func TestCore(t *testing.T) {
	t.Run(
		"render signal", func(t *testing.T) {
			assert.Equal(
				t,
				`data-signals-test="true"`,
				gox.Render(Signal("test", true)),
			)
		},
	)
	t.Run(
		"render signals map", func(t *testing.T) {
			assert.Equal(
				t,
				`data-signals="{'test1':true,'test2':'test string','test3':156}"`,
				gox.Render(Signals(Map{"test1": true, "test2": "test string", "test3": 156})),
			)
		},
	)
	t.Run(
		"render signal with modifier", func(t *testing.T) {
			assert.Equal(
				t,
				`data-signals-test__case.kebab="1"`,
				gox.Render(Signal("test", 1, WithCase(KEBAB))),
			)
		},
	)
	t.Run(
		"render computed", func(t *testing.T) {
			assert.Equal(
				t,
				`data-computed-test3="$test1 + $test2"`,
				gox.Render(Computed("test3", "$test1 + $test2")),
			)
		},
	)
	t.Run(
		"render computed with modifier", func(t *testing.T) {
			assert.Equal(
				t,
				`data-computed-test3__case.kebab="$test1 + $test2"`,
				gox.Render(Computed("test3", "$test1 + $test2", WithCase(KEBAB))),
			)
		},
	)
	t.Run(
		"render star ignore", func(t *testing.T) {
			assert.Equal(
				t,
				`data-star-ignore`,
				gox.Render(StarIgnore()),
			)
		},
	)
}
