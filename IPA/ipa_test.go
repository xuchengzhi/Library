package ipa

import (
	// "github.com/stretchr/testify/assert"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"testing"
)

func TestParseAPKFile(t *testing.T) {
	ipa, err := OpenFile("../testdata/1.ipa")
	if err != nil {

	}
	// if !assert.NoError(t, err) {
	// 	return
	// }
	defer ipa.Close()

	icon, err := ipa.IpaInfo()
	if err != nil {

	}
	log.Println(icon)
	// assert.NoError(t, err)
	// assert.NotNil(t, icon)

	// label, err := apk.Label(nil)
	// assert.NoError(t, err)
	// assert.Equal(t, "HelloWorld", label)
	// t.Log("app label:", label)
	// log.Println
	// assert.Equal(t, "com.example.helloworld", apk.PackageName())

	// manifest := apk.Manifest()
	// assert.Equal(t, manifest.SDK.Target, 24)

	// mainActivity, err := apk.MainActivity()
	// assert.NoError(t, err)
	// assert.Equal(t, "com.example.helloworld.MainActivity", mainActivity)

}
