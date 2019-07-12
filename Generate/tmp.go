package main

/*
   time: {{time}}
   authod: {{.authod}}
*/

import (
	"fmt"
	"github.com/xuchengzhi/Library/Time"
)

var times = GetTime.TF()

func main() {
	fmt.Println(times)
}
