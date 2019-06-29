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
)

// autoBindEvents 自动关联事件。
func autoBindEvents(vForm reflect.Value, form *TForm, subComponenstEvent, afterBindSubComponentsEvents bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("autoBindEvents error: ", err)
		}
	}()
	// OnFormCreate
	var formCreate reflect.Value

	vt := vForm.Type()

	// 提取所有符合规则的事件
	eventMethods := make(map[string]reflect.Value, 0)
	for i := 0; i < vt.NumMethod(); i++ {
		m := vt.Method(i)
		// 保存窗口创建事件
		if m.Name == "OnFormCreate" {
			formCreate = vForm.Method(i)
			continue
		}
		if strings.HasPrefix(m.Name, "On") {
			eventMethods[m.Name] = vForm.Method(i)
		}
	}

	type tempItem struct {
		Type   string
		Method reflect.Value
	}

	// 临时方法表
	tempEventTypes := make(map[string]tempItem, 0)

	// 设置事件
	setEvent := func(component IComponent) {
		name1 := component.Name()
		name2 := name1
		if component.Equals(form) {
			name1 = "Form"
			name2 = "TForm"
		} else if component.Equals(Application) {
			name1 = "Application"
			name2 = name1
		}
		prefix := "On" + name1
		for mName, method := range eventMethods {
			if !strings.HasPrefix(mName, prefix) {
				continue
			}
			eventType := mName[len(prefix):]
			// 将事件名与事件类型对应，之后会用到的
			tempEventTypes[mName] = tempItem{eventType, method}

			if component.Equals(Application) {
				addApplicationNotifyEvent(eventType, method)
			} else {
				addComponentNotifyEvent(vForm, name2, method, eventType)
			}
		}
	}

	// 设置Form事件
	setEvent(form)

	// 子组件事件
	bindSubComponentsEvents := func() {

		var i int32
		for i = 0; i < form.ComponentCount(); i++ {
			setEvent(form.Components(i))
		}

		// 提取字段中的事件关联
		for i := 0; i < vt.Elem().NumField(); i++ {
			field := vt.Elem().Field(i)
			eventsTag := field.Tag.Get("events")
			if eventsTag == "" {
				// 兼容前面的
				eventsTag = field.Tag.Get("event")
			}
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
					findAndSetEvent(vCtl, field.Name, item.Type, item.Method)
				}
			}
		}
	}

	// 设置子组件事件
	if subComponenstEvent {
		bindSubComponentsEvents()
	}

	// 设置Application事件
	setEvent(Application)

	// 最后调用OnCreate
	callEvent(formCreate, []reflect.Value{vForm})

	// 设定了之后绑定子组件事件并且之前没有指定要绑定子组件事件
	if afterBindSubComponentsEvents && !subComponenstEvent {
		// 因为手动创建的组件没有名称，所以这里设置下，名称在当前TForm必须是唯一的
		for i := 0; i < vt.Elem().NumField(); i++ {
			field := vt.Elem().Field(i)
			if field.Type.Kind() != reflect.Ptr || field.Anonymous ||
				!strings.Contains(field.Type.String(), ".T") {
				//!strings.HasPrefix(field.Type.String(), "*vcl.") {
				continue
			}
			if vCtl := vForm.Elem().Field(i); vCtl.IsValid() {
				findAndSetComponentName(vCtl, field.Name)
			}
		}
		bindSubComponentsEvents()
	}
}

// callEvent 调用事件。
func callEvent(event reflect.Value, params []reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("callEvent error:", err)
		}
	}()
	if !event.IsValid() {
		return
	}
	event.Call(params)
}

// findAndSetEvent 公用的call SetOnXXXX方法
func findAndSetEvent(v reflect.Value, name, eventType string, method reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("findAndSetEvent error: ", err, ", eventType:", eventType)
		}
	}()
	if event := v.MethodByName("SetOn" + eventType); event.IsValid() {
		event.Call([]reflect.Value{method})
	} else {
		fmt.Printf("\"%s\" does not support the \"%s\" event.\n", name, eventType)
	}
}

// findAndSetComponentName 查找并设置组件名称
func findAndSetComponentName(v reflect.Value, name string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("findAndSetComponentName error: ", err)
		}
	}()
	if setName := v.MethodByName("SetName"); setName.IsValid() {
		setName.Call([]reflect.Value{reflect.ValueOf(name)})
	}
}

// addComponentNotifyEvent
func addComponentNotifyEvent(vForm reflect.Value, compName string, method reflect.Value, eventType string) {
	if vCtl := vForm.Elem().FieldByName(compName); vCtl.IsValid() {
		findAndSetEvent(vCtl, compName, eventType, method)
	}
}

// addApplicationNotifyEvent
// 添加Application的关联事件，在一个程序内，application中的事件只有最后一次设置的才会生效。
// 因为Application是单例存在，推荐在主窗口内处理就行了。
func addApplicationNotifyEvent(eventType string, method reflect.Value) {
	if app := reflect.ValueOf(Application); app.IsValid() {
		findAndSetEvent(app, "Application", eventType, method)
	}
}
