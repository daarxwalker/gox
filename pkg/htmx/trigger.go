package htmx

import "github.com/daarxwalker/gox"

func Trigger(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-trigger")(value...)
}
