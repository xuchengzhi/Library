package main

import (
	// "fmt"
	// "github.com/xuchengzhi/Library/Http"
	"os/exec"
	// "sync"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Logcat(app string) {
	// cmdstr := fmt.Sprintf("adb logcat | grep %v", app)
	cmd := exec.Command("adb", "logcat")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error(), stderr.String())
	} else {
		log.Println(out.String())
	}

}

func Connect() {
	cmd := exec.Command("./windows/adb.exe", "tcpip", "5555")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error(), stderr.String())
	} else {

		if strings.Index("restarting in TCP mode port", out.String()) == -1 {
			resp, err := http.Get("http://localhost/api/connect_ip/192.168.247.177")
			if err != nil {
				log.Println(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}

			log.Println(string(body))
		}
	}

}

func Test() {
	cmd := exec.Command("adb", "logcat")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}

func main() {

	// SyncTest()
	// Logcat("com.huawei.android.thememanager")
	Connect()

}
