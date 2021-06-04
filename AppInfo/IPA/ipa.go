package ipa

import (
	"archive/zip"
	"bytes"
	// "encoding/json"
	// "encoding/xml"
	"fmt"
	// "github.com/pkg/errors"
	// "github.com/shogo82148/androidbinary"
	"image"
	"image/jpeg"  // handle jpeg format
	_ "image/png" // handle png format
	"io"
	// "io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	// "reflect"
	"strconv"
)

// IPA is an application package file for android.
type IPA struct {
	f         *os.File
	zipreader *zip.Reader
	apps      AppInfo
	name      string
}

type PyRes struct {
	Name     string
	Num      string
	PageName string
}

// OpenFile will open the file specified by filename and return IPA
func OpenFile(filename string) (ipa *IPA, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			f.Close()
		}
	}()
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	ipa, err = OpenZipReader(f, fi.Size())
	if err != nil {
		return nil, err
	}
	ipa.f = f
	ipa.name = filename
	return
}

func Icon(k *IPA) image.Image {
	var img image.Image
	Payload_ := getPath(k)
	if Payload_ == "-1" {
		fmt.Println("找不到Payload")
		return img
	}
	icopath := fmt.Sprintf("%vAppIcon76x76~ipad.png", Payload_)
	log.Println(icopath)
	m, err := k.readZipFile(icopath)
	// _, err := ioutil.ReadAll(bytes.NewReader(ico))
	ico, _, err := image.Decode(bytes.NewReader(m))
	if err != nil {
		log.Println(err, "Info.plist", "读取ICO失败")
		return img
	}
	return ico
}

func execPy(app string) PyRes {
	var py PyRes
	cmd := exec.Command("python", "app.py", app)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		// return false, apps
	}
	log.Println(string(out))
	str := strings.Split(string(out), ",")
	log.Println(str)
	Name := strings.Replace(str[0], "[", "", -1)
	log.Println(Name)
	PageName := strings.Replace(str[1], "'", "", -1)
	log.Println(PageName)
	Num := strings.Replace(str[2], "]", "", -1)
	log.Println(Num)
	Num = strings.Replace(Num, "'", "", -1)
	Num = strings.Replace(Num, "\n", "", -1)
	py.Name = strings.Replace(Name, "'", "", -1)
	py.Num = strings.Replace(Num, " ", "", -1)
	py.PageName = strings.Replace(PageName, " ", "", -1)
	return py
}

func IcoSave(Name string, img image.Image) bool {
	file, err := os.Create(Name)
	if err != nil {
		log.Println("错误：", err.Error())
		return false
	}
	defer file.Close()
	res := jpeg.Encode(file, img, &jpeg.Options{100})
	if res != nil {
		log.Println("错误：", err.Error())
		return false
	}
	return true
}

func (k *IPA) IpaInfo() (AppInfo, error) {
	var apps AppInfo
	// Payload_ := getPath(k)
	// if Payload_ == "-1" {
	// 	return apps, errors.New("找不到Payload")
	// }

	// plistpath := fmt.Sprintf("%vInfo.plist", Payload_)
	// fmt.Println("plist", plistpath)

	// xmlData, err := k.readZipFile(plistpath)
	// data, err := ioutil.ReadAll(bytes.NewReader(xmlData))

	// if err != nil {
	// 	return apps, errors.New("读取配置文件失败")
	// }
	// v := Result{}

	// newxml := NewXml{}
	// err = xml.Unmarshal(data, &newxml)
	// if err != nil {
	// 	fmt.Println("erros", err)
	// }
	// jsonstr, _ := json.MarshalIndent(newxml, "", "\t")
	// fmt.Println(string(jsonstr))
	// err = xml.Unmarshal(data, &v)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return apps, errors.New("格式化失败")
	// }
	res := execPy(k.name)

	apps.Version = res.Num
	apps.Name = res.Name
	apps.Ico = Icon(k)
	apps.PackgeName = res.PageName
	return apps, nil
}

func (k *IPA) IpaName() string {
	return k.apps.Name
}

// OpenZipReader has same arguments like zip.NewReader
func OpenZipReader(r io.ReaderAt, size int64) (*IPA, error) {
	zipreader, err := zip.NewReader(r, size)

	// for _, file := range zipreader.File {
	// 	log.Println(file.Name)
	// }
	if err != nil {
		return nil, err
	}
	ipa := &IPA{
		zipreader: zipreader,
	}

	// if err = ipa.parseManifest(); err != nil {
	// 	return nil, errors.Wrap(err, "parse-manifest")
	// }
	// if err = ipa.parseResources(); err != nil {
	// 	return nil, err
	// }
	return ipa, nil
}

// Close is avaliable only if ipa is created with OpenFile
func (k *IPA) Close() error {
	if k.f == nil {
		return nil
	}
	return k.f.Close()
}

func (k *IPA) Test() string {
	for index, file := range k.zipreader.File {
		fmt.Println(index, file.Name)
		if strings.HasSuffix(file.Name, "app/") {
			return file.Name
		}

	}

	fmt.Println(k.zipreader.File[1].Name)
	return k.zipreader.File[1].Name
}

func getPath(k *IPA) string {

	for _, file := range k.zipreader.File {
		if strings.HasSuffix(file.Name, "app/") {
			return file.Name
		}

	}
	return "-1"
}

func (k *IPA) readZipFile(name string) (data []byte, err error) {
	buf := bytes.NewBuffer(nil)

	for _, file := range k.zipreader.File {

		if file.Name != name {
			continue
		}
		rc, er := file.Open()
		if er != nil {
			err = er
			return
		}
		// log.Println(reflect.TypeOf(rc))
		defer rc.Close()
		_, err = io.Copy(buf, rc)
		if err != nil {
			return
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("File %s not found", strconv.Quote(name))
}
