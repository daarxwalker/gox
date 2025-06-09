package gox

import "testing"

func BenchmarkFlatRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(
			Div(
				Div(Text("Div 1")),
				Div(Text("Div 2")),
				Div(Text("Div 3")),
				Div(Text("Div 4")),
				Div(Text("Div 5")),
			),
		)
	}
}

func TestFlatRender(t *testing.T) {
	Render(
		Div(
			Div(Text("Div 1")),
			Div(Text("Div 2")),
			Div(Text("Div 3")),
			Div(Text("Div 4")),
			Div(Text("Div 5")),
		),
	)
}
