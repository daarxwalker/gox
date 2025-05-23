package hx

import "github.com/daarxwalker/gox"

func History(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-history")(value...)
}

func HistoryElt(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-history-elt")(value...)
}

func PushUrl() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-push-url")("true")
}

func ReplaceUrl() gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-replace-url")("true")
}
