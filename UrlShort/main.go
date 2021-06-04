package UrlShort

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"math/big"
	"strconv"
)

//生成短链接
func MakeShort(query string) string {
	base32 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5"}
	baseHex := uint(0x3FFFFFFF)
	queryHex := Md5(query)
	subHex := hexToUint("0x" + queryHex[24:])
	seed := baseHex & subHex

	result := ""
	for i := 0; i < 6; i++ {
		index := uint(0x0000001F) & seed
		result += base32[index]
		seed = seed >> 5
	}

	return result
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func hexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)

	return n
}

func hexToUint(hex string) uint {
	val := hex[2:]
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		panic(err)
	}

	return uint(n)
}
