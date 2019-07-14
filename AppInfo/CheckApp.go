package CheckApp

// package main

import (
	// "bytes"
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/lunny/axmlParser"
	// "github.com/xuchengzhi/Library/FileZip"
	"io/ioutil"
	"os"
	"os/exec"
	// "path"
	"path/filepath"
	"log"
	// "reflect"
	"strings"
)

type AppInfo struct {
	Name     string
	Version  string
	Appsname string
}

func ZipRename(file string) string {
	name := strings.Split(file, ".")[0]
	os.Rename(file, name+".zip")
	return name + ".zip"
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

type Result struct {
	XMLName    xml.Name `xml:"plist"` //标签上的标签名
	StringList []string `xml:"dict>string"`
	KeyList    []string `xml:"dict>key"`
}

func FileFormat() (AppInfo, bool) {
	files := "./Payload/SJZZ.app/Info.plist"
	stats, err := PathExists(files)
	var info AppInfo
	if err != nil {
		fmt.Println(err)
		return info, false
	}
	var name, builds, Appname string

	if stats {
		var result Result
		content, _ := ioutil.ReadFile(files)
		xml.Unmarshal(content, &result)

		strs := result.StringList
		keys := result.KeyList

		fmt.Println(len(strs), len(keys))
		// name = strs[7]
		// buildnum = strs[19]
		// builds = strs[11]
		// Appname = strs[23]

		// fmt.Println(name, Appname, buildnum, builds)
		for i := 0; i < len(strs); i++ {
			fmt.Println(keys[i], strs[i])
			if keys[i] == "CFBundleName" {
				fmt.Println(strs[i])
			}
		}
		info.Name = Appname
		info.Appsname = name
		info.Version = builds
		// os.RemoveAll("./Payload")
		return info, true
	} else {
		os.RemoveAll("./Payload")
		return info, false
	}

}

func Adr(app string) AppJson {
	listener := new(axmlParser.AppNameListener)
	axmlParser.ParseApk(app, listener)
	var info AppJson
	info.Name = listener.PackageName
	info.Version = listener.VersionName
	info.VCode = listener.VersionCode
	return info
}

type AppJson struct {
	Name    string
	Version string
	VCode   string
}

func IOS(app string) AppJson {
	abspath,_ := filepath.Abs(filepath.Dir("CheckApp.jar"))
	Apath := fmt.Sprintf("%v/CheckApp.jar",abspath)
	log.Println(Apth)
	cmd := exec.Command("java", "-jar", Apath, app)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	str := strings.Replace(string(out), ",", ",", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "=", ":", -1)
	str = strings.Replace(str, "{", "", -1)
	str = strings.Replace(str, "}", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	t := strings.Split(str, ",")
	var apps AppJson
	for i := 0; i < len(t); i++ {
		s := strings.Split(t[i], ":")
		if s[0] == "package" {
			apps.Name = s[1]
		} else if s[0] == "versionName" {
			apps.Version = s[1]

		} else if s[0] == "versionCode" {
			apps.VCode = s[1]

		}

	}
	return apps
}

// func main() {

// 	fmt.Println(IOS("/Users/yinmiaomiao/Desktop/code/go/ceshi.ipa"))
// 	// fmt.Println(Adr("./ceshi.apk"))
// 	// IOS("./SJZZ.ipa")
// }
