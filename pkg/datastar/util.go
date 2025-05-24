package datastar

import (
	"encoding/json"
	"html/template"
	"regexp"
	"strings"
)

var (
	stateValueMatcher = regexp.MustCompile(`"\$(\w+)"`)
)

func createAttributeValue(value any) string {
	var result string
	switch v := value.(type) {
	case string:
		result = v
	case map[string]string, map[string]any, Map:
		marshaledMap, err := json.Marshal(v)
		if err == nil {
			result = string(marshaledMap)
		}
		if strings.Contains(result, `"$`) {
			result = stateValueMatcher.ReplaceAllString(result, `$$$1`)
		}
	default:
		result = template.HTMLEscaper(v)
	}
	return result
}

func applyModifiers(attributeName string, modifiers ...Modifier) string {
	builder := new(strings.Builder)
	builder.WriteString(attributeName)
	for _, modifier := range modifiers {
		if len(modifier.Type) == 0 {
			continue
		}
		builder.WriteString(modifier.Type)
		if modifier.Values != nil {
			for _, modifierValue := range modifier.Values {
				builder.WriteString(".")
				builder.WriteString(modifierValue)
			}
		}
	}
	return builder.String()
}
