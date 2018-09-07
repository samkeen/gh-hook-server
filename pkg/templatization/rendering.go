package templatization

import (
	"bytes"
	"github.com/samkeen/github-webhook-serverless/templates"
	"log"
	"text/template"
)


type Payload struct {
	RepoName, StartCount, Username, Url string
}

func GetSlackMessageTemplate(templateName string) (*template.Template) {
	tmpl := template.New(templateName + ".txt")
	tmpl, err := tmpl.Parse(templates.SlackMessages[templateName])
	if err != nil {
		log.Fatal("Parse: ", err)
		panic("Unable to parse tmpl")
	}
	return tmpl
}

func ExecuteTemplate(tmpl *template.Template, tmplatePayload Payload) string {
	var buf bytes.Buffer
	var err = tmpl.Execute(&buf, tmplatePayload)
	if err != nil {
		log.Fatal("Execute: ", err)
		panic("Execution of template failed")
	}
	return buf.String()
}