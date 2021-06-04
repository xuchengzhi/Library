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
	"regexp"
	"runtime"
	"strings"
	"time"
)

var (
	h        bool
	s        string
	ip       string
	device   bool
	getip    bool
	tcpip    bool
	version  bool
	conn     string
	logcat   string
	sys      string
	adb_cmd  string
	_version string
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
	_version = "0.0.1"
	flag.BoolVar(&h, "h", false, "查看帮助")
	flag.BoolVar(&device, "device", false, "获取devices")
	flag.BoolVar(&tcpip, "t", false, "开启远程连接，port：5555")
	flag.BoolVar(&getip, "getip", false, "获取手机ip")
	flag.StringVar(&conn, "conn", "nil", "将手机通过远程连接到服务器，服务器可操作该手机")
	flag.StringVar(&logcat, "logcat", "nil", "输入包名查看日志")
	// flag.StringVar(&s, "s", "", "测试")
	flag.BoolVar(&version, "v", false, "版本号")
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
	cmdstr := fmt.Sprintf("logcat | grep %v", app)
	log.Println(cmdstr)
	cmd := exec.Command(adb_cmd, "logcat", "| grep ", app)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	log.Println(stderr.String())
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
			log.Println("当前设备ip:", ip)
		} else {
			log.Println("设备未连接或未开启调试模式")
		}
	}

	return ip
}

func Tcpip() string {
	ip := Getip()

	cmd := exec.Command(adb_cmd, "tcpip", "5555")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error(), stderr.String())
		return ""
	}
	log.Println(fmt.Sprintf("tcp start %v:5555", ip))
	return fmt.Sprintf("tcp start %v:5555", ip)

}

func Connect(ip string) {
	url := ""
	if len(ip) > 0 {

		url = fmt.Sprintf("http://192.168.248.188/api/connect_ip/%v", ip)

	} else {
		log.Println("ip无效")
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

func ConntToIP(ip string) {
	if len(ip) <= 0 {
		log.Println("ip无效")
		return
	}
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

func Version() {
	fmt.Println(_version)
}

func Device() []string {
	var devlist []string
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
		// res := strings.Split("device", string(opBytes))
		tmp := regexp.MustCompile(`\n[\P{Han}]+\t`)
		tmp1 := tmp.FindAllString(string(opBytes), -1)
		for i := 0; i < len(tmp1); i++ {
			tmp2 := strings.Replace(tmp1[i], "\n", "", -1)
			dev := strings.Replace(tmp2, "\t", "", -1)
			devlist = append(devlist, dev)
		}

	}
	res := fmt.Sprintf("在线设备(%v台):%v", len(devlist), devlist)
	log.Println(res)
	return devlist
}

func main() {
	// 	// SyncTest()
	// Logcat("com.huawei.android.thememanager")
	// 	Device()
	// 	ip := Getip()
	// 	Connect(ip)
	// 	time.Sleep(time.Second * 5)
	flag.Parse()

	// ConntToIP(ip)
	conn := Formats(flag.Lookup("conn"))
	logcat := Formats(flag.Lookup("logcat"))
	if h {
		flag.Usage()
	} else if conn != "nil" {
		Connect(conn)
	} else if logcat != "nil" {
		Logcat(logcat)
	} else if device {
		Device()
	} else if getip {
		Getip()
	} else if tcpip {
		Tcpip()
	} else if version {
		Version()
	} else {
		flag.Usage()
	}

	// ConntToIP("192.168.247.134")
	// msg := Formats(flag.Lookup("test"))
	// fmt.Println(msg)
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: Gadb.exe [-h] [-device] [-t tcpip] [-getip] [-ip ipaddress] [-conn ip] [-logcat packageName]
Options:
`)
	flag.PrintDefaults()
}
