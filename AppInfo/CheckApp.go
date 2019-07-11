package CheckApp

import (
	// "bytes"
	"encoding/xml"
	"fmt"
	"github.com/lunny/axmlParser"
	"github.com/xuchengzhi/Library/FileZip"
	"io/ioutil"
	"os"
	"path"
	// "reflect"
	"strings"
)

type AppInfo struct {
	Name    string
	Version string
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
}

func FileFormat() (AppInfo, bool) {
	files := "./SJZZ/Payload/SJZZ.app/Info.plist"
	stats, err := PathExists(files)
	var info AppInfo
	if err != nil {
		fmt.Println(err)
		return info, false
	}
	var name, buildnum, builds, Appname string

	if stats {
		var result Result
		content, _ := ioutil.ReadFile(files)
		xml.Unmarshal(content, &result)
		strs := result.StringList
		name = strs[7]
		buildnum = strs[19]
		builds = strs[11]
		Appname = strs[23]

		fmt.Println(name, Appname, buildnum, builds)

		info.Name = Appname
		info.Version = builds
		return info, true
	} else {
		return info, false
	}

}

func test() {

}

func Adr(app string) AppInfo {
	listener := new(axmlParser.AppNameListener)
	axmlParser.ParseApk(app, listener)
	var info AppInfo
	info.Name = listener.PackageName
	info.Version = listener.VersionName
	return info
}

func IOS(app string) AppInfo {
	if path.Ext(app) == "ipa" {
		app = ZipRename(app)
	}
	ZIP.Unzip(app, "SJZZ")

	info, _ := FileFormat()

	return info
}
