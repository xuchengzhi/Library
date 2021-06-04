package ipa

import (
	"encoding/xml"
	"image"
)

// Instrumentation is an application instrumentation code.
type Instrumentation struct {
	Name            string `xml:"name,attr"`
	Target          string `xml:"targetPackage,attr"`
	HandleProfiling bool   `xml:"handleProfiling,attr"`
	FunctionalTest  bool   `xml:"functionalTest,attr"`
}

// ActivityAction is an action of an activity.
type ActivityAction struct {
	Name string `xml:"name,attr"`
}

// ActivityCategory is a category of an activity.
type ActivityCategory struct {
	Name string `xml:"name,attr"`
}

// ActivityIntentFilter is an intent filter of an activity.
type ActivityIntentFilter struct {
	Actions    []ActivityAction   `xml:"action"`
	Categories []ActivityCategory `xml:"category"`
}

// AppActivity is an activity in an application.
type AppActivity struct {
	Theme         string                 `xml:"theme,attr"`
	Name          string                 `xml:"name,attr"`
	Label         string                 `xml:"label,attr"`
	IntentFilters []ActivityIntentFilter `xml:"intent-filter"`
}

type Result struct {
	XMLName xml.Name `xml:"plist"` //标签上的标签名
	// DictList   []string `xml:"dict"`
	Strings     []string `xml:"dict>string"`
	Keys        []string `xml:"dict>key"`
	String_List []string `xml:"dict>dict>dict>string"`
	Key_List    []string `xml:"dict>dict>dict>key"`
	KeyList     []string `xml:"dict>dict>key"`
	StringList  []string `xml:"dict>dict>string"`
}

type NewXml struct {
	XMLName  xml.Name `xml:"plist"`
	DictList []Dicts  `xml:"dict"`
}

type Dicts struct {
	XMLName xml.Name      `xml:"dict"`
	Key     []string      `xml:"key,string"`
	String  []string      `xml:"string"`
	Arr     []dict        `xml:"array"`
	Dict    []interface{} `xml:"dict>dict"`
}

type dict struct {
	Key    []string `xml:"key"`
	String []string `xml:"string"`
	Array  []arrs   `xml:"array"`
}
type arrs struct {
	String string `xml:"string"`
}

type AppInfo struct {
	XMLName    xml.Name `xml:"dict"`
	Name       string
	Ico        image.Image
	PackgeName string
	Version    string
}

// AppActivityAlias https://developer.android.com/guide/topics/manifest/activity-alias-element
type AppActivityAlias struct {
	Name           string                 `xml:"name,attr"`
	Label          string                 `xml:"label,attr"`
	TargetActivity string                 `xml:"targetActivity,attr"`
	IntentFilters  []ActivityIntentFilter `xml:"intent-filter"`
}

// MetaData is a metadata in an application.
type MetaData struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Application is an application in an APK.
type Application struct {
	AllowTaskReparenting  bool               `xml:"allowTaskReparenting,attr"`
	AllowBackup           bool               `xml:"allowBackup,attr"`
	BackupAgent           string             `xml:"backupAgent,attr"`
	Debuggable            bool               `xml:"debuggable,attr"`
	Description           string             `xml:"description,attr"`
	Enabled               bool               `xml:"enabled,attr"`
	HasCode               bool               `xml:"hasCode,attr"`
	HardwareAccelerated   bool               `xml:"hardwareAccelerated,attr"`
	Icon                  string             `xml:"icon,attr"`
	KillAfterRestore      bool               `xml:"killAfterRestore,attr"`
	LargeHeap             bool               `xml:"largeHeap,attr"`
	Label                 string             `xml:"label,attr"`
	Logo                  string             `xml:"logo,attr"`
	ManageSpaceActivity   string             `xml:"manageSpaceActivity,attr"`
	Name                  string             `xml:"name,attr"`
	Permission            string             `xml:"permission,attr"`
	Persistent            bool               `xml:"persistent,attr"`
	Process               string             `xml:"process,attr"`
	RestoreAnyVersion     bool               `xml:"restoreAnyVersion,attr"`
	RequiredAccountType   string             `xml:"requiredAccountType,attr"`
	RestrictedAccountType string             `xml:"restrictedAccountType,attr"`
	SupportsRtl           bool               `xml:"supportsRtl,attr"`
	TaskAffinity          string             `xml:"taskAffinity,attr"`
	TestOnly              bool               `xml:"testOnly,attr"`
	Theme                 string             `xml:"theme,attr"`
	UIOptions             string             `xml:"uiOptions,attr"`
	VMSafeMode            bool               `xml:"vmSafeMode,attr"`
	Activities            []AppActivity      `xml:"activity"`
	ActivityAliases       []AppActivityAlias `xml:"activity-alias"`
	MetaData              []MetaData         `xml:"meta-data"`
}

// UsesSDK is target SDK version.
type UsesSDK struct {
	Min    int `xml:"minSdkVersion,attr"`
	Target int `xml:"targetSdkVersion,attr"`
	Max    int `xml:"maxSdkVersion,attr"`
}

// Manifest is a manifest of an APK.
type Manifest struct {
	Package     string          `xml:"package,attr"`
	VersionCode int             `xml:"versionCode,attr"`
	VersionName string          `xml:"versionName,attr"`
	App         Application     `xml:"application"`
	Instrument  Instrumentation `xml:"instrumentation"`
	SDK         UsesSDK         `xml:"uses-sdk"`
}
