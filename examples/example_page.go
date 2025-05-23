package examples

import . "github.com/daarxwalker/gox"

func createPage(lang, title string, children ...Node) Node {
	page :=
		Html(
			Lang(lang),
			Head(
				Title(Text(title)),
				Meta(Name("viewport"), Content("width=device-width,initial-scale=1.0")),
			),
			Body(children...),
		)
	return page
}
