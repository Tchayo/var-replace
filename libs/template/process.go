package template

import (
	"bytes"
	"text/template"
)

func process(t *template.Template, vars interface{}) string {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)

	if err != nil {
		panic(err)
	}
	return tmplBytes.String()
}

// ProcessString : processing a template string
func ProcessString(str string, vars interface{}) string {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		panic(err)
	}
	return process(tmpl, vars)
}

// ProcessFile : processing a template file
func ProcessFile(fileName string, vars interface{}) string {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		panic(err)
	}
	return process(tmpl, vars)
}
