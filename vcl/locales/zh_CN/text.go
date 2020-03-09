//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package zh_CN

import "github.com/ying32/govcl/vcl/locales"

// 默认资源
var resItems = map[string]string{
	// VCL/LCL共有
	"SOpenFileTitle":     "打开文件",
	"SOKButton":          "确定(&O)",
	"SCancelButton":      "取消(&C)",
	"SYesButton":         "是(&Y)",
	"SNoButton":          "否(&N)",
	"SHelpButton":        "帮助(&H)",
	"SCloseButton":       "关闭(&C)",
	"SIgnoreButton":      "忽略(&I)",
	"SRetryButton":       "重试(&R)",
	"SAbortButton":       "终止(&B)",
	"SAllButton":         "全部(&A)",
	"SMsgDlgWarning":     "警告",
	"SMsgDlgError":       "错误",
	"SMsgDlgInformation": "信息",
	"SMsgDlgConfirm":     "询问",
	"SMsgDlgYes":         "是",
	"SMsgDlgNo":          "否",
	"SMsgDlgOK":          "确定",
	"SMsgDlgCancel":      "取消(&C)",
	"SMsgDlgHelp":        "帮助(&H)",
	"SMsgDlgHelpNone":    "没有帮助",
	"SMsgDlgHelpHelp":    "帮助",
	"SMsgDlgAbort":       "终止(&A)",
	"SMsgDlgRetry":       "重试(&R)",
	"SMsgDlgIgnore":      "忽略(&I)",
	"SMsgDlgAll":         "全部(&A)",
	"SMsgDlgNoToAll":     "全否(&N)",
	"SMsgDlgYesToAll":    "全是(&A)",
	"SMsgDlgClose":       "关闭(&C)",

	// VCL专有
	"SSelectDirCap":  "选择目录",
	"SDirNameCap":    "目录名(&N)：",
	"SDrivesCap":     "驱动器(&R)：",
	"SDirsCap":       "目录(&D)：",
	"SFilesCap":      "文件(&F)：(*.*)",
	"SNetworkCap":    "网络(&T)...",
	"SEncodingLabel": "编码：",

	// LCL专有
	// TFindDialog And TReplaceDialog
	"rsFind":           "查找",
	"rsHelp":           "帮助",
	"rsWholeWordsOnly": "全字匹配",
	"rsCaseSensitive":  "区分大小写",
	"rsEntireScope":    "搜索整个文件",
	"rsText":           "文本",
	"rsDirection":      "方向",
	"rsForward":        "向前",
	"rsBackward":       "向后",

	"rsFindMore":   "查找更多",
	"rsReplace":    "替换",
	"rsReplaceAll": "替换所有",

	"rsSelectcolorTitle": "选择颜色",
	"rsSelectFontTitle":  "选择一个字体",
}

func init() {
	locales.ModifyResouces(resItems)
}
