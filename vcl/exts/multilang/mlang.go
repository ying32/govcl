// 多语言包，用于本地化操作

package multilang

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fmt"
	"reflect"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
)

// TLangItem 本地已经存在语言列表的项目定义
type TLangItem struct {
	Language struct {
		Id          int    // 2052
		Name        string // zh-CN
		Description string // 简体中文
		Author      string // ying32
		AuthorEmail string // 1444386932@qq.com

	} `json:"!language"`
}

var (
	// 本地已经添加了的语言列表
	LocalLangs = make(map[int]TLangItem, 0)

	// 语言存放目录
	langsPath = extractFilePath(os.Args[0]) + "Langs" + string(filepath.Separator)

	// 公共资源
	commonResouces map[string]string

	// 当前app资源
	appResouces map[string]string

	// 当前app节点信息
	appNode map[string]interface{}

	// 默认应用的节点名称
	AppNodeName string

	// 当前语言
	CurrentLang string

	// 已经注册的Form
	regForms = make(map[uintptr]vcl.IComponent, 0)

	// 需要注册的资源
	regResouces = make(map[string]*string, 0)
)

func extractFilePath(path string) string {
	filename := filepath.Base(path)
	return path[:len(path)-len(filename)]
}

func parseLangFile(lang string) {
	filename := langsPath + lang + ".lang"
	if bs, err := ioutil.ReadFile(filename); err == nil {
		var temp interface{}
		if json.Unmarshal(bs, &temp) == nil {
			// 公共资源
			commonResouces = make(map[string]string, 0)
			appResouces = make(map[string]string, 0)
			appNode = make(map[string]interface{}, 0)

			if v, ok := temp.(map[string]interface{}); ok {
				for key, val := range v["!resources"].(map[string]interface{}) {
					commonResouces[key] = val.(string)
				}
			}
			// 当前app资源
			if v, ok := temp.(map[string]interface{}); ok {
				if node, ok := v[strings.ToLower(AppNodeName)]; ok {
					appNode = node.(map[string]interface{})
					if v, ok := appNode["!resources"]; ok {
						for key, val := range v.(map[string]interface{}) {
							appResouces[key] = val.(string)
						}
					}
				}
			}
		}
	}
}

// 翻译资源，这里不UI上的资源，只是一些常量什么的
func translateStrings() {

}

func InitDefaultLang() {
	if v, ok := LocalLangs[int(rtl.SysLocale.DefaultLCID)]; ok {
		ChangeLang(v.Language.Name)
	}
}

// 改变语言
func ChangeLang(lang string) {
	if lang == CurrentLang {
		return
	}
	CurrentLang = lang
	if AppNodeName == "" {
		AppNodeName = filepath.Base(os.Args[0])
		AppNodeName = AppNodeName[:len(AppNodeName)-len(filepath.Ext(AppNodeName))]
	}
	parseLangFile(CurrentLang)
	// 翻译语言
	translateStrings()
}

// 初始一个Form的语言
func InitComponentLang(aOwner vcl.IComponent) {
	ptr := vcl.CheckPtr(aOwner)
	if ptr == 0 {
		return
	}
	if _, ok := regForms[ptr]; !ok {
		regForms[ptr] = aOwner
	}
	if node, ok := appNode[aOwner.Name()]; ok {
		for propName, propValue := range node.(map[string]interface{}) {
			propName = strings.Trim(propName, " ")
			propValue, _ := propValue.(string)
			if strings.Contains(propName, ".") {
				arr := strings.Split(propName, ".")
				if len(arr) > 1 {
					obj := aOwner.FindComponent(arr[0])
					if obj.IsValid() {
						switch len(arr) {
						case 2:
							rtl.SetPropertyValue(obj.Instance(), arr[1], propValue)
						case 3:
							rtl.SetPropertySecValue(obj.Instance(), arr[1], arr[2], propValue)
						}
					}
				}
			} else {
				rtl.SetPropertyValue(ptr, propName, propValue)
			}
		}
	}
}

// RegsiterVarString 注册需要翻译的字符
func RegsiterVarString(name string, value *string) {
	regResouces[name] = value
}

// RegsiterVarString 注册需要翻译的字符
func RegsiterVar(value *string) {
	v := reflect.TypeOf(value)
	fmt.Println("ttt:", v)
}

func initLoadLocalLangsInfo() {
	filepath.Walk(langsPath, func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(info.Name())) == ".lang" {
			bs, err := ioutil.ReadFile(path)
			if err == nil {
				item := TLangItem{}
				if json.Unmarshal(bs, &item) == nil {
					LocalLangs[item.Language.Id] = item
				}
			}
		}
		return nil
	})
}

func init() {
	initLoadLocalLangsInfo()
}
