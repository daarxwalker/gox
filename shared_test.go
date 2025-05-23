package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharedNodes(t *testing.T) {
	t.Run(
		"render shared element", func(t *testing.T) {
			styleEl := Style(Element())
			assert.Equal(t, "<style></style>", Render(styleEl))
		},
	)
	t.Run(
		"render shared attribute", func(t *testing.T) {
			styleValue := `background:blue;`
			styleAttr := Style(Text(styleValue))
			assert.Equal(t, fmt.Sprintf(`style="%s"`, styleValue), Render(styleAttr))
		},
	)
	t.Run(
		"render data attribute", func(t *testing.T) {
			dataItem := Data(Name("item"), Text("test"))
			assert.Equal(t, `data-item="test"`, Render(dataItem))
		},
	)
	t.Run(
		"render data element", func(t *testing.T) {
			dataItem := Data(Element(), Value(123), Text("Test"))
			assert.Equal(t, `<data value="123">Test</data>`, Render(dataItem))
		},
	)
}
