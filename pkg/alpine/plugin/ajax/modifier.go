package ajax

import "strings"

func Alias(alias string) string {
	return ":" + alias
}

func applyDirectiveModifiers(attributeName string, modifiers ...string) string {
	for _, modifier := range modifiers {
		if len(modifier) == 0 {
			continue
		}
		if strings.HasPrefix(modifier, ":dynamic") {
			attributeName += modifier
			continue
		}
		if strings.HasPrefix(modifier, ":") {
			continue
		}
		attributeName += "." + modifier
	}
	return attributeName
}

func applyValueModifiers(attributeName string, modifiers ...string) string {
	for _, modifier := range modifiers {
		if len(modifier) == 0 || strings.HasPrefix(modifier, ":dynamic") {
			continue
		}
		if strings.HasPrefix(modifier, ":") {
			attributeName += modifier
			continue
		}
	}
	return attributeName
}
