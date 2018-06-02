/*
   这里主要是窗口资源用来反射关联事件，就可以简化相关代码。
   感觉有点蹩脚，暂无好办法替代。

   事件名命名的规则：
     On + 组件名 + 事件，相关支持的事件见： _events中
   例如窗口建完后：
     func (f *TForm1) OnForm1Create(sender vcl.IObject)
   又如按钮：
     func (f *TForm1) OnButton1Click(sender vcl.IObject)

*/

package vcl

import (
	"reflect"
	"strings"

	"fmt"

	. "github.com/ying32/govcl/vcl/types"
)

type tComponentItem struct {
	Name      string
	Component *TComponent
}

type tMethod struct {
	M    reflect.Method
	This reflect.Value
}

var (
	_eventsTable = make(map[string]tMethod, 0)
	// 这里的事件还不全，以后继续补库，先满足基本需求
	_events = map[string]struct{ call1, call2 interface{} }{
		"Click":             {onClick, nil},
		"Update":            {onUpdate, nil},
		"Execute":           {onExecute, nil},
		"MouseDown":         {onMouseDown, nil},
		"MouseUp":           {onMouseUp, nil},
		"MouseMove":         {onMouseMove, nil},
		"MouseWheel":        {onMouseWheel, nil},
		"MouseEnter":        {onMouseEnter, nil},
		"MouseLeave":        {onMouseLeave, nil},
		"CloseQuery":        {onCloseQuery, nil},
		"Destroy":           {onDestroy, nil},
		"Show":              {onShow, nil},
		"Hide":              {onHide, nil},
		"Close":             {onClose, onFormClose},
		"Find":              {onFind, nil},
		"Replace":           {onReplace, nil},
		"LinkClick":         {onLinkClick, nil},
		"Change":            {onChange, nil},
		"KeyPress":          {onKeyPress, nil},
		"DblClick":          {onDblClick, nil},
		"Resize":            {onResize, nil},
		"KeyDown":           {onKeyDown, nil},
		"KeyUp":             {onKeyUp, nil},
		"DropFiles":         {onDropFiles, nil},
		"Activate":          {onActivate, nil},
		"Deactivate":        {onDeactivate, nil},
		"ConstrainedResize": {onConstrainedResize, nil},
		"Paint":             {onPaint, nil},
		"ContextPopup":      {onContextPopup, nil},
		"DragOver":          {onDragOver, nil},
		"DragDrop":          {onDragDrop, nil},
		"StartDrag":         {onStartDrag, nil},
		"EndDrag":           {onEndDrag, nil},
		"DockDrop":          {onDockDrop, nil},
		"DockOver":          {onDockOver, nil},
		"UnDock":            {onUnDock, nil},
		"StartDock":         {onStartDock, nil},
		"GetSiteInfo":       {onGetSiteInfo, nil},
		"MouseWheelDown":    {onMouseWheelDown, nil},
		"MouseWheelUp":      {onMouseWheelUp, nil},
		"Timer":             {onTimer, nil}}

	// Application 独有事件
	_appEvents = map[string]interface{}{
		"Hint":      onHint,
		"Exception": onException,
		"Minimize":  onMinimize,
		"Restore":   onRestore,
		"Message":   onMessage,
		"Help":      onHelp,
		"ShortCut":  onShortCut}
)

func getComponents(f *TForm) (ret []tComponentItem) {
	ret = make([]tComponentItem, f.ComponentCount()+1)
	ret[0].Name = f.Name()
	ret[0].Component = ComponentFromObj(f)
	var i int32
	for i = 0; i < f.ComponentCount(); i++ {
		c := f.Components(i)
		ret[i+1].Name = c.Name()
		ret[i+1].Component = c
	}
	return
}

// addMethod 将相关方法添加到表中
func addMethod(M reflect.Method, obj IObject, eventTpy string, This reflect.Value) {
	key := fmt.Sprintf("%d_%s", obj.Instance(), eventTpy)
	var mm tMethod
	mm.M = M
	mm.This = This
	_eventsTable[key] = mm
}

// getMethod 查表，找此对象事件对应的方法
func getMethod(sender IObject, eventTpy string) (tMethod, bool) {
	key := fmt.Sprintf("%d_%s", sender.Instance(), eventTpy)
	m, ok := _eventsTable[key]
	return m, ok
}

func iifstr(b bool, atrue, afalse string) string {
	if b {
		return atrue
	}
	return afalse
}

// associatedEvents 关联事件，诶，感觉很蹩脚。。。。这方面感觉还是Delphi方便。
func associatedEvents(vv reflect.Value, comps []tComponentItem) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("associatedEvents ERROR: ", err)
		}
	}()
	// --- test
	// test test test
	// 关联事件, 查找组件
	//vv := reflect.ValueOf(inst)
	vt := vv.Type()
	eventMethods := make(map[string]reflect.Method, 0)
	for i := 0; i < vt.NumMethod(); i++ {
		m := vt.Method(i)
		if strings.HasPrefix(m.Name, "On") {
			eventMethods[m.Name] = m
		}
	}
	// 简化查找
	findMethod := func(name, event string) (reflect.Method, bool) {
		v, ok := eventMethods[fmt.Sprintf("On%s%s", name, event)]
		return v, ok
	}
	// 遍沥组件，找寻方法
	for i, item := range comps {
		if item.Name != "" {
			for k, v := range _events {
				if m, ok := findMethod(item.Name, k); ok {
					callx := v.call1
					if (i == 0) && k == "Close" {
						callx = v.call2
					}
					addNotifyEvent(vv, iifstr(i == 0, "TForm", item.Name), m, item.Component, k, callx)
				}
			}
		}
	}
	// 一些特殊的，比如Application
	addApplicationNotifyEvent(vv, eventMethods)
	// 最后调用OnCreate
	callFormCreate(vv, comps[0].Name)
}

func callFormCreate(vv reflect.Value, name string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("callFormCreate ERROR: ", err)
		}
	}()
	// 查找默认的Form创建，命名规则以 On+组件名+方法名
	m := vv.MethodByName(fmt.Sprintf("On%sCreate", name))
	if m.IsValid() {
		m.Call([]reflect.Value{vv})
	}
}

func mcall(eventType string, params ...interface{}) {
	if len(params) < 1 {
		return
	}
	// 这里要防止异常，不然全部挂了
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("mcall ERROR:", err)
		}
	}()
	sender := params[0].(IObject)
	if mmm, ok := getMethod(sender, eventType); ok {
		ps := make([]reflect.Value, len(params)+1)
		ps[0] = mmm.This
		for i, p := range params {
			ps[i+1] = reflect.ValueOf(p)
		}
		mmm.M.Func.Call(ps)
	}
}

func onClick(sender IObject) {
	mcall("Click", sender)
}

func onUpdate(sender IObject) {
	mcall("Update", sender)
}

func onExecute(sender IObject) {
	mcall("Execute", sender)
}

func onDestroy(sender IObject) {
	mcall("Destroy", sender)
}

func onCloseQuery(sender IObject, canClose *bool) {
	mcall("CloseQuery", sender, canClose)
}

func onMouseEnter(sender IObject) {
	mcall("MouseEnter", sender)
}

func onMouseLeave(sender IObject) {
	mcall("MouseLeave", sender)
}

func onShow(sender IObject) {
	mcall("Show", sender)
}

func onHide(sender IObject) {
	mcall("Hide", sender)
}

func onClose(sender IObject) {
	mcall("Close", sender)
}

func onFormClose(sender IObject, action *TCloseAction) {
	mcall("Close", sender, action)
}

func onFind(sender IObject) {
	mcall("Find", sender)
}

func onReplace(sender IObject) {
	mcall("Replace", sender)
}

func onMouseDown(sender IObject, button TMouseButton, shift TShiftState, x, y int32) {
	mcall("MouseDown", sender, button, shift, x, y)
}

func onMouseUp(sender IObject, button TMouseButton, shift TShiftState, x, y int32) {
	mcall("MouseUp", sender, button, shift, x, y)
}

func onMouseMove(sender IObject, shift TShiftState, x, y int32) {
	mcall("MouseMove", sender, shift, x, y)
}

func onMouseWheel(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool) {
	mcall("MouseWheel", sender, shift, wheelDelta, x, y, handled)
}

func onLinkClick(sender IObject, link string, linktype TSysLinkType) {
	mcall("LinkClick", sender, link, linktype)
}

func onChange(sender IObject) {
	mcall("Change", sender)
}

func onKeyPress(sender IObject, key *Char) {
	mcall("KeyPress", sender, key)
}

func onDblClick(sender IObject) {
	mcall("DblClick", sender)
}

func onResize(sender IObject) {
	mcall("Resize", sender)
}

func onKeyDown(sender IObject, key *Char, shift TShiftState) {
	mcall("KeyDown", sender, key, shift)
}

func onKeyUp(sender IObject, key *Char, shift TShiftState) {
	mcall("KeyUp", sender, key, shift)
}

func onDropFiles(sender IObject, aFileNames []string) {
	mcall("DropFiles", sender, aFileNames)
}

func onPaint(sender IObject) {
	mcall("Paint", sender)
}

func onActivate(sender IObject) {
	mcall("Activate", sender)
}

func onDeactivate(sender IObject) {
	mcall("Deactivate", sender)
}

func onConstrainedResize(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
	mcall("ConstrainedResize", sender, minWidth, minHeight, maxWidth, maxHeight)
}

func onContextPopup(sender IObject, mousePos TPoint, handled *bool) {
	mcall("ContextPopup", sender, mousePos, handled)
}

func onDragOver(sender, source IObject, x, y int32, state TDragState, accept *bool) {
	mcall("DragOver", sender, source, x, y, state, accept)
}

func onDragDrop(sender, source IObject, x, y int32) {
	mcall("DragDrop", sender, source, x, y)
}

func onStartDrag(sender IObject, dragObject *TDragObject) {
	mcall("StartDrag", sender, dragObject)
}

func onEndDrag(sender, target IObject, x, y int32) {
	mcall("EndDrag", sender, target, x, y)
}

func onDockDrop(sender IObject, source *TDragDockObject, x, y int32) {
	mcall("DockDrop", sender, x, y)
}

func onDockOver(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool) {
	mcall("DockOver", sender, source, x, y, state, accept)
}

func onUnDock(sender IObject, client *TControl, newTarget *TControl, allow *bool) {
	mcall("UnDock", sender, client, newTarget, allow)
}

func onStartDock(sender IObject, dragObject *TDragDockObject) {
	mcall("StartDock", sender, dragObject)
}

func onGetSiteInfo(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool) {
	mcall("GetSiteInfo", sender, dockClient, influenceRect, mousePos, canDock)
}

func onMouseWheelDown(sender IObject, shift TShiftState, mousePos TPoint, handled *bool) {
	mcall("MouseWheelDown", sender, shift, mousePos, handled)
}

func onMouseWheelUp(sender IObject, shift TShiftState, mousePos TPoint, handled *bool) {
	mcall("MouseWheelUp", sender, shift, mousePos, handled)
}

func onTimer(sender IObject) {
	mcall("Timer", sender)
}

// -- Application

func onException(sender IObject, e *Exception) {
	mcall("Exception", sender, e)
}
func onHint(sender IObject) {
	mcall("Hint", sender)
}

func onMinimize(sender IObject) {
	mcall("Minimize", sender)
}

func onRestore(sender IObject) {
	mcall("Restore", sender)
}

func onMessage(msg *TMsg, handled *bool) {
	mcall("Message", msg, handled)
}

func onHelp(command uint16, data THelpEventData, callhelp, result *bool) {
	mcall("Help", data, callhelp, result)
}

func onShortCut(msg *TWMKey, handled *bool) {
	mcall("ShortCut", msg, handled)
}

// callSetEventMethod 公用的call SetOnXXXX方法
func callSetEventMethod(v reflect.Value, eventType string, component IObject, this reflect.Value, m reflect.Method, fn interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("callSetEventMethod ERROR: ", err, ", eventType:", eventType, ", Component:", component, ", this:", this)
		}
	}()
	event := v.MethodByName("SetOn" + eventType)
	if event.IsValid() {
		addMethod(m, component, eventType, this)
		event.Call([]reflect.Value{reflect.ValueOf(fn)})
	}
}

// addNotifyEvent
func addNotifyEvent(vv reflect.Value, compName string, m reflect.Method, c IObject, eventType string, fn interface{}) {
	vCtl := vv.Elem().FieldByName(compName)
	if !vCtl.IsValid() {
		return
	}
	callSetEventMethod(vCtl, eventType, c, vv, m, fn)
}

// addApplicationNotifyEvent
// 添加Application的关联事件，在一个程序内，application只的事件只有最后一次设置的才会生效。
// 因为Application是单例存在，推荐在主窗口内处理就行了。
func addApplicationNotifyEvent(vv reflect.Value, eventMethods map[string]reflect.Method) {
	app := reflect.ValueOf(Application)
	if app.IsValid() {
		findMethod := func(event string) (reflect.Method, bool) {
			v, ok := eventMethods[fmt.Sprintf("OnApplication%s", event)]
			return v, ok
		}
		for eventName, fn := range _appEvents {
			if m, ok := findMethod(eventName); ok {
				callSetEventMethod(app, eventName, Application, vv, m, fn)
			}
		}
	}
}

//var typ TNotifyEvent
//var fnx = func(args []reflect.Value) []reflect.Value {
//	m.Func.Call([]reflect.Value{vv, vCtl})
//	return nil
//}
//fn := reflect.MakeFunc(reflect.TypeOf(typ), fnx)
//event.Call([]reflect.Value{fn})
