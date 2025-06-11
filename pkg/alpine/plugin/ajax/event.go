package ajax

import "github.com/daarxwalker/gox"

const (
	EventBefore   = "ajax:before"
	EventSend     = "ajax:send"
	EventRedirect = "ajax:redirect"
	EventSuccess  = "ajax:success"
	EventError    = "ajax:error"
	EventSent     = "ajax:sent"
	EventMissing  = "ajax:missing"
	EventMerge    = "ajax:merge"
	EventMerged   = "ajax:merged"
	EventAfter    = "ajax:after"
)

func OnBefore(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventBefore)(script)
}

func OnSend(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventSend)(script)
}

func OnRedirect(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventRedirect)(script)
}

func OnSuccess(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventSuccess)(script)
}

func OnError(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventError)(script)
}

func OnSent(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventSent)(script)
}

func OnMissing(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventMissing)(script)
}

func OnMerge(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventMerge)(script)
}

func OnMerged(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventMerged)(script)
}

func OnAfter(script string) gox.Node {
	return gox.CreateAttribute[string]("@" + EventAfter)(script)
}
