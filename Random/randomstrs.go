package Randoms

import (
	// "log"
	"math/rand"
	"time"
)

//获取随机数字
func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//获取随机字符串
func GetRandomString(num int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// func main() {
// 	msg := GetRandomInt(10000, 99999) //GetRandomString(18)
// 	log.Println(msg)
// }
