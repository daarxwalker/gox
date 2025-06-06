package gox

const (
	attributeOnAfterprint     = "onafterprint"
	attributeOnBeforeprint    = "onbeforeprint"
	attributeOnBeforeunload   = "onbeforeunload"
	attributeOnError          = "onerror"
	attributeOnHashchange     = "onhashchange"
	attributeOnLoad           = "onload"
	attributeOnMessage        = "onmessage"
	attributeOnOffline        = "onoffline"
	attributeOnOnline         = "ononline"
	attributeOnPagehide       = "onpagehide"
	attributeOnPageshow       = "onpageshow"
	attributeOnPopstate       = "onpopstate"
	attributeOnResize         = "onresize"
	attributeOnStorage        = "onstorage"
	attributeOnUnload         = "onunload"
	attributeOnBlur           = "onblur"
	attributeOnChange         = "onchange"
	attributeOnContextmenu    = "oncontextmenu"
	attributeOnFocus          = "onfocus"
	attributeOnInput          = "oninput"
	attributeOnInvalid        = "oninvalid"
	attributeOnReset          = "onreset"
	attributeOnSearch         = "onsearch"
	attributeOnSelect         = "onselect"
	attributeOnSubmit         = "onsubmit"
	attributeOnKeydown        = "onkeydown"
	attributeOnKeypress       = "onkeypress"
	attributeOnKeyup          = "onkeyup"
	attributeOnClick          = "onclick"
	attributeOnDblclick       = "ondblclick"
	attributeOnMousedown      = "onmousedown"
	attributeOnMouseup        = "onmouseup"
	attributeOnMousemove      = "onmousemove"
	attributeOnMouseover      = "onmouseover"
	attributeOnMouseout       = "onmouseout"
	attributeOnMouseenter     = "onmouseenter"
	attributeOnMouseleave     = "onmouseleave"
	attributeOnDrag           = "ondrag"
	attributeOnDragend        = "ondragend"
	attributeOnDragenter      = "ondragenter"
	attributeOnDragleave      = "ondragleave"
	attributeOnDragover       = "ondragover"
	attributeOnDragstart      = "ondragstart"
	attributeOnDrop           = "ondrop"
	attributeOnCopy           = "oncopy"
	attributeOnCut            = "oncut"
	attributeOnPaste          = "onpaste"
	attributeOnAbort          = "onabort"
	attributeOnCanplay        = "oncanplay"
	attributeOnCanplaythrough = "oncanplaythrough"
	attributeOnCuechange      = "oncuechange"
	attributeOnDurationchange = "ondurationchange"
	attributeOnEmptied        = "onemptied"
	attributeOnEnded          = "onended"
	attributeOnLoadeddata     = "onloadeddata"
	attributeOnLoadedmetadata = "onloadedmetadata"
	attributeOnLoadstart      = "onloadstart"
	attributeOnPause          = "onpause"
	attributeOnPlay           = "onplay"
	attributeOnPlaying        = "onplaying"
	attributeOnProgress       = "onprogress"
	attributeOnRatechange     = "onratechange"
	attributeOnSeeked         = "onseeked"
	attributeOnSeeking        = "onseeking"
	attributeOnStalled        = "onstalled"
	attributeOnSuspend        = "onsuspend"
	attributeOnTimeupdate     = "ontimeupdate"
	attributeOnVolumechange   = "onvolumechange"
	attributeOnWaiting        = "onwaiting"
	attributeOnToggle         = "ontoggle"
)

func OnAfterprint(values ...string) Node {
	return createAttribute(attributeOnAfterprint, values...)
}

func OnBeforeprint(values ...string) Node {
	return createAttribute(attributeOnBeforeprint, values...)
}

func OnBeforeunload(values ...string) Node {
	return createAttribute(attributeOnBeforeunload, values...)
}

func OnError(values ...string) Node {
	return createAttribute(attributeOnError, values...)
}

func OnHashchange(values ...string) Node {
	return createAttribute(attributeOnHashchange, values...)
}

func OnLoad(values ...string) Node {
	return createAttribute(attributeOnLoad, values...)
}

func OnMessage(values ...string) Node {
	return createAttribute(attributeOnMessage, values...)
}

func OnOffline(values ...string) Node {
	return createAttribute(attributeOnOffline, values...)
}

func OnOnline(values ...string) Node {
	return createAttribute(attributeOnOnline, values...)
}

func OnPagehide(values ...string) Node {
	return createAttribute(attributeOnPagehide, values...)
}

func OnPageshow(values ...string) Node {
	return createAttribute(attributeOnPageshow, values...)
}

func OnPopstate(values ...string) Node {
	return createAttribute(attributeOnPopstate, values...)
}

func OnResize(values ...string) Node {
	return createAttribute(attributeOnResize, values...)
}

func OnStorage(values ...string) Node {
	return createAttribute(attributeOnStorage, values...)
}

func OnUnload(values ...string) Node {
	return createAttribute(attributeOnUnload, values...)
}

func OnBlur(values ...string) Node {
	return createAttribute(attributeOnBlur, values...)
}

func OnChange(values ...string) Node {
	return createAttribute(attributeOnChange, values...)
}

func OnContextmenu(values ...string) Node {
	return createAttribute(attributeOnContextmenu, values...)
}

func OnFocus(values ...string) Node {
	return createAttribute(attributeOnFocus, values...)
}

func OnInput(values ...string) Node {
	return createAttribute(attributeOnInput, values...)
}

func OnInvalid(values ...string) Node {
	return createAttribute(attributeOnInvalid, values...)
}

func OnReset(values ...string) Node {
	return createAttribute(attributeOnReset, values...)
}

func OnSearch(values ...string) Node {
	return createAttribute(attributeOnSearch, values...)
}

func OnSelect(values ...string) Node {
	return createAttribute(attributeOnSelect, values...)
}

func OnSubmit(values ...string) Node {
	return createAttribute(attributeOnSubmit, values...)
}

func OnKeydown(values ...string) Node {
	return createAttribute(attributeOnKeydown, values...)
}

func OnKeypress(values ...string) Node {
	return createAttribute(attributeOnKeypress, values...)
}

func OnKeyup(values ...string) Node {
	return createAttribute(attributeOnKeyup, values...)
}

func OnClick(values ...string) Node {
	return createAttribute(attributeOnClick, values...)
}

func OnDblclick(values ...string) Node {
	return createAttribute(attributeOnDblclick, values...)
}

func OnMousedown(values ...string) Node {
	return createAttribute(attributeOnMousedown, values...)
}

func OnMouseup(values ...string) Node {
	return createAttribute(attributeOnMouseup, values...)
}

func OnMousemove(values ...string) Node {
	return createAttribute(attributeOnMousemove, values...)
}

func OnMouseover(values ...string) Node {
	return createAttribute(attributeOnMouseover, values...)
}

func OnMouseout(values ...string) Node {
	return createAttribute(attributeOnMouseout, values...)
}

func OnMouseenter(values ...string) Node {
	return createAttribute(attributeOnMouseenter, values...)
}

func OnMouseleave(values ...string) Node {
	return createAttribute(attributeOnMouseleave, values...)
}

func OnDrag(values ...string) Node {
	return createAttribute(attributeOnDrag, values...)
}

func OnDragend(values ...string) Node {
	return createAttribute(attributeOnDragend, values...)
}

func OnDragenter(values ...string) Node {
	return createAttribute(attributeOnDragenter, values...)
}

func OnDragleave(values ...string) Node {
	return createAttribute(attributeOnDragleave, values...)
}

func OnDragover(values ...string) Node {
	return createAttribute(attributeOnDragover, values...)
}

func OnDragstart(values ...string) Node {
	return createAttribute(attributeOnDragstart, values...)
}

func OnDrop(values ...string) Node {
	return createAttribute(attributeOnDrop, values...)
}

func OnCopy(values ...string) Node {
	return createAttribute(attributeOnCopy, values...)
}

func OnCut(values ...string) Node {
	return createAttribute(attributeOnCut, values...)
}

func OnPaste(values ...string) Node {
	return createAttribute(attributeOnPaste, values...)
}

func OnAbort(values ...string) Node {
	return createAttribute(attributeOnAbort, values...)
}

func OnCanplay(values ...string) Node {
	return createAttribute(attributeOnCanplay, values...)
}

func OnCanplaythrough(values ...string) Node {
	return createAttribute(attributeOnCanplaythrough, values...)
}

func OnCuechange(values ...string) Node {
	return createAttribute(attributeOnCuechange, values...)
}

func OnDurationchange(values ...string) Node {
	return createAttribute(attributeOnDurationchange, values...)
}

func OnEmptied(values ...string) Node {
	return createAttribute(attributeOnEmptied, values...)
}

func OnEnded(values ...string) Node {
	return createAttribute(attributeOnEnded, values...)
}

func OnLoadeddata(values ...string) Node {
	return createAttribute(attributeOnLoadeddata, values...)
}

func OnLoadedmetadata(values ...string) Node {
	return createAttribute(attributeOnLoadedmetadata, values...)
}

func OnLoadstart(values ...string) Node {
	return createAttribute(attributeOnLoadstart, values...)
}

func OnPause(values ...string) Node {
	return createAttribute(attributeOnPause, values...)
}

func OnPlay(values ...string) Node {
	return createAttribute(attributeOnPlay, values...)
}

func OnPlaying(values ...string) Node {
	return createAttribute(attributeOnPlaying, values...)
}

func OnProgress(values ...string) Node {
	return createAttribute(attributeOnProgress, values...)
}

func OnRatechange(values ...string) Node {
	return createAttribute(attributeOnRatechange, values...)
}

func OnSeeked(values ...string) Node {
	return createAttribute(attributeOnSeeked, values...)
}

func OnSeeking(values ...string) Node {
	return createAttribute(attributeOnSeeking, values...)
}

func OnStalled(values ...string) Node {
	return createAttribute(attributeOnStalled, values...)
}

func OnSuspend(values ...string) Node {
	return createAttribute(attributeOnSuspend, values...)
}

func OnTimeupdate(values ...string) Node {
	return createAttribute(attributeOnTimeupdate, values...)
}

func OnVolumechange(values ...string) Node {
	return createAttribute(attributeOnVolumechange, values...)
}

func OnWaiting(values ...string) Node {
	return createAttribute(attributeOnWaiting, values...)
}

func OnToggle(values ...string) Node {
	return createAttribute(attributeOnToggle, values...)
}
