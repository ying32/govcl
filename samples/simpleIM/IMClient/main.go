package main

import (
	"fmt"
	"net"

	"time"

	"errors"

	"encoding/binary"

	"github.com/ying32/govcl/samples/simpleIM"
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
	"github.com/ying32/govcl/vcl/types"
)

// 我就是测试下随便写的林

var (
	memoMsg, memoSendMsg       *vcl.TMemo
	friendList                 *vcl.TListBox
	btnSend                    *vcl.TButton
	tcpConn                    net.Conn
	nickName                   string
	mainForm, frmNickNameInput *vcl.TForm

	editNickName *vcl.TEdit
)

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.SetOnException(applicationException)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	initMainUI()
	initNickNameInput()
	initNetConn()
	vcl.Application.Run()
}

func sendMsg(cmd int, nk, msg string) error {
	if tcpConn != nil {
		bs, err := simpleIM.Encode(cmd, nk, msg)
		if err != nil {
			return err
		}
		_, err = tcpConn.Write(bs)
		if err != nil {
			return err
		}
	} else {
		return errors.New("tcpConn = nil")
	}
	return nil
}

func btnSendClick(sender vcl.IObject) {

	if !checkNKName() {
		if nickName != "" {
			mainForm.SetCaption(nickName)
			sendMsg(1001, nickName, "")
		}
		return
	}
	txt := memoSendMsg.Text()
	if txt != "" {
		if sendMsg(1000, nickName, txt) == nil {
			memoMsg.Lines().Add(fmtMessage("我", txt))
			memoSendMsg.Clear()
			memoSendMsg.SetFocus()
		}
	}
}

func mainFormKeyPress(sender vcl.IObject, key *types.Char) {
	if *key == '\r' {
		btnSend.Click()
	}
}

func mainFormDestroy(sender vcl.IObject) {
	if tcpConn != nil {
		tcpConn.Close()
	}
}

func applicationException(sender vcl.IObject, e *vcl.Exception) {
	vcl.ShowMessage(e.Message())
}

func formatdatetime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func fmtMessage(nk, msg string) string {
	// 某人说:	2018年05月06日 01:29:35
	return fmt.Sprintf("%s: %s\r\n  %s", nk, formatdatetime(time.Now()), msg)
}

func checkNKName() bool {
	if nickName == "" {
		nickName = showNickNameInput()
		return false
	}
	return true
}

func initMainUI() {
	mainForm = vcl.Application.CreateForm()
	mainForm.SetCaption("IMClient")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(800)
	mainForm.SetHeight(600)
	mainForm.SetKeyPreview(true)
	mainForm.SetOnKeyPress(mainFormKeyPress)
	mainForm.SetOnDestroy(mainFormDestroy)

	//nickName = fmt.Sprintf("随机用户%d", mainForm.Handle())
	//mainForm.SetCaption(nickName)

	//mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
	//	*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
	//})
	pnl := vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetAlign(types.AlClient)
	pnl.SetBevelOuter(types.BvNone)

	pnl2 := vcl.NewPanel(mainForm)
	pnl2.SetParent(pnl)
	pnl2.SetAlign(types.AlClient)
	pnl2.SetBevelOuter(types.BvNone)

	memoMsg = vcl.NewMemo(mainForm)
	memoMsg.SetParent(pnl2)
	memoMsg.SetAlign(types.AlClient)
	memoMsg.SetScrollBars(types.SsVertical)
	memoMsg.SetReadOnly(true)

	pnl2 = vcl.NewPanel(mainForm)
	pnl2.SetParent(pnl)
	pnl2.SetAlign(types.AlRight)
	pnl2.SetBevelOuter(types.BvNone)

	friendList = vcl.NewListBox(mainForm)
	friendList.SetParent(pnl2)
	friendList.SetAlign(types.AlClient)

	// message
	pnl = vcl.NewPanel(mainForm)
	pnl.SetParent(mainForm)
	pnl.SetAlign(types.AlBottom)
	pnl.SetHeight(200)
	pnl.SetBevelOuter(types.BvNone)

	pnl2 = vcl.NewPanel(mainForm)
	pnl2.SetParent(pnl)
	pnl2.SetAlign(types.AlClient)
	pnl2.SetBevelOuter(types.BvNone)

	memoSendMsg = vcl.NewMemo(mainForm)
	memoSendMsg.SetParent(pnl2)
	memoSendMsg.SetAlign(types.AlClient)
	memoSendMsg.SetScrollBars(types.SsVertical)

	pnl2 = vcl.NewPanel(mainForm)
	pnl2.SetParent(pnl)
	pnl2.SetAlign(types.AlBottom)
	pnl2.SetBevelOuter(types.BvNone)

	pnl3 := vcl.NewPanel(mainForm)
	pnl3.SetParent(pnl2)
	pnl3.SetAlign(types.AlRight)
	pnl3.SetBevelOuter(types.BvNone)

	btnSend = vcl.NewButton(mainForm)
	btnSend.SetParent(pnl3)
	btnSend.SetCaption("发送(&C)")
	btnSend.SetTop((pnl3.Height() - btnSend.Height()) / 2)
	btnSend.SetLeft(pnl3.Width() - btnSend.Width() - 30)
	btnSend.SetOnClick(btnSendClick)
}

func initNickNameInput() {
	frmNickNameInput = vcl.Application.CreateForm()
	frmNickNameInput.SetCaption("输入昵称：")
	frmNickNameInput.ScreenCenter()
	frmNickNameInput.SetWidth(320)
	frmNickNameInput.SetHeight(100)
	frmNickNameInput.EnabledMaximize(false)
	frmNickNameInput.EnabledMinimize(false)
	editNickName = vcl.NewEdit(frmNickNameInput)
	editNickName.SetParent(frmNickNameInput)
	editNickName.SetBounds(10, (frmNickNameInput.ClientHeight()-editNickName.Height())/2, frmNickNameInput.ClientWidth()-130, 12)
	btn := vcl.NewButton(frmNickNameInput)
	btn.SetParent(frmNickNameInput)
	btn.SetLeft(editNickName.Left() + editNickName.Width() + 10)
	btn.SetTop(editNickName.Top() - 3)
	btn.SetCaption("确定")
	btn.SetModalResult(types.MrOk)
}

func showNickNameInput() string {
	if frmNickNameInput.ShowModal() == types.MrOk {
		return editNickName.Text()
	}
	return ""
}

func initNetConn() {
	var err error
	tcpConn, err = net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		vcl.ShowMessage("连接服务器失败，错误：" + err.Error())
		return
	}
	go func() {
		for {
			buff := make([]byte, 4)
			n, err := tcpConn.Read(buff)
			if n != 4 || err != nil {
				continue
			}
			nLen := binary.LittleEndian.Uint32(buff)
			if nLen <= 0 {
				continue
			}
			buff2 := make([]byte, nLen)
			n, err = tcpConn.Read(buff2)
			if err != nil {
				continue
			}
			p, err := simpleIM.Decode(buff2)
			if err != nil {
				continue
			}
			switch p.CMD {
			case 1000:
				vcl.ThreadSync(func() {
					memoMsg.Lines().Add(fmtMessage(p.NK, p.Content))
				})
			case 1001:
				vcl.ThreadSync(func() {
					index := friendList.Items().IndexOf(p.NK)
					if index == -1 {
						friendList.Items().Add(p.NK)
					}
				})
			case 1002:
				vcl.ThreadSync(func() {
					index := friendList.Items().IndexOf(p.NK)
					if index != -1 {
						friendList.Items().Delete(index)
					}
				})
			}
		}
	}()

}
