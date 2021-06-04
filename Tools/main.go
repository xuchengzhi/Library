package Tools

import (
	// "fmt"
	"bytes"
	"encoding/json"
	"log"
	"net"
	"os"
	"runtime"
)

func GetSystem() string {
	sys := runtime.GOOS
	return sys
}

func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

func GetIPv4ByInterface(name string) ([]string, error) {
	var ips []string

	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

func Printj(msg []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, msg, "", "	")

	if err != nil {
		log.Fatalln(err)
	}
	out.WriteTo(os.Stdout)
}
