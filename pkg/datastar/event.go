package datastar

// Core
const (
	DATASTAR_SSE   = "datastar-sse"
	STARTED        = "started"
	FINISHED       = "finished"
	ERROR          = "error"
	RETRYING       = "retrying"
	RETRIES_FAILED = "retries-failed"
)

// Window
const (
	AFTER_PRINT   = "afterprint"
	BEFORE_PRINT  = "beforeprint"
	BEFORE_UNLOAD = "beforeunload"
	HASH_CHANGE   = "hashchange"
	LOAD          = "load"
	MESSAGE       = "message"
	OFFLINE       = "offline"
	ONLINE        = "online"
	PAGE_HIDE     = "pagehide"
	PAGE_SHOW     = "pageshow"
	POP_STATE     = "popstate"
	RESIZE        = "resize"
	STORAGE       = "storage"
	UNLOAD        = "unload"
)

// Form
const (
	BLUR         = "blur"
	CHANGE       = "change"
	CONTEXT_MENU = "contextmenu"
	FOCUS        = "focus"
	INPUT        = "input"
	INVALID      = "invalid"
	RESET        = "reset"
	SEARCH       = "search"
	SELECT       = "select"
	SUBMIT       = "submit"
)

// Keyboard
const (
	KEY_DOWN  = "keydown"
	KEY_PRESS = "keypress"
	KEY_UP    = "keyup"
)

// Mouse
const (
	CLICK       = "click"
	DBL_CLICK   = "dblclick"
	MOUSE_DOWN  = "mousedown"
	MOUSE_MOVE  = "mousemove"
	MOUSE_OUT   = "mouseout"
	MOUSE_OVER  = "mouseover"
	MOUSE_UP    = "mouseup"
	MOUSE_WHEEL = "mousewheel"
	WHEEL       = "wheel"
)

// Drag
const (
	DRAG       = "drag"
	DRAG_END   = "dragend"
	DRAG_ENTER = "dragenter"
	DRAG_LEAVE = "dragleave"
	DRAG_OVER  = "dragover"
	DRAG_START = "dragstart"
	DROP       = "drop"
	SCROLL     = "scroll"
)

// Clipboard
const (
	COPY  = "copy"
	CUT   = "cut"
	PASTE = "paste"
)

// Other
const (
	INTERSECT     = "intersect"
	INTERVAL      = "interval"
	RAF           = "raf"
	SIGNAL_CHANGE = "signal-change"
)
