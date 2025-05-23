package htmx

import "github.com/daarxwalker/gox"

const (
	EventAbort                 = "htmx:abort"
	EventAfterOnLoad           = "htmx:afterOnLoad"
	EventAfterProcessNode      = "htmx:afterProcessNode"
	EventAfterRequest          = "htmx:afterRequest"
	EventAfterSettle           = "htmx:afterSettle"
	EventAfterSwap             = "htmx:afterSwap"
	EventBeforeCleanupElement  = "htmx:beforeCleanupElement"
	EventBeforeOnLoad          = "htmx:beforeOnLoad"
	EventBeforeProcessNode     = "htmx:beforeProcessNode"
	EventBeforeRequest         = "htmx:beforeRequest"
	EventBeforeSwap            = "htmx:beforeSwap"
	EventBeforeSend            = "htmx:beforeSend"
	EventConfigRequest         = "htmx:configRequest"
	EventConfirm               = "htmx:confirm"
	EventHistoryCacheError     = "htmx:historyCacheError"
	EventHistoryCacheMiss      = "htmx:historyCacheMiss"
	EventHistoryCacheMissError = "htmx:historyCacheMissError"
	EventHistoryCacheMissLoad  = "htmx:historyCacheMissLoad"
	EventHistoryRestore        = "htmx:historyRestore"
	EventBeforeHistorySave     = "htmx:beforeHistorySave"
	EventLoad                  = "htmx:load"
	EventNoSSESourceError      = "htmx:noSSESourceError"
	EventOnLoadError           = "htmx:onLoadError"
	EventOobBeforeSwap         = "htmx:oobBeforeSwap"
	EventOobErrorNoTarget      = "htmx:oobErrorNoTarget"
	EventPrompt                = "htmx:prompt"
	EventPushedIntoHistory     = "htmx:pushedIntoHistory"
	EventResponseError         = "htmx:responseError"
	EventSendError             = "htmx:sendError"
	EventSseError              = "htmx:sseError"
	EventSseOpen               = "htmx:sseOpen"
	EventSwapError             = "htmx:swapError"
	EventTargetError           = "htmx:targetError"
	EventTimeout               = "htmx:timeout"
	EventValidationValidate    = "htmx:validation:validate"
	EventValidationFailed      = "htmx:validation:failed"
	EventValidationHalted      = "htmx:validation:halted"
	EventXhrAbort              = "htmx:xhr:abort"
	EventXhrLoadEnd            = "htmx:xhr:loadend"
	EventXhrLoadStart          = "htmx:xhr:loadstart"
	EventXhrProgress           = "htmx:xhr:progress"
)

const (
	EventScrollTop    = "scroll:top"
	EventScrollBottom = "scroll:bottom"
	EventWindowTop    = "window:top"
	EventWindowBottom = "window:bottom"
)

func On(event string, value ...string) gox.Node {
	if len(event) == 0 {
		return gox.CreateAttribute[string](atrributePrefix + "-on")(value...)
	}
	return gox.CreateAttribute[string](atrributePrefix + "-on:" + event)(value...)
}
