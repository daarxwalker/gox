package ajax

import "github.com/daarxwalker/gox"

const (
	ariaBusyAttribute = "aria-busy"
)

func Indicator(show bool) gox.Node {
	return gox.CreateAttribute[bool](ariaBusyAttribute)(show)
}
