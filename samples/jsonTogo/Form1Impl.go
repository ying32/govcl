// Write your event here

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
}

// https://api.github.com/repos/ying32/govcl

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	//f.MmoOutput.SetText(f.jsonToGo(f.MmoInput.Text()))
}

func (f *TMainForm) OnPanel2Resize(sender vcl.IObject) {
	w := (f.Panel2.Width() - f.Splitter1.Width()) / 2
	f.MmoInput.SetWidth(w)
	f.MmoOutput.SetWidth(w)
}

func (f *TMainForm) OnMmoInputChange(sender vcl.IObject) {
	str, err := f.jsonToGo(f.MmoInput.Text())
	if err == nil {
		f.MmoOutput.SetText(str)
	} else {
		f.MmoOutput.SetText("错误：" + err.Error())
	}
}

// 首字母大写调整
func (f *TMainForm) firstUpCase(str string) string {
	if len(str) < 2 {
		return str
	}
	if str[0] >= 'a' && str[0] <= 'z' {
		return string(str[0]-('a'-'A')) + str[1:]
	}
	return ""
}

func (f *TMainForm) jsonToGo(str string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return "", err
	}
	switch data.(type) {
	case map[string]interface{}:
		// 主要是转以结构为准的，也就是起始为 {}
		buff := bytes.NewBufferString("type ")
		f.buildCode(buff, data, "Data", "", "")
		//buff.WriteString("}\r\n")
		return buff.String(), nil
	default:
		return "", errors.New("json起始格式必须为对象。")
	}
}

var (
	keyWords = map[string]string{
		"url": "URL"}
)

func (f *TMainForm) buildCode(buff *bytes.Buffer, data interface{}, keyName, jsonKeyName, spaceStr string) {

	if keyName != "" {
		buff.WriteString(spaceStr)
		buff.WriteString(keyName + " ")
	}
	switch data.(type) {
	case map[string]interface{}:

		buff.WriteString(fmt.Sprintf("struct {\r\n"))

		// 因为随机问题，考虑解决办法，有序总比乱序好。。。。
		keys := make([]string, 0)
		for key := range data.(map[string]interface{}) {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			val, ok := data.(map[string]interface{})[key]
			if ok {
				// 使用下划线分割
				keyName := ""
				if f.ChkRemoveUnderlineAndTryUseCamelCase.Checked() {
					for _, s := range strings.Split(key, "_") {
						if v, ok := keyWords[strings.ToLower(s)]; ok {
							keyName += v
							continue
						}
						keyName += f.firstUpCase(s)
					}
				} else {
					keyName = f.firstUpCase(key)
				}

				f.buildCode(buff, val, keyName, key, spaceStr+"    ")
			}
		}
		// 原始乱序。。。
		//for key, val := range data.(map[string]interface{}) {
		//	// 使用下划线分割
		//	keyName := ""
		//	for _, s := range strings.Split(key, "_") {
		//		if v, ok := keyWords[strings.ToLower(s)]; ok {
		//			keyName += v
		//			continue
		//		}
		//		keyName += f.firstUpCase(s)
		//	}
		//	f.buildCode(buff, val, keyName, key, spaceStr+"    ")
		//}
		buff.WriteString(spaceStr + "}")
	case []interface{}:

		buff.WriteString("[]")
		if len(data.([]interface{})) > 0 {
			f.buildCode(buff, data.([]interface{})[0], "", "", spaceStr+"    ")
		} else {
			// 空数组，不知道数据，所以使用interface{}
			buff.WriteString("interface{}")
		}

	case float64:
		// json中只有Number一值，所以这里根据有没有小数点判断使用int还是float
		// 会不太准确， 1.0也会被当成 int，大多数正确即可。
		if data.(float64)-float64(int(data.(float64))) == 0 {
			buff.WriteString("int")
		} else {
			buff.WriteString("float64")
		}

	default:
		// 无法知道其类型的一律为 interface{}
		typeName := "interface{}"
		if data != nil {
			typeName = reflect.TypeOf(data).String()
		}

		// 如果 选项开启 且 typeName == string 则尝试是否能转为int，如果是，则属性以int来做
		if typeName == "string" && f.ChkTryConvInt.Checked() {
			_, err := strconv.Atoi(data.(string))
			// 可以将值转为int
			if err == nil {
				// 改变属性类型
				typeName = "int"
				// 同时改变json属性中的
				if jsonKeyName != "" {
					jsonKeyName += ",string"
				}
			}
		}
		buff.WriteString(fmt.Sprintf("%v ", typeName))
	}
	if jsonKeyName != "" {
		buff.WriteString(fmt.Sprintf(" `json:\"%s\"`\r\n", jsonKeyName))
	}
}

func (f *TMainForm) OnBtnPasteClick(sender vcl.IObject) {
	f.MmoInput.Clear()
	f.MmoInput.PasteFromClipboard()
}

func (f *TMainForm) OnBtnCopyClick(sender vcl.IObject) {
	f.MmoOutput.SelectAll()
	f.MmoOutput.CopyToClipboard()
}
