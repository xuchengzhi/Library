package main

import (
	"flag"
	"fmt"
	// "github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/widgets"
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

func renderNow() {
	//TTF, b_png string, textList []string
	textList := Fenci()
	angles := []int{0, 15, -15, 90}
	colors := []*color.RGBA{
		{0x0, 0x60, 0x30, 0xff},
		{0x60, 0x0, 0x0, 0xff},
		// &color.RGBA{0x73, 0x73, 0x0, 0xff},
	}
	// log.Println(colors)
	render := wordcloud_go.NewWordCloudRender(60, 8,
		TTF,
		b, textList, angles, colors, "foot_template.png")
	render.Render()
}

// func UiRun() {
// 	app := widgets.NewQApplication(len(os.Args), os.Args)

// 	// 创建窗口
// 	window := widgets.NewQMainWindow(nil, 0)

// 	// 设置大小
// 	window.SetMinimumSize2(500, 500)

// 	// 设置窗口标题
// 	window.SetWindowTitle("生成云词图")

// 	// 默认窗口是隐藏的，需要显示出来
// 	window.Show()
// 	app.Exec()
// }

func Fenci() []string {
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

func main() {
	startedAt := time.Now().Unix()
	// textList := []string{"恭喜", "发财", "万事", "如意"}
	// textList := Fenci("我与父亲不相见已二年余了，我最不能忘记的是他的背影。那年冬天，祖母死了，父亲的差使也交卸了，正是祸不单行的日子，我从北京到徐州，打算跟着父亲奔丧回家。到徐州见着父亲，看见满院狼藉的东西，又想起祖母，不禁簌簌地流下眼泪。父亲说，“事已如此，不必难过，好在天无绝人之路！”回家变卖典质，父亲还了亏空；又借钱办了丧事。这些日子，家中光景很是惨淡，一半为了丧事，一半为了父亲赋闲。丧事完毕，父亲要到南京谋事，我也要回北京念书，我们便同行。到南京时，有朋友约去游逛，勾留了一日；第二日上午便须渡江到浦口，下午上车北去。父亲因为事忙，本已说定不送我，叫旅馆里一个熟识的茶房陪我同去。他再三嘱咐茶房，甚是仔细。但他终于不放心，怕茶房不妥帖；颇踌躇了一会。其实我那年已二十岁，北京已来往过两三次，是没有甚么要紧的了。他踌躇了一会，终于决定还是自己送我去。我两三回劝他不必去；他只说，“不要紧，他们去不好！”我们过了江，进了车站。我买票，他忙着照看行李。行李太多了，得向脚夫行些小费，才可过去。他便又忙着和他们讲价钱。我那时真是聪明过分，总觉他说话不大漂亮，非自己插嘴不可。但他终于讲定了价钱；就送我上车。他给我拣定了靠车门的一张椅子；我将他给我做的紫毛大衣铺好坐位。他嘱我路上小心，夜里警醒些，不要受凉。又嘱托茶房好好照应我。我心里暗笑他的迂；他们只认得钱，托他们直是白托！而且我这样大年纪的人，难道还不能料理自己么？唉，我现在想想，那时真是太聪明了！我说道，“爸爸，你走吧。”他望车外看了看，说，“我买几个橘子去。你就在此地，不要走动。”我看那边月台的栅栏外有几个卖东西的等着顾客。走到那边月台，须穿过铁道，须跳下去又爬上去。父亲是一个胖子，走过去自然要费事些。我本来要去的，他不肯，只好让他去。我看见他戴着黑布小帽，穿着黑布大马褂，深青布棉袍，蹒跚地走到铁道边，慢慢探身下去，尚不大难。可是他穿过铁道，要爬上那边月台，就不容易了。他用两手攀着上面，两脚再向上缩；他肥胖的身子向左微倾，显出努力的样子。这时我看见他的背影，我的泪很快地流下来了。我赶紧拭干了泪，怕他看见，也怕别人看见。我再向外看时，他已抱了朱红的橘子望回走了。过铁道时，他先将橘子散放在地上，自己慢慢爬下，再抱起橘子走。到这边时，我赶紧去搀他。他和我走到车上，将橘子一股脑儿放在我的皮大衣上。于是扑扑衣上的泥土，心里很轻松似的，过一会说，“我走了；到那边来信！”我望着他走出去。他走了几步，回过头看见我，说，“进去吧，里边没人。”等他的背影混入来来往往的人里，再找不着了，我便进来坐下，我的眼泪又来了。近几年来，父亲和我都是东奔西走，家中光景是一日不如一日。他少年出外谋生，独力支持，做了许多大事。那知老境却如此颓唐！他触目伤怀，自然情不能自已。情郁于中，自然要发之于外；家庭琐屑便往往触他之怒。他待我渐渐不同往日。但最近两年的不见，他终于忘却我的不好，只是惦记着我，惦记着我的儿子。我北来后，他写了一信给我，信中说道，“我身体平安，惟膀子疼痛利害，举箸提笔，诸多不便，大约大去之期不远矣。”我读到此处，在晶莹的泪光中，又看见那肥胖的，青布棉袍，黑布马褂的背影。唉！我不知何时再能与他相见！")
	// TTF := "xin_shi_gu_yin.TTF"
	// b_png := "foot.png"
	flag.Parse()
	if h {
		flag.Usage()
	} else if render {
		renderNow()
	} else {
		flag.Usage()
	}

	endAt := time.Now().Unix()
	fmt.Printf("时间消耗:%d\n", endAt-startedAt)

}

func usage() {
	fmt.Fprintf(os.Stderr, `版本号: Test/1.0.0
使用: ceshi [-h 查看帮助] [-TTF TTF路径] [-b 背景图] [-msg 内容]
选项说明(咱这程序必须全套输入才可以):
`)
	flag.PrintDefaults()
	time.Sleep(3 * time.Second)
}
