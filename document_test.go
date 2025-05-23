package gox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	t.Run(
		"render doctype", func(t *testing.T) {
			assert.Equal(
				t, "<!DOCTYPE html>", Render(Doctype()),
			)
		},
	)
}
