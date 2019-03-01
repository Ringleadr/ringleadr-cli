package Init

import (
	"errors"
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
	var agogosArgs = []string{"-background"}
	//Grab the connect flag in case it's non empty
	address := ctx.String("connect")
	if address != "" {
		agogosArgs = append(agogosArgs, "-connect", address)
	}

	proxy := ctx.Bool("proxy")
	if proxy {
		agogosArgs = append(agogosArgs, "-proxy")
	}

	addr := ctx.String("addr")
	if addr != "" {
		agogosArgs = append(agogosArgs, "-addr", addr)
	}
	//Check if there is already an instance running
	if _, err := Requests.GetRequest("http://localhost:14440/ping"); err == nil {
		return errors.New("agogos host already running on this machine")
	}

	//if not, try to run one
	cmd, err := tryStartCommand("agogos-host", agogosArgs)
	if err == nil {
		//Release the process from our current instance so it doesn't quit when we do
		err = cmd.Process.Release()
		if err != nil {
			return err
		}

		log.Println("Using the existing binary. If you need to update the binary use `agogos-cli host update`")
		log.Println("agogos-host started with args", agogosArgs)
		log.Println("It may take about a minute to start up")
		return nil
	}

	log.Println("agogos-host not found on this machine. Downloading latest binary.")
	//finally, download binary
	err = DownloadLatestBinary()
	if err != nil {
		return err
	}

	log.Println("agogos-host installed to /usr/local/bin/agogos-host")

	//Try to start the binary
	cmd, err = tryStartCommand("agogos-host", agogosArgs)
	if err != nil {
		return err
	}

	//Release from our process
	err = cmd.Process.Release()
	if err != nil {
		return err
	}

	log.Println("agogos-host started with args", agogosArgs)
	log.Println("The service is now running, but may take about a minute to start up")
	log.Println("Check $HOME/.agogos/host/host.log for logging messages")
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

func tryStartCommand(name string, args []string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func DownloadLatestBinary() error {
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
	return nil
}
