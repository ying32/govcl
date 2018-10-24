## 多语言文件格式说明

多语文件为一个UTF8编码的JSON格式文本。  
```json
{
	"!language": {                    // 语言信息节点，必填
	    "Id": 2052,                   // 语言ID，必填
		"Name": "zh-CN",              // 语言名称，必填
		"Description": "简体中文",     // 语言描述名，必填
		"Author": "ying32",           // 作者名，选填
		"AuthorEmail": "1444386932@qq.com" // 作者Email，选填
	},
	"!libresources": {                // libvcl或者liblcl中的字符资源，固定项目
		"SOpenFileTitle": "打开",
		"SOKButton": "确定",
		"SCancelButton": "取消",
		"SYesButton": "是(&Y)",
		"SNoButton": "否(&N)",
		"SHelpButton": "帮助(&H)",
		"SCloseButton": "帮助(&C)",
		"SIgnoreButton": "忽略(&I)",
		"SRetryButton": "重试(&R)",
		"SAbortButton": "终止",
		"SAllButton": "全部(&A)",
		"SMsgDlgWarning": "警告",
		"SMsgDlgError": "错误",
		"SMsgDlgInformation": "信息",
		"SMsgDlgConfirm": "询问",
		"SMsgDlgYes": "是(&Y)",
		"SMsgDlgNo": "否(&N)",
		"SMsgDlgOK": "确定",
		"SMsgDlgCancel": "取消",
		"SMsgDlgHelp": "帮助(&H)",
		"SMsgDlgHelpNone": "没有有效的帮助",
		"SMsgDlgHelpHelp": "帮助",
		"SMsgDlgAbort": "终止(&A)",
		"SMsgDlgRetry": "重试(&R)",
		"SMsgDlgIgnore": "忽略(&I)",
		"SMsgDlgAll": "全部(&A)",
		"SMsgDlgNoToAll": "全否(&o)",
		"SMsgDlgYesToAll": "全是(&A)",
		"SMsgDlgClose": "关闭(&C)"
	},

	"!resources": { // 所有应用程序共享资源节点，比如app1.exe , app2.exe他们之前有相同的资源就放这里
	},

	"multilanguage": { // app节点起始，名称一般为当前应用程序名称，如果改为其他，需要在程序中指定节点
		"!resources": { // 当前app中的共享资源
			"testMessage": "这是一个测试消息！"
		},

		"Form1": { // 窗口名称，必须一致，之后在OnFormCreate中首行添加multilang.InitComponentLang(f)代码初始注册
		   "Caption": "多语言测试",  // 一级标题属性，名称都是根据控件名+属性名来操作的，这个就是直接设置窗口的了,二级类的比如 Font.Size: 12等等
		   "Button1.Caption": "按钮", // 子控件
		   "MenuItem1.Caption": "文件(&F)",
           "MenuItem2.Caption": "菜单项目2",
           "MenuItem3.Caption": "菜单项目3",
           "MenuItem4.Caption": "菜单项目4",
           "MenuItem5.Caption": "菜单项目5",
           "MenuItem6.Caption": "菜单项目6",
           "MenuItem7.Caption": "菜单项目7",
           "MenuItem8.Caption": "菜单项目8",
           "MenuItem9.Caption": "菜单项目9",
           "MenuItem10.Caption": "菜单项目10"
		}
	}
}
```
