package gox

import (
	"fmt"
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestRenderNodes(t *testing.T) {
	t.Run(
		"render attribute without value", func(t *testing.T) {
			checkedAttr := nodeRenderer{node{nodeType: nodeAttribute, name: attributeChecked}}
			builder := new(strings.Builder)
			checkedAttr.renderAttribute(builder)
			assert.Equal(t, attributeChecked, builder.String())
		},
	)
	t.Run(
		"render attribute with value", func(t *testing.T) {
			classes := "underline hover:no-underline"
			checkedAttr := nodeRenderer{node{nodeType: nodeAttribute, name: attributeClass, value: classes}}
			builder := new(strings.Builder)
			checkedAttr.renderAttribute(builder)
			assert.Equal(t, fmt.Sprintf(`%s="%s"`, attributeClass, classes), builder.String())
		},
	)
	t.Run(
		"render text", func(t *testing.T) {
			text := "test"
			checkedAttr := nodeRenderer{node{nodeType: nodeText, value: text}}
			assert.Equal(t, text, checkedAttr.renderText())
		},
	)
	t.Run(
		"render non-void element", func(t *testing.T) {
			divEl := nodeRenderer{node{nodeType: nodeElement, name: elementDiv}}
			builder := new(strings.Builder)
			divEl.renderElement(builder)
			assert.Equal(t, "<div></div>", builder.String())
		},
	)
	t.Run(
		"render void element", func(t *testing.T) {
			inputEl := nodeRenderer{node{nodeType: nodeElement, name: elementInput}}
			builder := new(strings.Builder)
			inputEl.renderElement(builder)
			assert.Equal(t, "<input />", builder.String())
		},
	)
	t.Run(
		"render non-void element with attributes", func(t *testing.T) {
			classes := "bg-blue-400 hover:bg-blue-600"
			buttonEl := nodeRenderer{
				node{
					nodeType: nodeElement,
					name:     elementButton,
					attributes: []node{
						{nodeType: nodeAttribute, name: attributeType, value: "button"},
						{nodeType: nodeAttribute, name: attributeClass, value: classes},
					},
				},
			}
			builder := new(strings.Builder)
			buttonEl.renderElement(builder)
			assert.Equal(t, fmt.Sprintf(`<button type="button" class="%s"></button>`, classes), builder.String())
		},
	)
	t.Run(
		"render non-void element with children", func(t *testing.T) {
			divEl := nodeRenderer{
				node{
					nodeType: nodeElement,
					name:     elementDiv,
					children: []node{
						{
							nodeType: nodeElement,
							name:     elementButton,
							children: []node{
								{nodeType: nodeText, value: "test"},
							},
						},
					},
				},
			}
			builder := new(strings.Builder)
			divEl.renderElement(builder)
			assert.Equal(t, "<div><button>test</button></div>", builder.String())
		},
	)
	t.Run(
		"render non-void element with attributes and children", func(t *testing.T) {
			classes := "bg-blue-400 hover:bg-blue-600"
			divEl := nodeRenderer{
				node{
					nodeType: nodeElement,
					name:     elementDiv,
					attributes: []node{
						{nodeType: nodeAttribute, name: attributeClass, value: classes},
					},
					children: []node{
						{
							nodeType: nodeElement,
							name:     elementButton,
							attributes: []node{
								{nodeType: nodeAttribute, name: attributeType, value: "button"},
							},
							children: []node{
								{nodeType: nodeText, value: "test"},
							},
						},
					},
				},
			}
			
			builder := new(strings.Builder)
			divEl.renderElement(builder)
			assert.Equal(
				t,
				fmt.Sprintf(`<div class="%s"><button type="button">test</button></div>`, classes),
				builder.String(),
			)
		},
	)
}
