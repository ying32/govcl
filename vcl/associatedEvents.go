/*
   这里主要是窗口资源用来反射关联事件，就可以简化相关代码。
   感觉有点蹩脚，暂无好办法替代。

   事件名命名的规则：
     On + 组件名 + 事件，TForm特殊性，固定为名称为Form， 相关支持的事件见： _events中
   例如窗口建完后：
     func (f *TForm1) OnFormCreate(sender vcl.IObject)
   又如按钮：
     func (f *TForm1) OnButton1Click(sender vcl.IObject)

*/

package vcl

import (
	"fmt"
	"reflect"
	"strings"
)

var (

	// 这里的事件还不全，以后继续补库，先满足基本需求
	_events = []string{
		"Click",
		"Update",
		"Execute",
		"MouseDown",
		"MouseUp",
		"MouseMove",
		"MouseWheel",
		"MouseEnter",
		"MouseLeave",
		"CloseQuery",
		"Destroy",
		"Show",
		"Hide",
		"Close",
		"Find",
		"Replace",
		"LinkClick",
		"Change",
		"KeyPress",
		"DblClick",
		"Resize",
		"KeyDown",
		"KeyUp",
		"DropFiles",
		"Activate",
		"Deactivate",
		"ConstrainedResize",
		"Paint",
		"ContextPopup",
		"DragOver",
		"DragDrop",
		"StartDrag",
		"EndDrag",
		"DockDrop",
		"DockOver",
		"UnDock",
		"StartDock",
		"GetSiteInfo",
		"MouseWheelDown",
		"MouseWheelUp",
		"Timer",
		"StyleChanged",
		// Grid
		"ColumnMoved",
		"DrawCell",
		"FixedCellClick", // libvcl
		"GetEditMask",
		"GetEditText",
		"RowMoved",
		"SelectCell",
		"SetEditText",
		"TopLeftChanged",

		//headercontrol
		"DrawSection",  // livcl
		"SectionCheck", // libvcl
		"SectionClick", // libvcl
		"SectionDrag",
		"SectionEndDrag",
		"SectionResize",
		"SectionTrack"}

	// Application 独有事件
	_appEvents = []string{
		"Hint",
		"Exception",
		"Minimize",
		"Restore",
		"Message",
		"Help",
		"ShortCut"}
)

// RegisterFormEvent 这个函数主要是我没有在这里添加时，可直接注册相关事件
func RegisterFormEvent(eventType string) {
	for _, tpy := range _events {
		if strings.Compare(strings.ToLower(tpy), strings.ToLower(eventType)) == 0 {
			return
		}
	}
	_events = append(_events, eventType)
}

func iifstr(b bool, atrue, afalse string) string {
	if b {
		return atrue
	}
	return afalse
}

// associatedEvents 关联事件。
func associatedEvents(vForm reflect.Value, form *TForm) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("associatedEvents ERROR: ", err)
		}
	}()
	// test test test
	// 关联事件, 查找组件
	vt := vForm.Type()
	eventMethods := make(map[string]reflect.Value, 0)
	for i := 0; i < vt.NumMethod(); i++ {
		m := vt.Method(i)
		if strings.HasPrefix(m.Name, "On") {
			eventMethods[m.Name] = vForm.Method(i)
		}
	}
	// 简化查找
	findMethod := func(name, event string) (reflect.Value, bool) {
		v, ok := eventMethods[fmt.Sprintf("On%s%s", name, event)]
		return v, ok
	}

	setEvent := func(component IComponent, isForm bool) {
		name := iifstr(isForm, "TForm", component.Name())
		name2 := iifstr(isForm, "Form", name)
		if name == "" {
			return
		}
		for _, eventType := range _events {
			// 遍沥组件，找寻方法
			if method, ok := findMethod(name2, eventType); ok {
				// 因为TForm的特殊性，所以需要重新命名下
				addNotifyEvent(vForm, name, method, component, eventType)
			}
		}
	}

	setEvent(form, true)
	var i int32
	for i = 0; i < form.ComponentCount(); i++ {
		setEvent(form.Components(i), false)
	}

	// 一些特殊的，比如Application
	addApplicationNotifyEvent(vForm, eventMethods)
	// 最后调用OnCreate
	callFormCreate(vForm)
}

func callFormCreate(vv reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("callFormCreate ERROR: ", err)
		}
	}()
	// 查找默认的Form创建，命名规则以 On+组件名+方法名
	m := vv.MethodByName("OnFormCreate")
	if m.IsValid() {
		m.Call([]reflect.Value{vv})
	}
}

// callSetEventMethod 公用的call SetOnXXXX方法
func callSetEventMethod(v reflect.Value, eventType string, component IObject, method reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("callSetEventMethod ERROR: ", err, ", eventType:", eventType, ", Component:", component)
		}
	}()
	event := v.MethodByName("SetOn" + eventType)
	if event.IsValid() {
		//fnx := reflect.MakeFunc(method.Type(), func(args []reflect.Value) (results []reflect.Value) {
		//	return method.Call(args)
		//})
		event.Call([]reflect.Value{method})
	}
}

// addNotifyEvent
func addNotifyEvent(vv reflect.Value, compName string, method reflect.Value, component IObject, eventType string) {
	vCtl := vv.Elem().FieldByName(compName)
	if !vCtl.IsValid() {
		return
	}
	callSetEventMethod(vCtl, eventType, component, method)
}

// addApplicationNotifyEvent
// 添加Application的关联事件，在一个程序内，application只的事件只有最后一次设置的才会生效。
// 因为Application是单例存在，推荐在主窗口内处理就行了。
func addApplicationNotifyEvent(vv reflect.Value, eventMethods map[string]reflect.Value) {
	app := reflect.ValueOf(Application)
	if app.IsValid() {
		findMethod := func(event string) (reflect.Value, bool) {
			v, ok := eventMethods[fmt.Sprintf("OnApplication%s", event)]
			return v, ok
		}
		for _, eventType := range _appEvents {
			if m, ok := findMethod(eventType); ok {
				callSetEventMethod(app, eventType, Application, m)
			}
		}
	}
}
