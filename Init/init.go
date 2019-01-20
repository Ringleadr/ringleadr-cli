package Init

import (
	"errors"
	"fmt"
	"github.com/GodlikePenguin/agogos-cli/Requests"
	"github.com/buger/jsonparser"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Init(ctx *cli.Context) error {
	//Tell the host it will run in the background
	var agogosArgs = "-background "
	//Grab the connect flag in case it's non empty
	address := ctx.String("connect")
	if address != "" {
		agogosArgs += fmt.Sprintf("-connect %s ", address)
	}
	//Check if there is already an instance running
	if _, err := Requests.GetRequest("http://localhost:14440/ping"); err == nil {
		return errors.New("agogos host already running on this machine")
	}

	//if not, try to run one
	cmd, err := tryStartCommand("agogos-host " + agogosArgs)
	if err == nil {
		//Release the process from our current instance so it doesn't quit when we do
		err = cmd.Process.Release()
		if err != nil {
			return err
		}

		log.Println("You already have agogos-host installed. In the future you can run it using `agogos-host`. It is now running.")
		return nil
	}

	log.Println("agogos-host not found on this machine. Downloading latest binary.")
	//finally, download binary
	//Get URL of latest download
	body, err := Requests.GetRequest("https://api.github.com/repos/GodlikePenguin/agogos-host-release/releases/latest")
	if err != nil {
		return err
	}

	downloadURL, err := getDownloadURL(body)
	if downloadURL == "" || err != nil {
		return errors.New("could not get latest download URL for agogos-host")
	}

	err = DownloadFile("agogos-host", downloadURL)
	if err != nil {
		return err
	}

	//chmod to make executable
	err = os.Chmod("agogos-host", 0755)
	if err != nil {
		return err
	}

	//Move downloaded file to new location
	err = os.Rename("agogos-host", "/usr/local/bin/agogos-host")
	if err != nil {
		return err
	}

	//Try to start the binary
	cmd, err = tryStartCommand("agogos-host " + agogosArgs)
	if err != nil {
		return err
	}

	//Release from our process
	err = cmd.Process.Release()
	if err != nil {
		return err
	}

	log.Println("Downloaded latest agogos-release binary. It has been added to your path and is called `agogos-host`.",
		"The service is now running. This may take a few minutes to start up")
	return nil
}

func getDownloadURL(body []byte) (string, error) {
	var downloadURL string
	//Iterate through JSON response and get download for this OS
	_, err := jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		val, _, _, err := jsonparser.Get(value, "browser_download_url")
		if err == nil && strings.Contains(string(val), runtime.GOOS) {
			downloadURL = string(val)
		}
	}, "assets")
	if err != nil {
		return "", err
	}
	return downloadURL, nil
}

func tryStartCommand(name string) (*exec.Cmd, error) {
	cmd := exec.Command(name)
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
