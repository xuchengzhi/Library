package ADB

import (
	"fmt"
	// "github.com/xuchengzhi/Library/Http"
	"os/exec"
	// "sync"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
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

func Getip() string {
	cmd := exec.Command("./windows/adb.exe", "shell", "ifconfig wlan0")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	ip := ""
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error(), stderr.String())
	} else {
		if strings.Index("inet addr", out.String()) == -1 {
			start_num := strings.Index(out.String(), "addr:") + 5
			ip = out.String()[start_num : start_num+16]
			log.Println(ip)
		} else {
			log.Println("设备未连接或未开启调试模式")
		}
	}

	return ip
}

func Connect(ip string) {
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

			url := ""
			if len(ip) > 0 {
				url = fmt.Sprintf("http://192.168.248.188/api/connect_ip/%v", ip)
				log.Println(url)
			} else {

				return
			}
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}

			defer resp.Body.Close()
			ip = strings.Replace(ip, " ", "", -1)
			urls := "http://" + ip + ":7912/remote"
			exec.Command("cmd", "/c", "start", urls).Start()
			time.Sleep(1000)
		}
	}

}

func Device() {
	cmd := exec.Command("./windows/adb.exe", "devices")
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

// func main() {

// 	// SyncTest()
// 	// Logcat("com.huawei.android.thememanager")
// 	Device()
// 	ip := Getip()
// 	Connect(ip)
// 	time.Sleep(time.Second * 5)
// }
