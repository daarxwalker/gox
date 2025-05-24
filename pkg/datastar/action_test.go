package datastar

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	t.Run(
		"create get", func(t *testing.T) {
			assert.Equal(
				t,
				`@get('/test')`,
				GET("/test"),
			)
		},
	)
	t.Run(
		"create get with config", func(t *testing.T) {
			assert.Equal(
				t,
				`@get('/test',{'headers':{'test':1},'includeLocal':true})`,
				GET("/test", Map{INCLUDE_LOCAL: true, HEADERS: Map{"test": 1}}),
			)
		},
	)
	t.Run(
		"create post", func(t *testing.T) {
			assert.Equal(
				t,
				`@post('/test')`,
				POST("/test"),
			)
		},
	)
	t.Run(
		"create put", func(t *testing.T) {
			assert.Equal(
				t,
				`@put('/test')`,
				PUT("/test"),
			)
		},
	)
	t.Run(
		"create patch", func(t *testing.T) {
			assert.Equal(
				t,
				`@patch('/test')`,
				PATCH("/test"),
			)
		},
	)
	t.Run(
		"create delete", func(t *testing.T) {
			assert.Equal(
				t,
				`@delete('/test')`,
				DELETE("/test"),
			)
		},
	)
	t.Run(
		"create clipboard", func(t *testing.T) {
			assert.Equal(
				t,
				`@clipboard('test')`,
				Clipboard("test"),
			)
		},
	)
	t.Run(
		"create setAll", func(t *testing.T) {
			assert.Equal(
				t,
				`@setAll('test',true)`,
				SetAll("test", true),
			)
		},
	)
	t.Run(
		"create toggleAll", func(t *testing.T) {
			assert.Equal(
				t,
				`@toggleAll('test')`,
				ToggleAll("test"),
			)
		},
	)
}
