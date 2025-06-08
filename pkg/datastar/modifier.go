package datastar

type Modifier struct {
	Type   string
	Values []string
}

func CreateModifier(name string, values []string) Modifier {
	return Modifier{
		Type:   name,
		Values: values,
	}
}

func WithAuto() Modifier {
	return CreateModifier("__auto", nil)
}

func WithCapture() Modifier {
	return CreateModifier("__capture", nil)
}

func WithCase(value string) Modifier {
	return CreateModifier("__case", []string{value})
}

func WithDebounce(values ...string) Modifier {
	return CreateModifier("__debounce", values)
}

func WithDelay(values ...string) Modifier {
	return CreateModifier("__delay", values)
}

func WithDuration(values ...string) Modifier {
	return CreateModifier("__duration", values)
}

func WithFocus() Modifier {
	return CreateModifier("__focus", nil)
}

func WithFull() Modifier {
	return CreateModifier("__full", nil)
}

func WithIfMissing() Modifier {
	return CreateModifier("__ifmissing", nil)
}

func WithHalf() Modifier {
	return CreateModifier("__half", nil)
}

func WithHCenter() Modifier {
	return CreateModifier("__hcenter", nil)
}

func WithHEnd() Modifier {
	return CreateModifier("__hend", nil)
}

func WithHNearest() Modifier {
	return CreateModifier("__hnearest", nil)
}

func WithHStart() Modifier {
	return CreateModifier("__hstart", nil)
}

func WithInstant() Modifier {
	return CreateModifier("__instant", nil)
}

func WithOnce() Modifier {
	return CreateModifier("__once", nil)
}

func WithOutside() Modifier {
	return CreateModifier("__viewtransition", nil)
}

func WithPassive() Modifier {
	return CreateModifier("__passive", nil)
}

func WithPrevent() Modifier {
	return CreateModifier("__prevent", nil)
}

func WithSelf() Modifier {
	return CreateModifier("__self", nil)
}

func WithSession() Modifier {
	return CreateModifier("__session", nil)
}

func WithSmooth() Modifier {
	return CreateModifier("__smooth", nil)
}

func WithStop() Modifier {
	return CreateModifier("__stop", nil)
}

func WithThrottle(values ...string) Modifier {
	return CreateModifier("__throttle", values)
}

func WithViewTransition() Modifier {
	return CreateModifier("__viewtransition", nil)
}

func WithVCenter() Modifier {
	return CreateModifier("__vcenter", nil)
}

func WithVEnd() Modifier {
	return CreateModifier("__vend", nil)
}

func WithVNearest() Modifier {
	return CreateModifier("__vnearest", nil)
}

func WithVStart() Modifier {
	return CreateModifier("__vstart", nil)
}

func WithWindow() Modifier {
	return CreateModifier("__viewtransition", nil)
}
