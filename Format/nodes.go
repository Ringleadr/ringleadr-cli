package Format

import (
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-datatypes"
	"os"
	"text/tabwriter"
)

func PrintNodes(nodes *[]Datatypes.Node) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, err := fmt.Fprintln(w, "NAME\tADDRESS\tACTIVE")
	if err != nil {
		return Errors.FormatError()
	}
	for _, node := range *nodes {
		_, err := fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%t", node.Name, node.Address, node.Active))
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
