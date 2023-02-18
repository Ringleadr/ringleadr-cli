package Config

import "os"

var (
	agogosHostUrl string
)

func SetupConfig() {
	if tempHost := os.Getenv("AGOGOS_HOST_URL"); tempHost != "" {
		agogosHostUrl = tempHost
	} else {
		agogosHostUrl = "http://localhost:14440"
	}

}

func GetAgogosHostUrl() string {
	return agogosHostUrl
}
