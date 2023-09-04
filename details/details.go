package details

import (
	"log"
	"net"
	"os"
)

func GetHostname() (string, error) {
	hostname, _ := os.Hostname()
	return hostname, nil
}

func GetIPAddress() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
