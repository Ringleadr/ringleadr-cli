package Format

import (
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-datatypes"
	"html/template"
	"os"
)

const applicationList = `NAME	#COMPONENTS
{{range .}}{{.Name}}	{{len .Components}}
{{end}}`

func PrintApplications(apps *[]Datatypes.Application) error {
	templ, err := template.New("apps").Parse(applicationList)
	if err != nil {
		return Errors.FormatError()
	}

	err = templ.Execute(os.Stdout, apps)
	if err != nil {
		return Errors.FormatError()
	}
	return nil
}
