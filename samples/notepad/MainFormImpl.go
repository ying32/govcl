package main

import (
	"time"

	"fmt"

	"github.com/ying32/govcl/vcl"
)

// 此方法会在将资源窗口处理完后调用，命名规则为  On+窗口Name+Create，除TForm外
func (m *TMainForm) OnFormCreate(sender vcl.IObject) {
	// 这里可以初始化些东西

}

// File

func (m *TMainForm) OnActEditCopyUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(MainForm.MemoBody.SelLength() > 0)
}

func (m *TMainForm) OnActFileNewExecute(sender vcl.IObject) {
	m.MemoBody.Clear()
}

func (m *TMainForm) OnActFileOpenExecute(sender vcl.IObject) {
	if m.DlgOpen.Execute() {
		m.MemoBody.Lines().LoadFromFile(m.DlgOpen.FileName())
	}
}

func (m *TMainForm) OnActFileSaveExecute(sender vcl.IObject) {
	if m.DlgSave.Execute() {
		m.MemoBody.Lines().SaveToFile(m.DlgSave.FileName())
	}
}

func (m *TMainForm) OnActFileSaveAsExecute(sender vcl.IObject) {
	if m.DlgSave.Execute() {
		m.MemoBody.Lines().SaveToFile(m.DlgSave.FileName())
	}
}

func (m *TMainForm) OnActFilePageSettingsExecute(sender vcl.IObject) {
	m.PageSetupDialog1.Execute()
}

func (m *TMainForm) OnActFilePrintExecute(sender vcl.IObject) {
	m.PrinterSetupDialog1.Execute()
}

func (m *TMainForm) OnActFileExitExecute(sender vcl.IObject) {
	vcl.Application.Terminate()
}

// Edit
func (m *TMainForm) OnActEditUndoExecute(sender vcl.IObject) {
	m.MemoBody.Undo()
}

func (m *TMainForm) OnActEditUndoUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(m.MemoBody.CanUndo())
}

func (m *TMainForm) OnActEditCopyExecute(sender vcl.IObject) {
	m.MemoBody.CopyToClipboard()
}

func (m *TMainForm) OnActEditCutExecute(sender vcl.IObject) {
	m.MemoBody.CutToClipboard()
}

func (m *TMainForm) OnActEditPasteExecute(sender vcl.IObject) {
	m.MemoBody.PasteFromClipboard()
}

func (m *TMainForm) OnActEditFindExecute(sender vcl.IObject) {
	m.FindDialog1.Execute()
}

func (m *TMainForm) OnFindDialog1Find(sender vcl.IObject) {
	fmt.Println("查找 查找")
}

func (m *TMainForm) OnActEditReplaceExecute(sender vcl.IObject) {
	m.ReplaceDialog1.Execute()
}

func (m *TMainForm) OnReplaceDialog1Find(sender vcl.IObject) {
	fmt.Println("替换 查找")
}

func (m *TMainForm) OnReplaceDialog1Replace(sender vcl.IObject) {
	fmt.Println("替换 替换")
}

func (m *TMainForm) OnActEditGoLineExecute(sender vcl.IObject) {
	s := vcl.InputBox("跳转行数", "行数：", "")
	if s != "" {

	}
}

func (m *TMainForm) OnActEditDeleteExecute(sender vcl.IObject) {
	m.MemoBody.SetSelText("")
}

func (m *TMainForm) OnActEditSelectAllExecute(sender vcl.IObject) {
	m.MemoBody.SelectAll()
}

func (m *TMainForm) OnActEditSelectAllUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(m.MemoBody.GetTextLen() > 0)
}

func (m *TMainForm) OnActEditTimeOrDateExecute(sender vcl.IObject) {
	m.MemoBody.Lines().Add(time.Now().Format("2006-01-02 15:04:05"))
}

// Format

func (m *TMainForm) OnActFormatWordWapExecute(sender vcl.IObject) {
	val := !vcl.AsAction(sender).Checked()
	vcl.AsAction(sender).SetChecked(val)
	m.StatusBar1.SetVisible(!val && m.ActViewStatusBar.Checked())
	m.MemoBody.SetWordWrap(val)
}

func (m *TMainForm) OnActFormatFontExecute(sender vcl.IObject) {
	if m.FontDialog1.Execute() {
		m.MemoBody.SetFont(MainForm.FontDialog1.Font())
	}
}

func (m *TMainForm) OnActViewStatusBarExecute(sender vcl.IObject) {
	val := !vcl.AsAction(sender).Checked()
	vcl.AsAction(sender).SetChecked(val)
	m.StatusBar1.SetVisible(val)
}

func (m *TMainForm) OnActViewStatusBarUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(!m.ActFormatWordWap.Checked())
}
