package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	t.Run(
		"create and validate element node", func(t *testing.T) {
			name := "test"
			n, ok := createElement(name).(node)
			assert.Equal(t, true, ok)
			assert.Equal(t, nodeElement, n.nodeType)
			assert.Equal(t, n.name, name)
		},
	)
	t.Run(
		"create and validate attribute node", func(t *testing.T) {
			name := "test-name"
			value := "test-value"
			n, ok := createAttribute(name, value).(node)
			assert.Equal(t, true, ok)
			assert.Equal(t, nodeAttribute, n.nodeType)
			assert.Equal(t, n.name, name)
			assert.Equal(t, n.value, value)
		},
	)
	t.Run(
		"create and validate shared node as element", func(t *testing.T) {
			name := "test-name"
			n, ok := createShared(name, nodeElement).(node)
			assert.Equal(t, true, ok)
			assert.Equal(t, nodeElement, n.nodeType)
			assert.Equal(t, n.name, name)
		},
	)
	t.Run(
		"create and validate shared node as attribute", func(t *testing.T) {
			name := "test-name"
			value := "test-value"
			n, ok := createShared(name, nodeAttribute, Text(value)).(node)
			assert.Equal(t, true, ok)
			assert.Equal(t, nodeAttribute, n.nodeType)
			assert.Equal(t, n.name, name)
			assert.Equal(t, n.value, value)
		},
	)
}
