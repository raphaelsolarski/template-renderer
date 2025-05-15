package renderer

import (
	"bytes"
	"html/template"
)

func RenderTemplate(tpl string, data map[string]string) (string, error) {
	parsedTpl, parseErr := template.New("template").Parse(tpl)
	if parseErr != nil {
		return "", parseErr
	}

	var pngOutput bytes.Buffer
	executeErr := parsedTpl.Execute(&pngOutput, data)
	if executeErr != nil {
		return "", executeErr
	}
	return pngOutput.String(), nil
}
