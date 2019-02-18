package General

import (
	"bufio"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Config"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func Purge(c *cli.Context) error {
	fmt.Println("This will delete all resources in the system, are you sure you want to continue? (Y/N)")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	resp := scanner.Text()
	if strings.ToLower(resp) != "y" {
		return errors.New("Non 'y' response received. Will not continue.")
	}

	if err := scanner.Err(); err != nil {
		return errors.New("Error reading user input: " + err.Error())
	}
	_, err := Requests.DeleteRequest(fmt.Sprintf("%s/everything", Config.GetAgogosHostUrl()))
	return err
}
