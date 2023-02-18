package Networks

import (
	"encoding/json"
	"fmt"
	"github.com/Ringleadr/ringleadr-cli/internal/Config"
	"github.com/Ringleadr/ringleadr-cli/internal/Errors"
	"github.com/Ringleadr/ringleadr-cli/internal/Format"
	"github.com/Ringleadr/ringleadr-cli/internal/Requests"
	Datatypes "github.com/Ringleadr/ringleadr-datatypes"
	"github.com/urfave/cli"
)

func ListNetworks(c *cli.Context) error {
	resp, err := Requests.GetRequest(fmt.Sprintf("%s/networks", Config.GetAgogosHostUrl()))
	if err != nil {
		return err
	}

	var networks []Datatypes.Network

	//Bind response to struct
	if err := json.Unmarshal(resp, &networks); err != nil {
		return Errors.UnexpectedReponse()
	}
	return Format.PrintNetworks(&networks)
}
