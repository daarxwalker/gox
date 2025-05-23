package hx

import "github.com/daarxwalker/gox"

func Sync(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-sync")(value...)
}
