package Format

import (
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-datatypes"
	"html/template"
	"os"
	"strings"
)

const storageList = `NAME
{{if .}}{{range $item := .}}{{$item.Name | trimAgogos}}
{{end}}{{end}}`

func PrintStorage(storage *[]Datatypes.Storage) error {
	funcMap := template.FuncMap{
		"trimAgogos": func(input string) string {
			return strings.TrimPrefix(input, "agogos-")
		},
	}

	templ, err := template.New("storage").Funcs(funcMap).Parse(storageList)
	if err != nil {
		return Errors.FormatError()
	}

	err = templ.Execute(os.Stdout, storage)
	if err != nil {
		return Errors.FormatError()
	}
	return nil
}
