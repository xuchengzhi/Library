package main

import (
	// "fmt"
	"github.com/xuchengzhi/Library/Time"
	"io/ioutil"
	"os"
	"strings"
	// "text/template"
	// "time"
	// "bytes"
)

/*
	生成代码
*/

var times = GetTime.TF()

type Inventory struct {
	Times  string
	Authod string
}

func Readfile(filename, f string) {
	files, _ := ioutil.ReadFile(filename)
	tmp1 := strings.Replace(string(files), "{{time}}", times, -1)
	tmp2 := strings.Replace(tmp1, "{{.authod}}", "xuchengzhi", -1)
	tmp := []byte(tmp2)
	Writefile(f+".go", tmp)

}

func Writefile(filename string, str []byte) {
	// str := []byte("package main \n\nimport (\n   \"fmt\"\n)")
	files, _ := os.Create(filename)
	defer files.Close()
	files.Write(str)
	files.Close()
}

func News(files string) {
	Readfile("tmp.go", files)
}

func main() {
	News("ceshi")
}
