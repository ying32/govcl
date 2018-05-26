package main

import (
	"fmt"
	"net"
	"sync"

	"encoding/binary"

	"io"

	"gitee.com/ying32/govcl/samples/simpleIM"
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

// 我就是测试下随便写的林

var (
	clientMap         sync.Map
	onlineClientlbl   *vcl.TLabel
	onlineClientCount int
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("IMServer")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)
	//mainForm.SetOnCloseQuery(func(Sender vcl.IObject, CanClose *bool) {
	//	*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
	//})
	onlineClientlbl = vcl.NewLabel(mainForm)
	onlineClientlbl.SetAlign(types.AlClient)
	onlineClientlbl.SetParent(mainForm)
	onlineClientlbl.SetAutoSize(false)
	onlineClientlbl.Font().SetSize(13)
	onlineClientlbl.SetAlignment(types.TaCenter)
	onlineClientlbl.SetLayout(types.TlCenter)
	onlineClientlbl.SetCaption("当前在线人数：0")

	go initTCP()
	vcl.Application.Run()
}

func updateOnlineClient() {
	//vcl.ThreadSync(func() {
	onlineClientlbl.SetCaption(fmt.Sprintf("当前在线人数：%d", onlineClientCount))
	//})
}

func initTCP() {

	tcpConn, err := net.Listen("tcp", ":6666")
	if err != nil {
		vcl.ShowMessage(err.Error())
		return
	}
	defer tcpConn.Close()
	for {
		conn, err := tcpConn.Accept()
		if err != nil {
			fmt.Println("接受客户端连接异常：", err.Error())
			continue
		}
		fmt.Println("客户连接来自：", conn.RemoteAddr().String())
		go newConn(conn)
	}
}

func removeConn(conn net.Conn) {
	onlineClientCount--
	if v, ok := clientMap.Load(conn); ok {
		sendClientMsg(conn, simpleIM.TPacket{1002, v.(string), ""})
		updateOnlineClient()
		clientMap.Delete(conn)
	}
}

func sendClientMsg(conn net.Conn, p simpleIM.TPacket) {
	clientMap.Range(func(key, value interface{}) bool {
		if key.(net.Conn) != conn {
			bs, err := simpleIM.Encode(p.CMD, p.NK, p.Content)
			if err == nil {
				key.(net.Conn).Write(bs)
			}
		}
		return true
	})
}

func newConn(conn net.Conn) {
	clientMap.Store(conn, "")
	onlineClientCount++
	updateOnlineClient()
	defer conn.Close()
	data := make([]byte, 4)
	//conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
	for {
		n, err := conn.Read(data)
		if err == io.EOF {
			removeConn(conn)
			break
		}
		if n == 4 && err == nil {
			nLen := binary.LittleEndian.Uint32(data)
			buff := make([]byte, nLen)
			n, err = conn.Read(buff)
			if err == nil {
				p, err := simpleIM.Decode(buff)
				if err == nil {
					if p.CMD == 1001 {
						clientMap.Store(conn, p.NK)
					}
					fmt.Println("cmd:", p.CMD)
					fmt.Println("nk:", p.NK)
					fmt.Println("content:", p.Content)
					sendClientMsg(conn, p)
				} else {
					fmt.Println("3: n:", n, ", err:", err)
				}
			} else {
				if err == io.EOF {
					removeConn(conn)
					break
				}
			}
		}
	}
}
