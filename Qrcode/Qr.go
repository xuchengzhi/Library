package Qr

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
	"image"
	"image/png"
	"log"
	"os"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func builds(str, name string, x, y int) {
	base64 := str
	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Auto)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, x, y)
	if err != nil {
		log.Fatal(err)
	}

	writePng(name+".png", code)
}

func scants(file string) {
	fi, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}

// func main() {
// x, y := 500, 500
// name := "ceshi"
// str := "http://www.baidu.com"
// builds(str, name, x, y)
// scants(name + ".png")
// }
