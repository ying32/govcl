package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//// 车票价格
type TTicketPrice struct {
	P   string // 特等座
	A3  string // 硬卧
	A4  string // 软卧
	A1  string // 硬座
	A9  string // 商务座
	M   string // 一等座
	O   string // 二等座
	WZ  string // 无座
	MIN string // 其它
}

type TTrainInfo struct {
	TrainNo              string `json:"train_no"`               // 车次号  查票价可用
	StationTrainCode     string `json:"station_train_code"`     // 车次
	StartStationTelecode string `json:"start_station_telecode"` // 本车始发站代码
	StartStationName     string `json:"start_station_name"`     // 本车始发站名称
	EndStationTelecode   string `json:"end_station_telecode"`   // 本车终点站代码
	EndStationName       string `json:"end_station_name"`       // 本车终点站名称
	FromStationTelecode  string `json:"from_station_telecode"`  // 本次搭车起始站代码
	FromStationName      string `json:"from_station_name"`      // 本次搭车起始站名称
	ToStationTelecode    string `json:"to_station_telecode"`    // 本次搭车目标站代码
	ToStationName        string `json:"to_station_name"`        // 本次搭车目标站名
	StartTime            string `json:"start_time"`             // 发车时间
	ArriveTime           string `json:"arrive_time"`            // 到达时间
	DayDifference        int    `json:"day_difference,string"`  // 到达日差
	TrainClassName       string `json:"train_class_name"`       // 车火类型
	LiShi                string `json:"lishi"`                  // 历时　时：分
	CanWebBuy            string `json:"canWebBuy"`              // 是否支持网络构票
	LiShiValue           int    `json:"lishiValue,string"`      // 历时值
	YPInfo               string `json:"yp_info"`
	ControlTrainDay      string `json:"control_train_day"`
	StartTrainDate       string `json:"start_train_date"` // 起始日期
	SeatFeature          string `json:"seat_feature"`     // 席位的???
	YPEx                 string `json:"yp_ex"`
	TrainSeatFeature     string `json:"train_seat_feature"` // 席位特征
	SeatTypes            string `json:"seat_types"`         // 席位类型  查票价可用
	LocationCode         string `json:"location_code"`
	FromStationNo        string `json:"from_station_no"` // 查票价可用
	ToStationNo          string `json:"to_station_no"`   // 查票价可用
	ControlDay           int    `json:"control_day"`
	SaleTime             string `json:"sale_time"`
	IsSupportCard        string `json:"is_support_card"` // 是否支持直接刷2代身份证进站
	Note                 string `json:"note"`            // 备注
	GGNum                string `json:"gg_num"`          // 不知道－－－－, 这个得拼音几级才能解得出啊
	GRNum                string `json:"gr_num"`          // 高级软卧
	QTNum                string `json:"qt_num"`          // 其它
	RWNum                string `json:"rw_num"`          // 软卧
	RZNum                string `json:"rz_num"`          // 软座
	TZNnum               string `json:"tz_num"`          // 特等座
	WZNum                string `json:"wz_num"`          // 无座
	YBNum                string `json:"yb_num"`          // 不知道－－－
	YWNum                string `json:"yw_num"`          // 硬卧
	YZNum                string `json:"yz_num"`          // 硬座
	ZENum                string `json:"ze_num"`          // 二等座
	ZYNum                string `json:"zy_num"`          // 一等座
	SWZNum               string `json:"swz_num"`         // 商务座
	//price TTicketPrice //票价
}

type TTrainSearchResultData struct {
	ValidateMessagesShowId string `json:"validateMessagesShowId"`
	Status                 bool   `json:"status"`
	HttpStatus             int    `json:"httpstatus"`
	Data                   struct {
		Datas      []TTrainInfo `json:"datas"`
		Flag       bool         `json:"flag"`
		SearchDate string       `json:"searchDate"`
	} `json:"data"`
	Messages         []string `json:"messages"`
	ValidateMessages struct{} `json:"validateMessages"`
}

func trainCodeToGroupId(code string) int32 {
	if code == "" {
		return -1
	}
	switch code[0] {
	case 'G', 'g':
		return 0
	case 'D', 'd':
		return 1
	case 'Z', 'z':
		return 2
	case 'T', 't':
		return 3
	case 'K', 'k':
		return 4
	default:
		return -1
	}
}

func retStartorEnd(s1, e1, r1 string) string {
	if s1 == e1 {
		return "过"
	}
	return r1
}

func getTimeStr(n int) string {
	if n%60 != 0 {
		return fmt.Sprintf("%d时%d分", n/60, n%60)
	}
	return fmt.Sprintf("%d分", n)
}

func dayDifference(aDay int) string {
	switch aDay {
	case 0:
		return "当日到达"
	case 1:
		return "次日到达"
	case 2:
		return "两日后到达"
	case 3:
		return "三日后到达"
	case 4:
		return "四日后到达"
	default:
		return "五日之后到达"
	}
}

func parseFromFile(filename string) (*TTrainSearchResultData, error) {
	datas, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	result := new(TTrainSearchResultData)
	err = json.Unmarshal(datas, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
