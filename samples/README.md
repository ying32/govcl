## 目录下的例程并非所有都支持跨平台，只有部分基础组件支持跨平台。 

* All 

| 项目名(Project) | 说明(Description) | 平台限制(Platform limit) |
| :------ | :------ | ---- | 
| [action](action) | "动作"组件演示，演示action的公用使用 | |
| [basic](basic) | 基础窗口 | |
| [basicResForm](basicResForm) | 演示设计器输出的UI | |
| [checklistbox](checklistbox) | 复选列表框演示 | |
| [clock](clock) | 自绘的时钟 | |
| [comboboxEx](comboboxEx) | 可显示图标的combobox | |
| [customLibTest](customLibTest) | 自定义加载指定位置的liblcl | Windows, Linux |
| [draganddrop](draganddrop) | 数据拖放 | |
| [draw](draw) | canvas自绘演示 | |
| [drawchart](drawchart) | 使用[Charts for Go](https://github.com/vdobler/chart)绘制图表并显示到GoVCL的控件上 | |
| [drawfilterusegift](drawfilterusegift) | 使用[Go Image Filtering Toolkit](https://github.com/disintegration/gift)处理图片滤镜并显示到GoVCL的控件上，演示视频见`drawfilterusegift/video.mp4` | |
| [drawrose](drawrose) | canvas自绘一朵玫瑰花 | |
| [drawusegg](drawusegg) | 使用[Go Graphics - 2D](github.com/fogleman/gg)绘制并显示到GoVCL的控件上，演示视频见`drawusegg/video.mp4` | |
| [dropfiles](dropfiles) | 鼠标拖放文件演示  | |
| [eventpublic](eventpublic) | 控件事件的公用演示  | |
| [fileshelltree](fileshelltree) | 文件目录树  | |
| [formEvents](formEvents) | 演示自动关联事件    |  |
| [govcl](govcl) | 开发govcl时的总测试例程，包含很多功能的演示 | |  
| [grids/stringgrid](grids/stringgrid) | 表格控件1，常规 | |
| [grids/stringgrid2](grids/stringgrid2) | 表格控件2，高级 | |
| [grids/drawgrid](grids/drawgrid) | Draw表格控件 | |
| [imagebutton](imagebutton) | 四态图控件演示  | |
| [imageviewer](imageviewer) |  一个图片浏览器 |   |
| [inifile](inifile) | INI配置文件演示  | |
| [jsonTogo](jsonTogo) | 将一段json数据转为Go的结构，以方便json.Unmarshal填充 |  |  
| [jsonViewer](jsonViewer) | 将一段json数据以树的形式显示 |  |  
| [layout](layout) | 基础布局演示  | |
| [listboxcustomdraw](listboxcustomdraw) | ListBox部分自绘演示  | |
| [listboxcustomdraw2](listboxcustomdraw2) | ListBox部分自绘演示  | |
| [listview](listview) | ListView部分功能演示  | 部分限Windows |
| [listviewvirtualdata](listviewvirtualdata) | TListView虚拟数据，用于大数据显示 |   |
| [login](login) |  登录窗口演示  | |   
| [markdownEd](markdownEd) |  简单的markdown编辑器  | |   
| [memstream](memstream) | 内存流演示  | |
| [menu](menu) | 菜单演示  | |
| [messageTest](messageTest) | 跨平台的窗口消息捕获，与WindowsMessages例程功能一样，只是可以跨平台 | |  
| [miniwebview](miniwebview) | 跨平台浏览器组件演示 | libvcl, liblcl win32/win64,  liblcl macOS-cocoa, liblcl linux-gtk2 |  
| [mp3Player](mp3Player) | 基于bass.dll的音频播放器 |  |  
| [msgbox](msgbox) | 各种消息框演示  | |
| [multilanguage](multilanguage) | 多国语言演示例程 | |  
| [notepad](notepad) | 简单仿Windows记事本    |  |
| [OSVersion](OSVersion) | 系统版本信息获取    |  |
| [pageControlWizard](pageControlWizard) | PageControl向导程序演示,Linux与macOS下法隐藏tab，官方说明是平台限制，所以无法办到 | |  
| [printer](printer) | 打印机操作 | |  
| [redisViewer](redisViewer) | Redis视图客户端  | |  
| [res2goTest/Lazarus](res2goTest/Lazarus) | res2go工具测试例程，演示使用Lazarus设计器构建的UI | |  
| [res2goTest/Test](res2goTest/Test) | res2go工具测试例程，演示不通过Application.CreateForm创建窗口、TFrame演示| |  
| [richedit](richedit) | windows富文本框演示  |  |
| [rproxy](rproxy) | 简单的反向代理用于内网穿透，支持HTTP/HTTPS转发 | |  
| [simpleIM](simpleIM) | 简单的群聊，基于TCP    |  |
| [simplelibvlc](simplelibvlc) | 基于libvlc库的播放器，主要测试可行性 |  |  
| [statusbar](statusbar) | 状态条演示  | |
| [stdcontrols](stdcontrols) | 标准控件示例  | |
| [sysdialog](sysdialog) | 各种系统对话框演示  | |
| [syslocale](syslocale) | 本地化相关 | |
| [taskdialog](taskdialog) | 任务对话框 | |
| [trayicon](trayicon) | 系统拖盘图标演示  | Windows  MacOS 部分linux |
| [treeview](treeview) |  树型列表框演示  | |
| [valuelisteditor](valuelisteditor) |  键值编辑器 | | 
| [videosrtgui](videosrtgui) |  一个用govcl重写[video-srt-windows ](https://github.com/wxbool/video-srt-windows)的（仅UI和交互） | | 
| [simplecalc](simplecalc) | 简易计算器（演示布局用，内有布局解析）  | | 


* macOS 

| 项目名(Project) | 说明(Description) | 平台限制(Platform limit) |
| :------ | :------ | ---- | 
| [macOS/nswindowTest](macOS/nswindowTest) | 一个macOS下无标题栏样式窗口（不是无边框窗口）   | macOS,cocoa  |


* Windows

| 项目名(Project) | 说明(Description) | 平台限制(Platform limit) |
| :------ | :------ | ---- | 
| [Windows/listviewadvcustomdraw](Windows/listviewadvcustomdraw) | ListView高级自绘 | windows |  
| [Windows/listviewcustomdraw](Windows/listviewcustomdraw) |  ListView部分自绘演示  | 自绘部分限Windows |   
| [Windows/listviewitemedit](Windows/listviewitemedit) |  用于双击项目直接编辑ListView数据。  | Windows |  
| [Windows/memloaddll](Windows/memloaddll) |  内存加载dll（单文件）  | Windows 32bit |   
| [Windows/miniblinkWebview](Windows/miniblinkWebview) | 基于miniblink内核的浏览器  | Windows，目前还有问题，很多不能正常使用 |
| [Windows/registerHotkey](Windows/registerHotkey) | windows注册热键  | Windows |
| [Windows/registry](Windows/registry) | windows注册表演示  | Windows |
| [Windows/shortcut](Windows/shortcut) | 创建快捷方式相关 | Windows | 
| [Windows/treeview_checkbox](Windows/treeview_checkbox) |  树型列表框 + 复选框 演示  | windows,libvcl |
| [Windows/gdiplustest](Windows/gdiplustest) | GDI+画渐变文字及半透明背景    | Windows |
| [Windows/gdiplustest2](Windows/gdiplustest2) | 移植自Delphi IGDIPlus例程的    | Windows  |
| [Windows/WindowsMessages](Windows/WindowsMessages) | Windows下消息捕获    | Windows |
| [Windows/WindowsProcess](Windows/WindowsProcess) | Windows进程列表    | Windows |
| [Windows/windowsspy](Windows/windowsspy) | Windows下的一个窗口信息查看，类型于spy++ | Windows |
| [Windows/WindowsTest](Windows/WindowsTest) | Windows下的一些测试    | Windows |
| [Windows/winole](Windows/winole) | 使用go-ole库操作ole | Windows | 
| [Windows/wkeWebBrowser](Windows/wkeWebBrowser) | windows下wke浏览器嵌入演示  | Windows,32bit |
| [Windows/xunleidownloader](Windows/xunleidownloader) | 迅雷下载引擎演示    | Windows,32bit |
| [Windows/taskbar](Windows/taskbar) | Windows7+开始的任务栏新特性    | Windows7+ |


















