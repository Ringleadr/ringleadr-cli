package Format

import (
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Errors"
	"github.com/GodlikePenguin/agogos-datatypes"
	"os"
	"strings"
	"text/tabwriter"
)

func PrintNetworks(networks *[]Datatypes.Network) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	_, err := fmt.Fprintln(w, "NAME")
	if err != nil {
		return Errors.FormatError()
	}
	for _, net := range *networks {
		_, err := fmt.Fprintln(w, strings.Replace(net.Name, "agogos-", "", 1))
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
