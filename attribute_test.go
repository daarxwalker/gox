package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute(t *testing.T) {
	t.Run(
		"render class", func(t *testing.T) {
			classes := "test-1 test-2"
			class := Class("test-1 test-2")
			assert.Equal(t, fmt.Sprintf(`class="%s"`, classes), Render(class))
		},
	)
	t.Run(
		"render style", func(t *testing.T) {
			styles := "background:blue;font-weight:700;color:white;"
			style := Style(Raw(styles))
			assert.Equal(t, fmt.Sprintf(`style="%s"`, styles), Render(style))
		},
	)
	t.Run(
		"render checked", func(t *testing.T) {
			assert.Equal(t, `checked`, Render(Checked()))
			assert.Equal(t, `checked="true"`, Render(Checked(true)))
			assert.Equal(t, `checked="false"`, Render(Checked(false)))
		},
	)
	t.Run(
		"render selected", func(t *testing.T) {
			assert.Equal(t, `selected`, Render(Selected()))
			assert.Equal(t, `selected="true"`, Render(Selected(true)))
			assert.Equal(t, `selected="false"`, Render(Selected(false)))
		},
	)
	t.Run(
		"render on event", func(t *testing.T) {
			assert.Equal(t, `onclick="test.showModal()"`, Render(On("click", "test.showModal()")))
		},
	)
}
