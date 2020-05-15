// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"strings"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TNewConnectionFields struct {
}

func (f *TNewConnection) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
}

func (f *TNewConnection) OnBtnTestConnectionClick(sender vcl.IObject) {
	name := strings.TrimSpace(f.EdtConnName.Text())
	host := strings.TrimSpace(f.EdtHost.Text())
	port := strings.TrimSpace(f.EdtPort.Text())
	password := strings.TrimSpace(f.EdtPassword.Text())

	if name == "" {
		vcl.ShowMessage("连接名得要设置个吧。")
		return
	}
	if _, ok := config[name]; ok {
		vcl.ShowMessage("配置名称已经存在，换个吧。")
		return
	}

	if host == "" {
		vcl.ShowMessage("要连接的主机名或者IP地址没设置啊。")
		return
	}
	if port == "" {
		vcl.ShowMessage("端口呢？")
		return
	}

	conn, err := newRedisConn(host, port, password)
	if err != nil {
		vcl.ShowMessage(err.Error())
		return
	}
	defer conn.Close()

	vcl.ShowMessage("测试连接成功。")
}
