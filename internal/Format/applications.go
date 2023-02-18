package Format

import (
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Errors"
	Datatypes "github.com/Ringleadr/ringleadr-datatypes"
	"os"
	"strings"
	"text/tabwriter"
)

func PrintApplications(apps *[]Datatypes.Application) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, err := fmt.Fprintln(w, "NAME\tCOMPONENTS READY\tNODE\tMESSAGES")
	if err != nil {
		return Errors.FormatError()
	}
	for _, app := range *apps {
		readyComps := 0
		for _, comp := range app.Components {
			if strings.Contains(comp.Status, "running") {
				readyComps += 1
			}
		}
		_, err := fmt.Fprintln(w, fmt.Sprintf("%s\t(%d/%d)\t%s\t%s", app.Name, readyComps, len(app.Components), app.Node, app.Messages))
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
