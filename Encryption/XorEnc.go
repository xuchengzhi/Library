package XorEnc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	// "log"
	// "strings"
	"errors"
)

func XorEncodeByte(msg []byte, keys string) string {
	key := []byte(keys)
	ml := len(msg)
	kl := len(key)
	var bytes []byte
	pwd := ""
	for i := 0; i < ml; i++ {
		pwd += (string((msg[i]) ^ (key[i%kl])))
		bytes = append(bytes, (msg[i])^(key[i%kl]))
	}
	b, _ := simplifiedchinese.GBK.NewDecoder().Bytes(bytes)

	return string(b)
}

func XorEncodeStr(msg, key string) string {
	ml := len(msg)
	kl := len(key)
	// fmt.Println(string(key[ml/kl]))
	pwd := ""
	for i := 0; i < ml; i++ {
		pwd += (string((key[i%kl]) ^ (msg[i])))
	}

	return pwd
}

func XorDecodeByte(msg []byte, keys string) string {
	key := []byte(keys)
	ml := len(msg)
	kl := len(key)
	pwd := ""
	for i := 0; i < ml; i++ {
		pwd += (string(((msg[i]) ^ key[i%kl])))
	}

	return pwd
}

func XorDecodeStr(msg, key string) string {

	ml := len(msg)
	kl := len(key)
	pwd := ""
	for i := 0; i < ml; i++ {
		pwd += (string(((msg[i]) ^ key[i%kl])))
	}
	return pwd
}

func HMAC_SHA256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

// base编码
func BASE64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

func BASE64EncodeByte(src []byte) string {
	return string(base64.StdEncoding.EncodeToString(src))
}

// base解码
func BASE64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		fmt.Println("base err", err)
		return ""
	}
	return string(a)
}

var ivspec = []byte("0000000000000000")

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding1(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func AesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	content := pKCS5Padding(data, aesBlockEncrypt.BlockSize())
	cipherBytes := make([]byte, len(content))
	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, iv)
	aesEncrypt.CryptBlocks(cipherBytes, content)
	return cipherBytes, nil
}

func AesCBCDecrypt(src, key, iv []byte) ([]byte, error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err := aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return pKCS5Trimming(decrypted), nil
}

func AESBase64Decrypt(encrypt_data string, key string) (origin_data string, err error) {
	iv := []byte(key)
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		fmt.Println(err)
		return
	}
	encrypt := cipher.NewCBCDecrypter(block, iv)
	var source []byte
	if source, err = base64.StdEncoding.DecodeString(encrypt_data); err != nil {
		fmt.Println(err)
		return
	}
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	origin_data = string(PKCS5UnPadding(dst))
	return
}

func AesEncode(origData []byte, key []byte, iv []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	//采用AES加密方法中CBC加密模式
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// blocMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	//执行加密
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//实现解密
func AesDecode(cypted []byte, key []byte) (string, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypted)
	//去除填充字符串
	origData, err = PKCS7UnPadding1(origData)
	if err != nil {
		return "", err
	}
	return string(origData), err
}

//MD5加密
func Gmd5(msg string) string {
	h := md5.New()
	h.Write([]byte(msg))
	cipherStr := h.Sum(nil)
	// fmt.Println(hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr)
}

func padding(src []byte, blocksize int) []byte {
	n := len(src)
	padnum := blocksize - n%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	dst := append(src, pad...)
	return dst
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	dst := src[:n-unpadnum]
	return dst
}

func encryptDES(src []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func decryptDES(src []byte, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

//CBC加密
func EncryptDES_CBC(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//CBC解密
func DecryptDES_CBC(src, key string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

// des ECB模式加密
func EntryptDesECB(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		fmt.Println("EntryptDesECB Need a multiple of the blocksize")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

// des ECB模式解密
func DecryptDESECB(data, key []byte) string {
	if len(key) > 8 {
		key = key[:8]
	}
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		fmt.Println("DecryptDES crypto/cipher: input not full blocks")
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}
