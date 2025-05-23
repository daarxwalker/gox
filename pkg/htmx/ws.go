package htmx

import "github.com/daarxwalker/gox"

func Ws(url string) gox.Node {
	return gox.Fragment(
		Ext("ws"),
		gox.CreateAttribute[string]("ws-connect")(url),
	)
}
