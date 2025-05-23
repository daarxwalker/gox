package htmx

import (
	"encoding/json"
	"net/http"

	"github.com/daarxwalker/gox"
)

const (
	IgnoreStatus = http.StatusNoContent
)

const (
	RequestHeaderBoosted               = "HX-Boosted"
	RequestHeaderCurrentUrl            = "HX-Current-URL"
	RequestHeaderHistoryRestoreRequest = "HX-History-Restore-Request"
	RequestHeaderPrompt                = "HX-Prompt"
	RequestHeaderRequest               = "HX-Request"
	RequestHeaderTarget                = "HX-Target"
	RequestHeaderTriggerName           = "HX-Trigger-Name"
	RequestHeaderTrigger               = "HX-Trigger"
)

func Get(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-get")(value...)
}

func Post(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-post")(value...)
}

func Put(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-put")(value...)
}

func Patch(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-patch")(value...)
}

func Delete(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-delete")(value...)
}

func Confirm(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-confirm")(value...)
}

func Encoding(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-encoding")(value...)
}

func Headers(value ...map[string]any) gox.Node {
	var v string
	if len(value) > 0 {
		valueBytes, err := json.Marshal(value[0])
		if err == nil {
			v = string(valueBytes)
		}
	}
	return gox.CreateAttribute[string](atrributePrefix + "-headers")(v)
}

func Include(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-include")(value...)
}

func Params(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-params")(value...)
}

func Prompt(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-prompt")(value...)
}

func Request(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-request")(value...)
}

func Validate(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-validate")(value...)
}

func Vals(value ...string) gox.Node {
	return gox.CreateAttribute[string](atrributePrefix + "-vals")(value...)
}
