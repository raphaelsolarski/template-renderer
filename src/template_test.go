package renderer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	tpl := `<svg width="{{.width}}" height="100" xmlns="http://www.w3.org/2000/svg"/>`
	data := map[string]string{
		"width": "200",
	}
	output, err := RenderTemplate(tpl, data)
	assert.Nil(t, err)
	assert.Equal(t, output, `<svg width="200" height="100" xmlns="http://www.w3.org/2000/svg"/>`)
}

func TestRenderTemplateShouldEscapeValuesAndText(t *testing.T) {
	tpl := `<svg width="{{.someAttrValue}}" height="100" xmlns="http://www.w3.org/2000/svg">{{.someText}}</svg>`
	data := map[string]string{
		"someAttrValue": "\"",
		"someText":      "<div>abc</div>",
	}
	output, err := RenderTemplate(tpl, data)
	assert.Nil(t, err)
	assert.Equal(t, output, `<svg width="&#34;" height="100" xmlns="http://www.w3.org/2000/svg">&lt;div&gt;abc&lt;/div&gt;</svg>`)
}
