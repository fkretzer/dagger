package templates

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed src
var srcs embed.FS

// New creates a new template with all the template dependencies set up.
func New(schemaVersion string, moduleName string, moduleParentPath string) *template.Template {
	topLevelTemplate := "api"
	templateDeps := []string{
		topLevelTemplate, "header", "objects", "object", "method", "method_solve", "call_args", "method_comment", "types", "args", "default",
	}

	fileNames := make([]string, 0, len(templateDeps))
	for _, tmpl := range templateDeps {
		fileNames = append(fileNames, fmt.Sprintf("src/%s.ts.gtpl", tmpl))
	}

	funcs := TypescriptTemplateFuncs(schemaVersion, moduleName, moduleParentPath)
	tmpl := template.Must(template.New(topLevelTemplate).Funcs(funcs).ParseFS(srcs, fileNames...))
	return tmpl
}
