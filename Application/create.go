package Application

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/urfave/cli"
	"io/ioutil"
)

func CreateApplication(c *cli.Context) error {
	filePath := c.String("file")
	if filePath == "" {
		cli.ShowSubcommandHelp(c)
		return errors.New("config file not specified")
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if !Config.IsJson(file) {
		return errors.New("provided file is not valid JSON")
	}

	_, err = Requests.PostRequest(fmt.Sprintf("%s/applications", Config.GetAgogosHostUrl()),
		bytes.NewReader(file))
	if err != nil {
		return err
	}
	return nil
}
