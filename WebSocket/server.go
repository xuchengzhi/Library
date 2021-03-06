package main

import (
	"fmt"

	"golang.org/x/net/websocket"

	"html/template" //支持模板html

	"log"

	"net/http"

	"io/ioutil"
)

func Echo(ws *websocket.Conn) {

	var err error

	for {

		var reply string

		//websocket接受信息

		if err = websocket.Message.Receive(ws, &reply); err != nil {

			fmt.Println("receive failed:", err)

			break

		}

		fmt.Println("reveived from client: " + reply)

		msg := "received:哈哈，" + reply

		fmt.Println("send to client:" + msg)

		//这里是发送消息

		if err = websocket.Message.Send(ws, msg); err != nil {

			fmt.Println("send failed:", err)

			break

		}

	}

}

func ReadLog(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return string(data), err
	}
	return string(data), err
}

func Logs(ws *websocket.Conn) {

	var err error

	for {

		// var reply string

		//websocket接受信息

		// if err = websocket.Message.Receive(ws, &reply); err != nil {

		// 	fmt.Println("receive failed:", err)

		// 	break

		// }

		// fmt.Println("reveived from client: " + reply)

		// msg := "received:哈哈，" + reply

		msg, _ := ReadLog("ceshi.log")

		fmt.Println("send to client:" + msg)

		//这里是发送消息

		if err = websocket.Message.Send(ws, msg); err != nil {

			fmt.Println("send failed:", err)

			break

		}

	}

}

func web(w http.ResponseWriter, r *http.Request) {

	//打印请求的方法

	fmt.Println("method", r.Method)

	if r.Method == "GET" { //如果请求方法为get显示login.html,并相应给前端

		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			panic("err")
		}
		t.Execute(w, nil)

	} else {

		//否则走打印输出post接受的参数username和password

		fmt.Println(r.PostFormValue("username"))

		fmt.Println(r.PostFormValue("password"))

	}

}

func main() {

	// 接受websocket的路由地址

	http.Handle("/websocket", websocket.Handler(Echo))

	http.Handle("/log", websocket.Handler(Logs))

	//html页面

	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal("ListenAndServe:", err)

	}
	// msg, _ := ReadLog("ceshi.log")
	// log.Println(msg)

}
