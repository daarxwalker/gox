package gox

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectiveRange(t *testing.T) {
	data := []int{1, 2}
	t.Run(
		"range", func(t *testing.T) {
			bodyEl := Body(
				Range(
					data, func(item int, _ int) Node {
						return Span(Text(fmt.Sprintf("%d", item)))
					},
				),
			)
			assert.Equal(t, "<body><span>1</span><span>2</span></body>", Render(bodyEl))
		},
	)
}

func TestDirectiveMapRange(t *testing.T) {
	data := map[string]string{"key1": "value1", "key2": "value2"}
	t.Run(
		"range", func(t *testing.T) {
			bodyEl := Body(
				MapRange(
					data, func(k, v string) Node {
						return Span(Text(fmt.Sprintf("%s:%s", k, v)))
					},
				),
			)
			assert.Equal(t, "<body><span>key1:value1</span><span>key2:value2</span></body>", Render(bodyEl))
		},
	)
}

func TestDirectiveIf(t *testing.T) {
	t.Run(
		"if", func(t *testing.T) {
			passed := Body(
				If(true, Div(Text("passed"))),
			)
			notPassed := Body(
				If(false, Div(Text("not passed"))),
			)
			assert.Equal(t, "<body><div>passed</div></body>", Render(passed), "should render div with text")
			assert.Equal(t, "<body></body>", Render(notPassed), "should not render div with text: not passed")
		},
	)
}

func TestDirectiveIfZone(t *testing.T) {
	t.Run(
		"if zone", func(t *testing.T) {
			passed := Body(
				IfZone(
					true, func() Node {
						return Div(Text("passed"))
					},
				),
			)
			notPassed := Body(
				IfZone(
					false, func() Node {
						return Div(Text("not passed"))
					},
				),
			)
			assert.Equal(t, "<body><div>passed</div></body>", Render(passed), "should render div with text")
			assert.Equal(t, "<body></body>", Render(notPassed), "should not render div with text: not passed")
		},
	)
}

func TestDirectiveZone(t *testing.T) {
	t.Run(
		"zone condition true", func(t *testing.T) {
			condition := true
			assert.Equal(
				t,
				"<div><span>ok</span></div>",
				Render(
					Div(
						Zone(
							func() Node {
								if condition {
									return Span(Text("ok"))
								}
								return Fragment()
							},
						),
					),
				),
			)
		},
	)
	t.Run(
		"zone condition false", func(t *testing.T) {
			condition := false
			assert.Equal(
				t,
				"<div></div>",
				Render(
					Div(
						Zone(
							func() Node {
								if condition {
									return Span(Text("ok"))
								}
								return Fragment()
							},
						),
					),
				),
			)
		},
	)
}
