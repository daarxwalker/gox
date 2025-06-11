package ajax

import (
	"strings"

	"github.com/daarxwalker/gox"
)

const (
	targetDirective    = "x-target"
	headersDirective   = "x-headers"
	mergeDirective     = "x-merge"
	autofocusDirective = "x-autofocus"
	syncDirective      = "x-sync"
)

func Target(target string, modifiers ...string) gox.Node {
	return gox.CreateAttribute[string](applyDirectiveModifiers(targetDirective, modifiers...))(
		applyValueModifiers(
			target, modifiers...,
		),
	)
}

func Targets(targets []string, modifiers ...string) gox.Node {
	return gox.CreateAttribute[string](
		applyDirectiveModifiers(
			targetDirective, modifiers...,
		),
	)(applyValueModifiers(strings.Join(targets, " "), modifiers...))
}

func Headers(headers Map) gox.Node {
	return gox.CreateAttribute[string](headersDirective)(headers.EscapedString())
}

func Merge(mergeType string, modifiers ...string) gox.Node {
	if len(mergeType) == 0 {
		mergeType = REPLACE
	}
	return gox.CreateAttribute[string](applyDirectiveModifiers(mergeDirective, modifiers...))(mergeType)
}

func AutoFocus() gox.Node {
	return gox.CreateAttribute[string](autofocusDirective)()
}

func Sync() gox.Node {
	return gox.CreateAttribute[string](syncDirective)()
}
