package templates

// TODO convert this to a struct for type safty
var SlackMessages = map[string]string{
	"NewStargazer": `New Github star for _{{ .RepoName }}_ repo!.
The {{ .RepoName }} repo now has {{ .StartCount }} stars! :tada:.
Your new fan is <{{ .Username }}|{{.Url}}>
`,
	"RepoEvent": `There was a Repo Event for Repo: _{{ .RepoName }}_.
The event action was: *{{ .Action }}*
The actor was <{{ .Username }}|{{.Url}}>`}









