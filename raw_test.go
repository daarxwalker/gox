package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawNodes(t *testing.T) {
	rawHtml := `<span class="test"></span>`
	t.Run(
		"render raw", func(t *testing.T) {
			assert.Equal(t, fmt.Sprintf("<div>%s</div>", rawHtml), Render(Div(Raw(rawHtml))))
		},
	)
}
