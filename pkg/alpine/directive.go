package alpine

import (
	"encoding/json"
	"strings"

	"github.com/daarxwalker/gox"
)

const (
	dataDirective       = "x-data"
	initDirective       = "x-init"
	showDirective       = "x-show"
	textDirective       = "x-text"
	htmlDirective       = "x-html"
	modelDirective      = "x-model"
	modelableDirective  = "x-model"
	forDirective        = "x-for"
	transitionDirective = "x-transition"
	effectDirective     = "x-effect"
	ignoreDirective     = "x-ignore"
	refDirective        = "x-ref"
	cloakDirective      = "x-cloak"
	teleportDirective   = "x-teleport"
	ifDirective         = "x-if"
	idDirective         = "x-id"
)

func Directive(name, value string) gox.Node {
	return gox.CreateAttribute[string]("x-" + name)(value)
}

func Data(data map[string]any) gox.Node {
	bytes, err := json.Marshal(data)
	if err != nil {
		return gox.CreateAttribute[string](dataDirective)("{}")
	}
	return gox.CreateAttribute[string](dataDirective)(string(bytes))
}

func Init(script string) gox.Node {
	return gox.CreateAttribute[string](initDirective)(script)
}

func Show(condition string) gox.Node {
	return gox.CreateAttribute[string](showDirective)(condition)
}

func Bind(name string, value ...string) gox.Node {
	if len(value) > 0 {
		return gox.CreateAttribute[string](":" + name)(value[0])
	}
	return gox.CreateAttribute[string](":" + name)(name)
}

func Class(value map[string]string) gox.Node {
	result := make([]string, len(value))
	i := 0
	for k, v := range value {
		result[i] = `"` + k + `": ` + v
		i++
	}
	return Bind("class", `{`+strings.Join(result, `,`)+`}`)
}

func Text(value string) gox.Node {
	return gox.CreateAttribute[string](textDirective)(value)
}

func Html(value string) gox.Node {
	return gox.CreateAttribute[string](htmlDirective)(value)
}

func Model(value string) gox.Node {
	return gox.CreateAttribute[string](modelDirective)(value)
}

func Modelable(condition string) gox.Node {
	return gox.CreateAttribute[string](modelableDirective)(condition)
}

func For(items string, item string, index ...string) gox.Node {
	if len(index) > 0 {
		return gox.CreateAttribute[string](forDirective)("(" + item + ", " + index[0] + ") in " + items)
	}
	return gox.CreateAttribute[string](forDirective)(item + " in " + items)
}

func Transition() gox.Node {
	return gox.CreateAttribute[string](transitionDirective)("")
}

func Effect(script string) gox.Node {
	return gox.CreateAttribute[string](effectDirective)(script)
}

func Ignore(script string) gox.Node {
	return gox.CreateAttribute[string](ignoreDirective)(script)
}

func Ref(name string) gox.Node {
	return gox.CreateAttribute[string](refDirective)(name)
}

func Cloak() gox.Node {
	return gox.CreateAttribute[string](cloakDirective)("")
}

func Teleport(selector string) gox.Node {
	return gox.CreateAttribute[string](teleportDirective)(selector)
}

func If(condition string) gox.Node {
	return gox.CreateAttribute[string](ifDirective)(condition)
}

func Id(name string) gox.Node {
	return gox.CreateAttribute[string](idDirective)("['" + name + "']")
}
