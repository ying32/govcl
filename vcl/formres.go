//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"errors"
	"reflect"
	"strings"
	"sync"
)

type formResItem struct {
	ClassName string
	Data      *[]byte
}

var (
	// 用于自动查找并加载资源的
	formResMap sync.Map
)

func getClassName(aClass interface{}) string {
	className := ""
	switch aClass.(type) {
	case string:
		className = aClass.(string)
	default:
		temp := strings.Split(reflect.TypeOf(aClass).String(), ".")
		if len(temp) > 0 {
			className = temp[len(temp)-1]
		}
	}
	if len(className) == 0 {
		return ""
	}
	return strings.ToUpper(className)
}

// 注册一个Form的资源
// 此种方式用于不指定Form资源，直接通过类名查找方式
func RegisterFormResource(aClass interface{}, data *[]byte) error {
	className := getClassName(aClass)
	if className == "" || data == nil {
		return errors.New("ClassName和data不能为空。")
	}
	// Delphi中不区分大小写的，所以统一转为大写
	formResMap.Store(className, &formResItem{ClassName: className, Data: data})
	return nil
}

// 查找对应的Form资源
func findFormResource(aClass interface{}) (*formResItem, error) {
	className := getClassName(aClass)
	if className != "" {
		if val, ok := formResMap.Load(className); ok {
			return val.(*formResItem), nil
		}
	}
	return nil, errors.New("Not found.")
}
