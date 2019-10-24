// package ADB
package main

import (
	"fmt"
	// "github.com/xuchengzhi/Library/Http"
	"os/exec"
	// "sync"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	h       bool
	s       string
	ip      string
	device  bool
	getip   bool
	conn    bool
	logcat  string
	sys     string
	adb_cmd string
)

func init() {
	sys = runtime.GOOS
	if sys == "windows" {
		adb_cmd = "./windows/adb.exe"
	} else if sys == "linux" {
		adb_cmd = "./linux/adb"
	} else {
		adb_cmd = "./mac/adb"
	}

	flag.BoolVar(&h, "h", false, "查看帮助")
	flag.BoolVar(&device, "device", false, "获取devices")
	flag.BoolVar(&getip, "getip", false, "获取手机ip")
	flag.BoolVar(&conn, "conn", false, "将手机通过远程连接到服务器，服务器可操作该手机")
	flag.String("logcat", "logcat", "查看日志")
	flag.StringVar(&s, "s", "", "测试")
	flag.StringVar(&ip, "ip", "", "服务器IP地址")
	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage

	// log.Println(sys)
}

func Formats(f *flag.Flag) string {
	res := ""
	if f != nil {
		res = fmt.Sprintf("%v", f.Value)
	} else {
		fmt.Println(nil)
	}
	return res
}

func Logcat(app string) {
	// cmdstr := fmt.Sprintf("adb logcat | grep %v", app)
	cmd := exec.Command(adb_cmd, "logcat")
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

	cmd := exec.Command(adb_cmd, "shell", "ifconfig wlan0")
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

func Connect() {

	ip := Getip()

	cmd := exec.Command(adb_cmd, "tcpip", "5555")

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
			// ConntToIP(ip)
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

func ConntToIP(ip string) {
	cmdstr := fmt.Sprintf("connect %v:5555", ip)
	log.Println(cmdstr)
	cmd := exec.Command(adb_cmd, cmdstr)

	// deviceinfo = nil
	// num = 0
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
	log.Println(ip)

}

func Device() {
	cmd := exec.Command(adb_cmd, "devices")
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
	// 	// SyncTest()
	// Logcat("com.huawei.android.thememanager")
	// 	Device()
	// 	ip := Getip()
	// 	Connect(ip)
	// 	time.Sleep(time.Second * 5)
	flag.Parse()

	if h {
		flag.Usage()
	} else if device {
		Device()
	} else if getip {
		Getip()
	} else if conn {
		Connect()
	}
	// ConntToIP("192.168.247.134")
	// msg := Formats(flag.Lookup("test"))
	// fmt.Println(msg)
}

func usage() {
	fmt.Fprintf(os.Stderr, `App version: Test/1.10.0
Usage: adbs.exe [-h] [-device] [-getip] [-ip ipaddress] [-conn]
Options:
`)
	flag.PrintDefaults()
}
