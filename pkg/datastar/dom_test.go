package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarxwalker/gox"
)

func TestDOM(t *testing.T) {
	t.Run(
		"render attrs", func(t *testing.T) {
			assert.Equal(
				t,
				`data-attr-test="$condition"`,
				gox.Render(Attr("test", "$condition")),
			)
		},
	)
	t.Run(
		"render attrs", func(t *testing.T) {
			assert.Equal(
				t,
				`data-attr="{'test1':$state1,'test2':$state2}"`,
				gox.Render(Attrs(Map{"test1": "$state1", "test2": "$state2"})),
			)
		},
	)
	t.Run(
		"render bind", func(t *testing.T) {
			assert.Equal(
				t,
				`data-bind-test`,
				gox.Render(Bind("test")),
			)
		},
	)
	t.Run(
		"render bind with modifiers", func(t *testing.T) {
			assert.Equal(
				t,
				`data-bind-test__case.kebab`,
				gox.Render(Bind("test", WithCase(KEBAB))),
			)
		},
	)
	t.Run(
		"render class", func(t *testing.T) {
			assert.Equal(
				t,
				`data-class-test="$condition"`,
				gox.Render(Class("test", "$condition")),
			)
		},
	)
	t.Run(
		"render classes", func(t *testing.T) {
			assert.Equal(
				t,
				`data-class="{'test1':$state1,'test2':$state2}"`,
				gox.Render(Classes(Map{"test1": "$state1", "test2": "$state2"})),
			)
		},
	)
	t.Run(
		"render event", func(t *testing.T) {
			assert.Equal(
				t,
				`data-on-click__prevent="$test = ''"`,
				gox.Render(On(CLICK, "$test = ''", WithPrevent())),
			)
		},
	)
	t.Run(
		"render ref", func(t *testing.T) {
			assert.Equal(
				t,
				`data-ref-test__case.kebab`,
				gox.Render(Ref("test", WithCase(KEBAB))),
			)
		},
	)
	t.Run(
		"render show", func(t *testing.T) {
			assert.Equal(
				t,
				`data-show="$condition"`,
				gox.Render(Show("$condition")),
			)
		},
	)
	t.Run(
		"render text", func(t *testing.T) {
			assert.Equal(
				t,
				`data-text="$value"`,
				gox.Render(Text("$value")),
			)
		},
	)
}
