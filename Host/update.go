package Host

import (
	"errors"
	"github.com/GodlikePenguin/agogos-cli/Init"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func UpdateHost(c *cli.Context) error {
	err := os.Remove("/usr/local/bin/agogos-host")
	if err != nil {
		//Ignore error if it turns out the host was not installed
		if !strings.Contains(err.Error(), "no such file or directory") {
			return errors.New("Error deleting existing binary: " + err.Error())
		}
	}
	err = Init.DownloadLatestBinary()
	if err != nil {
		return errors.New("Error downloading latest binary: " + err.Error())
	}
	log.Println("Latest binary has been installed and moved to /usr/local/bin/agogos-host")
	log.Println("It has not been executed but you can do so yourself manually, or by running 'agogos-cli init'")
	return nil
}
