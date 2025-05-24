package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	t.Run(
		"create primitive attribute value", func(t *testing.T) {
			assert.Equal(
				t,
				`true`,
				createAttributeValue(true),
			)
		},
	)
	t.Run(
		"create map attribute value", func(t *testing.T) {
			assert.Equal(
				t,
				`{"test":true}`,
				createAttributeValue(Map{"test": true}),
			)
		},
	)
	t.Run(
		"create map state attribute value", func(t *testing.T) {
			assert.Equal(
				t,
				`{"test1":$test2}`,
				createAttributeValue(Map{"test1": "$test2"}),
			)
		},
	)
	t.Run(
		"apply modifiers to attribute name", func(t *testing.T) {
			assert.Equal(
				t,
				`data-test__case.kebab`,
				applyModifiers("data-test", WithCase(KEBAB)),
			)
		},
	)
}
