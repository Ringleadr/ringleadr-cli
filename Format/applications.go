package Format

import (
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-datatypes"
	"os"
	"text/tabwriter"
)

const applicationList = `NAME		#COMPONENTS		#COPIES
{{if .}}{{range $item := .}}{{$item.Name}}		{{len $item.Components}}		{{$item.Copies}}
{{end}}{{end}}`

func PrintApplications(apps *[]Datatypes.Application) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, err := fmt.Fprintln(w, "NAME\t#COMPONENTS\t#COPIES")
	if err != nil {
		return Errors.FormatError()
	}
	for _, app := range *apps {
		_, err := fmt.Fprintln(w, fmt.Sprintf("%s\t%d\t%d", app.Name, len(app.Components), app.Copies))
		if err != nil {
			return err
		}
	}

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}
