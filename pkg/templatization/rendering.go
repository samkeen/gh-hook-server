package templatization

import (
	"bytes"
	"log"
	"text/template"
)

const MSG_TEMPLATE_PATH = "./msgtemplates"

type Payload struct {
	RepoName, StartCount, Username, Url string
}

func GetSlackMessageTemplate(templateName string) (*template.Template) {
	tmpl := template.New(templateName + ".txt")
	tmpl, err := tmpl.ParseFiles(MSG_TEMPLATE_PATH + "/" + templateName + ".txt")
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
