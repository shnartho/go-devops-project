package details

import "os"

func GetHostname() (string, error) {
	hostname, _ := os.Hostname()
	return hostname, nil
}
