package Qr

// package main

import (
	"errors"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
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

func Builds(str, fname, bcolors, colors string, x, y int) error {
	/*
	   func :生成二维码封装
	   params: str 二维码内容，fnama 二维码名称，bcolors 二维码背景颜色 ，color 二维码前景色 x,y 宽高

	*/
	base64 := str
	log.Println("二维码内容:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Auto)

	if err != nil {
		log.Println(err)
		return err
	}
	// log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")

		return errors.New("data differs")
	}

	code, err = barcode.Scale(code, x, y)
	if err != nil {
		log.Println(err)
		return err
	}
	var tmpname = "tmp.png"
	writePng(tmpname, code)
	log.Println("二维码已生成")
	errs := Changecolor(tmpname, colors, bcolors, fname+".png")
	if errs != nil {
		return errs
	}
	log.Println("二维码颜色已处理")
	return nil

}

func Scants(file string) {
	//二维码识别
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

func Changecolor(tmpname, colors, bcolors, fname string) error {
	//参考https://www.cnblogs.com/muamaker/p/10767942.html
	imgfile, err := os.Open(tmpname)

	if err != nil {
		log.Println(err)
		return err
	}
	defer imgfile.Close()

	img, err := png.Decode(imgfile)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	cimg := image.NewRGBA(bounds)
	draw.Draw(cimg, img.Bounds(), img, image.Point{}, draw.Over)

	r_new, _ := strconv.ParseUint(colors[0:2], 16, 10) //进行颜色转换
	g_new, _ := strconv.ParseUint(colors[2:4], 16, 10) //进行颜色转换
	b_new, _ := strconv.ParseUint(colors[4:6], 16, 10) //进行颜色转换

	r_bc, _ := strconv.ParseUint(bcolors[0:2], 16, 10) //进行背景颜色转换
	g_bc, _ := strconv.ParseUint(bcolors[2:4], 16, 10) //进行背景颜色转换
	b_bc, _ := strconv.ParseUint(bcolors[4:6], 16, 10) //进行背景颜色转换

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := img.At(i, j)

			r, g, b, a := colorRgb.RGBA()
			a = 255
			if r != 0 && g != 0 && b != 0 {
				cimg.Set(i, j, color.RGBA{uint8(r_bc), uint8(g_bc), uint8(b_bc), uint8(a)})
			} else if r == 0 && g == 0 && b == 0 {
				cimg.Set(i, j, color.RGBA{uint8(r_new), uint8(g_new), uint8(b_new), uint8(a)})
			}
		}
	}
	tmp, _ := os.Create(fname)
	png.Encode(tmp, cimg)
	tmp.Close()
	return nil
}

// func main() {
// 	x, y := 150, 150
// 	name := "ceshi"
// 	str := "https://itunes.apple.com/cn/app/id1473104002"
// 	bolors := "FF1493"
// 	colors := "000000"
// 	res := Builds(str, name, colors, bolors, x, y)
// 	if res != nil {
// 		log.Println(res)
// 	}
// 	// Scants(name + ".png")
// }
