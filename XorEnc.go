package XorEnc

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

func XorDecodeStr(msg, key string) string {
	ml := len(msg)
	kl := len(key)
	pwd := ""
	for i := 0; i < ml; i++ {
		pwd += (string(((msg[i]) ^ key[i%kl])))
	}
	return pwd
}
