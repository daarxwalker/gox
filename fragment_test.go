package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFragment(t *testing.T) {
	t.Run(
		"render children wrapped with fragment", func(t *testing.T) {
			fragmentWithChildren := Fragment(
				Div(Text("Test 1")),
				Div(Text("Test 2")),
				Div(Text("Test 3")),
			)
			assert.Equal(t, "<div>Test 1</div><div>Test 2</div><div>Test 3</div>", Render(fragmentWithChildren))
		},
	)
}
