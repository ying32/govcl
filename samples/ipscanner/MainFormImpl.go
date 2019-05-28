// 在这里写你的事件

package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TMainFormFields struct {
	scanning bool
	//pool // 一个go协程池，只固定创建多少个，
	pool *sync.Pool
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.pool = &sync.Pool{
		New: func() interface{} {
			return ""
		},
	}
}

func (f *TMainForm) OnActStartExecute(sender vcl.IObject) {
	text := f.EditIPRange.Text()
	if text == "" {
		return
	}
	// 也可以用正则哦
	// (\d+)\.(\d+)\.(\d+)\.(\d+)\-(\d+)\.(\d+)\.(\d+)\.(\d+)\,(\d+)\-(\d+)
	// (\d+\.\d+\.\d+\.\d+)\-(\d+\.\d+\.\d+\.\d+)\,(\d+)\-(\d+)
	// (\d+\.\d+\.\d+\.\d+)\-(\d+\.\d+\.\d+\.\d+)\,((\d+)\-(\d+)|(\d+))

	sp := strings.Split(text, ",")
	if len(sp) != 2 {
		return
	}
	fmt.Println("IP:", sp[0])
	fmt.Println("Port:", sp[1])

	ipRange := strings.Split(sp[0], "-")
	if len(ipRange) != 2 {
		return
	}
	portRange := strings.Split(sp[1], "-")
	if len(portRange) != 2 {
		return
	}
	fmt.Println("起始IP:", ipRange[0], ", 结束IP:", ipRange[1])
	fmt.Println("起始Port:", portRange[0], ", 结束Port:", portRange[1])

	ipStart := net.ParseIP(ipRange[0]).To4()
	ipEnd := net.ParseIP(ipRange[1]).To4()
	if (ipStart[0] != ipEnd[0]) || (ipStart[1] != ipEnd[1]) || (ipStart[2] != ipEnd[2]) {
		fmt.Println("起始与结束前24位必须一致。")
		return
	}
	if ipEnd[3] < ipStart[3] {
		fmt.Println("IP范围溢出。")
		return
	}

	go func() {

		pStart, _ := strconv.Atoi(portRange[0])
		pEnd, _ := strconv.Atoi(portRange[1])
		fmt.Println("start:", pStart, ", end:", pEnd)
		for i := int(ipStart[3]); i <= int(ipEnd[3]); i++ {
			if !f.scanning {
				return
			}

			ip := net.IPv4(ipStart[0], ipStart[1], ipStart[2], byte(i))
			for n := pStart; n <= pEnd; n++ {
				if !f.scanning {
					return
				}
				url := fmt.Sprintf("http://%s:%d", ip.To4().String(), n)
				f.pool.Put(url)
			}
		}

		// 创建协程
		for i := 0; i < 2000; i++ {
			go func(n int) {
				for f.scanning {
					item := f.pool.Get().(string)
					if item != "" {
						req, _ := http.NewRequest(http.MethodHead, item, nil)
						cli := new(http.Client)
						resp, err := cli.Do(req)
						if err == nil {
							resp.Body.Close()
							fmt.Println("端口OK:", item)
						}
						continue
					} else {
						break
					}
					time.Sleep(time.Millisecond * 100)
				}
				fmt.Println("协程结束：", n)
			}(i)
		}

	}()

	f.scanning = true

}

func (f *TMainForm) OnActStartUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(!f.scanning)
}

func (f *TMainForm) OnActStopExecute(sender vcl.IObject) {
	f.scanning = false

}

func (f *TMainForm) OnActStopUpdate(sender vcl.IObject) {
	vcl.ActionFromObj(sender).SetEnabled(f.scanning)
}
