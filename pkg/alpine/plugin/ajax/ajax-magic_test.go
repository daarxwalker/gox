package ajax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAjaxMagic(t *testing.T) {
	t.Run(
		"render get", func(t *testing.T) {
			assert.Equal(
				t,
				`$ajax('/test',{'method':'get'})`,
				GET("/test"),
			)
		},
	)
	t.Run(
		"render post map", func(t *testing.T) {
			assert.Equal(
				t,
				`$ajax('/test',{'body':{'test':true},'method':'post'})`,
				POST("/test", WithBody(Map{"test": true})),
			)
		},
	)
	t.Run(
		"render patch struct with headers", func(t *testing.T) {
			type test struct {
				Value int `json:"value"`
			}
			assert.Equal(
				t,
				`$ajax('/test',{'body':{'value':1},'method':'post','headers':{'x-authorization':'123456789'}})`,
				POST("/test", WithBody(test{Value: 1}), WithHeaders(Map{"x-authorization": "123456789"})),
			)
		},
	)
}
