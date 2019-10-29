package main

import (
	"flag"
	"fmt"
	// "github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/widgets"
	// "github.com/xuchengzhi/Library/Transfar"
	gojieba "github.com/xuchengzhi/gojieba"
	"github.com/xuchengzhi/wordcloud"
	"image/color"
	// "log"
	"os"
	// "reflect"
	"regexp"
	// "strings"
	"time"
)

var (
	TTF    string
	b      string
	msg    string
	bcolor string
	render bool
	h      bool
)

func init() {
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&render, "render", false, "生成云词图")
	flag.StringVar(&TTF, "TTF", "", "TTF路径")
	flag.StringVar(&msg, "msg", "", "要生成的内容")
	flag.StringVar(&b, "b", "", "背景图路径")
	flag.Usage = usage
}

func RenderNow(TTF, b_png, string, textList []string) {
	//TTF, b_png string, textList []string
	//生成云词图片
	angles := []int{0, 15, -15, 90}
	colors := []*color.RGBA{
		{0x0, 0x60, 0x30, 0xff},
		{0x60, 0x0, 0x0, 0xff},
		// &color.RGBA{0x73, 0x73, 0x0, 0xff},
	}
	// log.Println(colors)
	render := wordcloud_go.NewWordCloudRender(60, 8,
		TTF,
		b, textList, angles, colors, fmt.Sprintf("new_%v",b_png)
	render.Render()
}

func Fenci(msg string) []string {
	//使用jieba分词返回list
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	chiReg := regexp.MustCompile("[^\u4e00-\u9fa5]")
	s = chiReg.ReplaceAllString(msg, "")

	// s = "我来到北京清华大学"
	// words = x.CutAll(s)
	// fmt.Println(s)
	// fmt.Println("全模式:", strings.Join(words, " "))

	words = x.Cut(s, use_hmm)
	// fmt.Println(s)
	// fmt.Println("精确模式:", strings.Join(words, " "))

	// s = "他来到了网易杭研大厦"
	// words = x.Cut(s, use_hmm)
	// fmt.Println(s)
	// fmt.Println("新词识别:", strings.Join(words, "/"))

	// words = x.CutForSearch(s, use_hmm)
	// fmt.Println(words)
	// fmt.Println("搜索引擎模式:", reflect.TypeOf(strings.Join(words, ",")))

	// s = "长春市长春药店"
	// words = x.Tag(s)
	// fmt.Println(s)
	// fmt.Println("词性标注:", strings.Join(words, ","))
	return words
}
