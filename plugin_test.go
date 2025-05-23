package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLSXPlugin(t *testing.T) {
	t.Run(
		"render conditional classes #1", func(t *testing.T) {
			assert.Equal(
				t,
				`<div class="text-blue-400 underline"></div>`,
				Render(
					Div(
						CLSX{
							"text-blue-400 underline": true,
							"text-red-400":            false,
						},
					),
				),
				"should render div with only classes text-blue-400 and underline",
			)
		},
	)
	t.Run(
		"render conditional classes #2", func(t *testing.T) {
			assert.Equal(
				t,
				`<div class="text-blue-400 underline text-red-400"></div>`,
				Render(
					Div(
						CLSX{
							"text-blue-400 underline": true,
							"text-red-400":            true,
						},
					),
				),
				"should render div with all classes",
			)
		},
	)
}
