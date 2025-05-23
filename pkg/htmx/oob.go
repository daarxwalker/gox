package hx

import "github.com/daarxwalker/gox"

func SelectOob(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-select-oob")(value...)
}

func SwapOob(value ...string) gox.Node {
	if len(value) == 0 {
		value = append(value, "true")
	}
	return gox.CreateAttribute[string](atrributePrefix + "-swap-oob")(value...)
}
