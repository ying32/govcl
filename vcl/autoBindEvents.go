//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

/*
   这里主要是窗口资源用来反射关联事件，就可以简化相关代码。

   事件名命名的规则：
     On + 组件名 + 事件，TForm特殊性，固定为名称为Form。
   例如窗口建完后：
     func (f *TForm1) OnFormCreate(sender vcl.IObject)
   又如按钮：
     func (f *TForm1) OnButton1Click(sender vcl.IObject)

   原理：
     首次会收集Form以On开头的事件，然后根据 组件名称提取出事件的类型，再通过事件类型查找某个组件中的 SetOn + eventType方法。

   多个组件共享同一个事件：

   type TMainForm struct {
       *vcl.TForm
       Button1 *vcl.TButton
       Button2 *vcl.TButton `events:"OnButton1Click"`
       Button3 *vcl.TButton `events:"OnButton1Click,OnButton1Resize"`
   }

   // 这样只自动关联了Button1的事件，但此时我想将此事件关联到Button2, Button3上
   // 常规的做法就是 Button2.SetOnClick(f.OnButton1Click)
   // 现在提供一种新的方式，这种方式应对于res2go转换后不自动共享事件问题。

   func (f *TMainForm) OnButton1Click(sender vcl.IObject) {

   }

*/

package vcl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ying32/govcl/vcl/api"
)

type eventMethod struct {
	Method  reflect.Value
	FuncPtr uintptr
}

// autoBindEvents 自动关联事件。
func autoBindEvents(vForm reflect.Value, root IComponent, subComponentsEvent, afterBindSubComponentsEvents bool) {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Calling autoBindEvents exception:", err)
			}
		}()
	}
	// OnFormCreate or OnFrameCreate
	var doCreate reflect.Value

	vt := vForm.Type()

	// 提取所有符合规则的事件
	eventMethods := make(map[string]eventMethod, 0)
	// 遍历当前结构的方法
	for i := 0; i < vt.NumMethod(); i++ {
		m := vt.Method(i)
		// 保存窗口创建事件
		if m.Name == "OnFormCreate" || m.Name == "OnFrameCreate" {
			doCreate = vForm.Method(i)
			continue
		}
		if strings.HasPrefix(m.Name, "On") {
			eventMethods[m.Name] = eventMethod{Method: vForm.Method(i), FuncPtr: m.Func.Pointer()}
		}
	}

	type eventItem struct {
		Type   string
		Method eventMethod
	}

	// 临时方法表
	tempEventTypes := make(map[string]eventItem, 0)

	// 遍历结构中的字段
	//for i := 0; i < vt.Elem().NumField(); i++ {
	//	field := vt.Elem().Field(i)
	//	fmt.Println("field.Name:", field.Name)
	//
	//}

	// 用于之后显示提示的
	formName := root.Name()

	// 设置事件
	setEvent := func(component IComponent) {
		name1 := component.Name()
		name2 := name1
		if component.Equals(root) {
			name1 = "Form"
			name2 = "TForm"
			if root.ClassName() == "TFrame" {
				name1 = "Frame"
				name2 = "TFrame"
			}
		} else if component.Equals(Application) {
			name1 = "Application"
			name2 = name1
		}
		// 前缀 On + 组件名
		prefix := "On" + name1

		for mName, method := range eventMethods {
			if !strings.HasPrefix(mName, prefix) {
				continue
			}
			eventType := mName[len(prefix):]
			// 将事件名与事件类型对应，之后会用到的
			tempEventTypes[mName] = eventItem{eventType, method}

			if component.Equals(Application) {
				addApplicationNotifyEvent(eventType, method)
			} else {
				addComponentNotifyEvent(vForm, name2, method, eventType, formName)
			}
		}
	}

	// 设置Root事件
	setEvent(root)

	// 子组件事件
	bindSubComponentsEvents := func() {
		var i int32
		for i = 0; i < root.ComponentCount(); i++ {
			setEvent(root.Components(i))
		}

		// 提取字段中的事件关联
		for i := 0; i < vt.Elem().NumField(); i++ {
			field := vt.Elem().Field(i)
			eventsTag := field.Tag.Get("events")
			if eventsTag == "" {
				continue
			}
			eventArr := strings.Split(eventsTag, ",")
			for _, event := range eventArr {
				event = strings.TrimSpace(event)
				item, ok := tempEventTypes[event]
				if !ok {
					continue
				}
				if vCtl := vForm.Elem().Field(i); vCtl.IsValid() {
					findAndSetEvent(vCtl, field.Name, item.Type, item.Method, formName)
				}
			}
		}
	}

	// 设置子组件事件
	if subComponentsEvent {
		bindSubComponentsEvents()
	}

	// 设置Application事件
	setEvent(Application)

	// 最后调用OnCreate
	callEvent(doCreate, []reflect.Value{vForm})

	// 设定了之后绑定子组件事件并且之前没有指定要绑定子组件事件
	// 会造成冲突，先禁用吧
	//if afterBindSubComponentsEvents && !subComponentsEvent {
	//	// 因为手动创建的组件没有名称，所以这里设置下，名称在当前TForm必须是唯一的
	//	for i := 0; i < vt.Elem().NumField(); i++ {
	//		field := vt.Elem().Field(i)
	//		if field.Type.Kind() != reflect.Ptr || field.Anonymous ||
	//			!strings.Contains(field.Type.String(), ".T") {
	//			continue
	//		}
	//		// 检测首字母是否大写
	//		if len(field.Name) >= 1 {
	//			// 首字母不为A-Z之间的则排除。
	//			if c := field.Name[0]; !(c >= 'A' && c <= 'Z') {
	//				continue
	//			}
	//		}
	//		if vCtl := vForm.Elem().Field(i); vCtl.IsValid() {
	//			findAndSetComponentName(vCtl, field.Name, true)
	//		}
	//	}
	//	bindSubComponentsEvents()
	//}
}

// callEvent 调用事件。
func callEvent(event reflect.Value, params []reflect.Value) {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Calling callEvent exception:", err)
			}
		}()
	}
	if !event.IsValid() {
		return
	}
	event.Call(params)
}

// findAndSetEvent 公用的call SetOnXXXX方法
func findAndSetEvent(v reflect.Value, name, eventType string, method eventMethod, rootName string) {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Calling findAndSetEvent exception:", err, ", eventType:", eventType, ", rootName:", rootName)
			}
		}()
	}
	if event := v.MethodByName("SetOn" + eventType); event.IsValid() {
		// 设置EventId
		api.BeginAddEvent()
		defer api.EndAddEvent()
		api.SetCurrentEventId(api.GetUID(v.Pointer(), method.FuncPtr))
		event.Call([]reflect.Value{method.Method})
	} else {
		if len(eventType) > 0 {
			// 也许分析错误，所不打印错误消息。
			s := eventType[0]
			switch {
			case s >= '0' && s <= '9' || s == '_':
				return
			}
		}
		fmt.Printf("%s.%s does not support the %s event.\n", rootName, name, eventType)
	}
}

// findAndSetComponentName 查找并设置组件名称
func findAndSetComponentName(v reflect.Value, name string, clearDefault bool) {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Calling findAndSetComponentName exception:", err)
			}
		}()
	}
	if v.Pointer() == 0 {
		return
	}
	if setName := v.MethodByName("SetName"); setName.IsValid() {
		setName.Call([]reflect.Value{reflect.ValueOf(name)})
		if clearDefault {
			if setText := v.MethodByName("SetText"); setText.IsValid() {
				setText.Call([]reflect.Value{reflect.ValueOf("")})
			}
		}
	}
}

// addComponentNotifyEvent
func addComponentNotifyEvent(vForm reflect.Value, compName string, method eventMethod, eventType, rootName string) {
	if vCtl := vForm.Elem().FieldByName(compName); vCtl.IsValid() {
		findAndSetEvent(vCtl, compName, eventType, method, rootName)
	}
}

// addApplicationNotifyEvent
// 添加Application的关联事件，在一个程序内，application中的事件只有最后一次设置的才会生效。
// 因为Application是单例存在，推荐在主窗口内处理就行了。
func addApplicationNotifyEvent(eventType string, method eventMethod) {
	if app := reflect.ValueOf(Application); app.IsValid() {
		findAndSetEvent(app, "Application", eventType, method, "Application")
	}
}
