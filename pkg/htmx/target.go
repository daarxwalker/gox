package htmx

import "github.com/daarxwalker/gox"

func Target(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-target")(value...)
}
