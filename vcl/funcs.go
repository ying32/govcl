package vcl

import (
	"reflect"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/rtl"
	. "github.com/ying32/govcl/vcl/types"
)

// ShowMessage 显示一个消息框
func ShowMessage(msg string) {
	api.DShowMessage(msg)
}

// MessageDlg 消息框，Buttons为按钮样式，祥见types.TMsgDlgButtons
func MessageDlg(Msg string, DlgType TMsgDlgType, Buttons ...uint8) int32 {
	return api.DMessageDlg(Msg, DlgType, TMsgDlgButtons(rtl.Include(0, Buttons...)), 0)
}

// CheckPtr 检测接口是否被实例化，如果已经实例化则返回实例指针
func CheckPtr(value IObject) uintptr {
	if value == nil || reflect.ValueOf(value).Pointer() == 0 {
		return 0
	}
	return value.Instance()
}

// SelectDirectory1 选择目录
func SelectDirectory1(options TSelectDirOpts) (bool, string) {
	return api.DSelectDirectory1(options)
}

// SelectDirectory2 选择目录，一般 options默认是SdNewUI，parent默认为nil
func SelectDirectory2(caption, root string, options TSelectDirExtOpts, parent IObject) (bool, string) {
	return api.DSelectDirectory2(caption, root, options, CheckPtr(parent))
}

// ThreadSync 主线程中执行
func ThreadSync(fn TThreadProc) {
	api.DSynchronize(fn)
}

//
func InputBox(aCaption, aPrompt, aDefault string) string {
	return api.DInputBox(aCaption, aPrompt, aDefault)
}

func InputQuery(aCaption, aPrompt string, value *string) bool {
	return api.DInputQuery(aCaption, aPrompt, value)
}
