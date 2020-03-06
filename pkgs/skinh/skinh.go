// +build windows

/*
  skinsharp皮肤扩展接口
  http://www.skinsharp.com/htdocs/index.htm
  skinsharp是收费的，这里只是头文件的导入。

  注：不知道有没有x64版本的。。。

*/
package skinh

import (
	"syscall"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

var (
	skinsharpdll = syscall.NewLazyDLL("SkinH_DL.dll")

	skinH_Attach           = skinsharpdll.NewProc("SkinH_Attach")
	skinH_AttachEx         = skinsharpdll.NewProc("SkinH_AttachEx")
	skinH_AttachExt        = skinsharpdll.NewProc("SkinH_AttachExt")
	skinH_AttachRes        = skinsharpdll.NewProc("SkinH_AttachRes")
	skinH_Detach           = skinsharpdll.NewProc("SkinH_Detach")
	skinH_DetachEx         = skinsharpdll.NewProc("SkinH_DetachEx")
	skinH_SetWindowAlpha   = skinsharpdll.NewProc("SkinH_SetWindowAlpha")
	skinH_SetMenuAlpha     = skinsharpdll.NewProc("SkinH_SetMenuAlpha")
	skinH_GetColor         = skinsharpdll.NewProc("SkinH_GetColor")
	skinH_AdjustHSV        = skinsharpdll.NewProc("SkinH_AdjustHSV")
	skinH_Map              = skinsharpdll.NewProc("SkinH_Map")
	skinH_LockUpdate       = skinsharpdll.NewProc("SkinH_LockUpdate")
	skinH_SetAero          = skinsharpdll.NewProc("SkinH_SetAero")
	skinH_SetBackColor     = skinsharpdll.NewProc("SkinH_SetBackColor")
	skinH_SetForeColor     = skinsharpdll.NewProc("SkinH_SetForeColor")
	skinH_SetWindowMovable = skinsharpdll.NewProc("SkinH_SetWindowMovable")
	skinH_AdjustAero       = skinsharpdll.NewProc("SkinH_AdjustAero")
	skinH_NineBlt          = skinsharpdll.NewProc("SkinH_NineBlt")
)

/*
{
	 	功能:	加载程序当前目录下的文件名skinh.she皮肤进行换肤
	 	返回值:	成功返回0, 失败返回非0
}
function SkinH_Attach: Integer; stdcall; external SkinH_DLL;

{
		功能:	加载指定路径的皮肤进行换肤
    参数:
          strSkinFile,	//皮肤文件路径
					strPassword		//皮肤密钥

		返回值: 成功返回0, 失败返回非0
}

*/
func Attach() int32 {
	r, _, _ := skinH_Attach.Call()
	return int32(r)
}

/*
function SkinH_AttachEx(strSkin,strPwd:PChar):Integer; stdcall; external SkinH_DLL;

{
    功能:	加载指定路径的皮肤进行换肤并指定相应的色调，饱和度，亮度,成功返回0,失败返回非0
    参数:
          strSkin,       //皮肤文件路径
          strPwd,        //皮肤密钥
          nHue,          //色调，  取值范围-180-180,默认值0
          nSat,          //饱和度，取值范围-100-100,默认值0
          nBri           //亮度，  取值范围-100-100,默认值0

    返回值: 成功返回0, 失败返回非0
}
*/
func AttachEx(strSkin, strPwd string) int32 {
	r, _, _ := skinH_AttachEx.Call(api.GoStrToDStr(strSkin), api.GoStrToDStr(strPwd))
	return int32(r)
}

/*
function SkinH_AttachExt(strSkin,strPwd:PChar; nHue:Integer=0; nSat:Integer=0; nBri:Integer=0):Integer; stdcall; external SkinH_DLL;

{
		功能:	加载指定资源进行换肤并指定相应的色调，饱和度，亮度
    参数:
						pShe,			  //资源皮肤数据指针
            dwSize,			//资源皮肤数据长度
            strPassword,//皮肤密钥
            nHue,				//色调，	取值范围-180-180, 默认值0
            nSat,				//饱和度，取值范围-100-100, 默认值0
            nBri				//亮度，	取值范围-100-100, 默认值0

    返回值: 成功返回0, 失败返回非0
}

*/
func AttachExt(strSkin, strPwd string, nHue, nSat, nBri int32) int32 {
	r, _, _ := skinH_AttachExt.Call(api.GoStrToDStr(strSkin), api.GoStrToDStr(strPwd), uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return int32(r)
}

/*
function SkinH_AttachRes(pShe:Byte; nSize:Integer; strPwd:PChar; nHue:Integer=0; nSat:Integer=0; nBri:Integer=0):Integer; stdcall; external SkinH_DLL;

{
		功能:	卸载换肤
		返回值: 成功返回0, 失败返回非0
}
*/
func AttachRes(pSh uint8, nSize int32, strPwd string, nHue, nSat, nBri int32) int32 {
	r, _, _ := skinH_AttachRes.Call(uintptr(pSh), uintptr(nSize), api.GoStrToDStr(strPwd), uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return int32(r)
}

/*
function SkinH_Detach:Integer; stdcall; external SkinH_DLL;

{
		功能:	卸载指定句柄的窗体或者控件的皮肤
    参数:
          hWnd				//指定卸载皮肤的窗体或控件的句柄

		返回值: 成功返回0, 失败返回非0
}
*/
func Detach() int32 {
	r, _, _ := skinH_Detach.Call()
	return int32(r)
}

/*
function SkinH_DetachEx(hWnd:HWND):Integer; stdcall; external SkinH_DLL;

{
		功能:	设置指定窗体的透明度

    参数:
          hWnd,				//窗体的句柄
          nAlpha			//透明度

    返回值: 成功返回0, 失败返回非0
}
*/
func DetachEx(hWnd types.HWND) int32 {
	r, _, _ := skinH_DetachEx.Call(uintptr(hWnd))
	return int32(r)
}

/*
function SkinH_SetWindowAlpha(hWnd:HWND; nAlpha:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	设置菜单透明度
		参数:
    			nAlpha				//菜单透明度，取值范围 0 - 255

    返回值: 成功返回0, 失败返回非0
}
*/
func SetWindowAlpha(hWnd types.HWND, nAlpha int32) int32 {
	r, _, _ := skinH_SetWindowAlpha.Call(uintptr(hWnd), uintptr(nAlpha))
	return int32(r)
}

/*
function SkinH_SetMenuAlpha(nAlpha:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	获取指定窗口或控件在nX,nY处的颜色值
		参数:
          hWnd,				//指定窗体或控件的句柄
					nX,					//横坐标
					nY					//纵坐标

    返回值: 成功返回0, 失败返回非0
}

*/
func SetMenuAlpha(nAlpha int32) int32 {
	r, _, _ := skinH_SetMenuAlpha.Call(uintptr(nAlpha))
	return int32(r)
}

/*
function SkinH_GetColor(hWnd:HWND; nPosX,nPosY:Integer):TColorRef; stdcall; external SkinH_DLL;

{
		功能:	调整当前皮肤的色调，饱和度，亮度
		参数:
          nHue,				//色调，	取值范围0-360, 默认值0
					nSat,				//饱和度，	取值范围0-256, 默认值0
					nBri				//亮度，	取值范围0-256, 默认值0

    返回值: 成功返回0, 失败返回非0
}
*/
func GetColor(hWnd types.HWND, nPosX, nPosY int32) uint32 {
	r, _, _ := skinH_GetColor.Call(uintptr(hWnd), uintptr(nPosX), uintptr(nPosY))
	return uint32(r)
}

/*
function SkinH_AdjustHSV(nHue,nSat,nBri:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	指定窗体和控件的换肤类型
    参数:
          hWnd,				//指定窗体或控件的句柄
          nType				//换肤类型

    返回值: 成功返回0, 失败返回非0
}
*/
func AdjustHSV(nHue, nSat, nBri int32) int32 {
	r, _, _ := skinH_AdjustHSV.Call(uintptr(nHue), uintptr(nSat), uintptr(nBri))
	return int32(r)
}

/*
function SkinH_Map(hWnd:HWND;nType:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	用于填充表格或者列表控件数据时，重复绘制影响执行效率问题
		参数:
          hWnd,				//指定窗体或控件的句柄
					bUpdate		  //1为锁定绘制，0为解锁绘制

    返回值: 成功返回0, 失败返回非0
}
*/
func Map(hWnd types.HWND, nType int32) int32 {
	r, _, _ := skinH_Map.Call(uintptr(hWnd), uintptr(nType))
	return int32(r)
}

/*
function SkinH_LockUpdate(hWnd:HWND; nLocked:Integer):Integer; stdcall; external SkinH_DLL;
{
		功能:	设置Aero特效
    参数:
				  bAero				//1为开启特效,0为关闭特效

    返回值: 成功返回0, 失败返回非0
}
*/
func LockUpdate(hWnd types.HWND, nLocked int32) int32 {
	r, _, _ := skinH_LockUpdate.Call(uintptr(hWnd), uintptr(nLocked))
	return int32(r)
}

/*
function SkinH_SetAero(hWnd:HWND):Integer; stdcall; external SkinH_DLL;

{
		功能:	设置控件的背景色(目前仅对单选框, 复选框, 分组框有效)
		参数:
          hWnd,				//控件句柄
					nRed,				//红色分量
					nGreen,			//绿色分量
					nBlue				//蓝色分量

    返回值: 成功返回0, 失败返回非0
}
*/
func SetAero(hWnd types.HWND) int32 {
	r, _, _ := skinH_SetAero.Call(uintptr(hWnd))
	return int32(r)
}

/*
function SkinH_SetBackColor(hWnd:HWND; nRed,nGreen,nBlue:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	设置控件的文本颜色色(目前仅对单选框,复选框,分组框有效)
		参数:
    			hWnd,				//控件句柄
					nRed,				//红色分量
					nGreen,			//绿色分量
					nBlue				//蓝色分量

    返回值: 成功返回0, 失败返回非0
}
*/
func SetBackColor(hWnd types.HWND, nRed, nGreen, nBlue int32) int32 {
	r, _, _ := skinH_SetBackColor.Call(uintptr(hWnd), uintptr(nRed), uintptr(nGreen), uintptr(nBlue))
	return int32(r)
}

/*
function SkinH_SetForeColor(hWnd:HWND; nRed,nGreen,nBlue:Integer):Integer; stdcall; external SkinH_DLL;

{
		功能:	设置窗体是否可以移动
    参数:
				  hWnd,				//窗口句柄
					bMovable		//0为不可移动, 1为可移动

    返回值: 成功返回0, 失败返回非0
}

*/
func SetForeColor(hWnd types.HWND, nRed, nGreen, nBlue int32) int32 {
	r, _, _ := skinH_SetForeColor.Call(uintptr(hWnd), uintptr(nRed), uintptr(nGreen), uintptr(nBlue))
	return int32(r)
}

/*
function SkinH_SetWindowMovable(hWnd:HWND;bMove:Boolean):Integer; stdcall; external SkinH_DLL;
{
		功能:	设置Aero特效参数
		参数:
    		  nAlpha,				//透明度,   0-255, 默认值0
					nShwDark,			//亮度,     0-255, 默认值0
					nShwSharp,		//锐度,	    0-255, 默认值0
					nShwSize,			//阴影大小, 2-19,  默认值2
					nX,					  //水平偏移, 0-25,  默认值0 (目前不支持)
					nY,					  //垂直偏移, 0-25,  默认值0 (目前不支持)
					nRed,				  //红色分量, 0-255, 默认值 -1
					nGreen,				//绿色分量, 0-255, 默认值 -1
					nBlue				  //蓝色分量, 0-255, 默认值 -1

		返回值: 成功返回0, 失败返回非0
}
*/
func SetWindowMovable(hWnd types.HWND, bMove bool) int32 {
	r, _, _ := skinH_SetWindowMovable.Call(uintptr(hWnd), api.GoBoolToDBool(bMove))
	return int32(r)
}

/*
function SkinH_AdjustAero(nAlpha:Integer=0; nShwDark:Integer=0; nShwSharp:Integer=0; nShwSize:Integer=2; nX:Integer=0; nY:Integer=0; nRed:Integer=-1; nGreen:Integer=-1; nBlue:Integer=-1):Integer; stdcall; external SkinH_DLL;

{
		功能:	绘制指定设备上下文的元素
		返回值: 成功返回0, 失败返回非0
    参数:
					hDtDC,				//目标设备上下文
					left,				  //左上角水平坐标
					top,				  //左上角垂直坐标
					right,				//右下角水平坐标
					bottom,				//右下角垂直坐标
					nMRect				//元素id
}

*/
func AdjustAero(nAlpha, nShwDark, nShwSharp, nShwSize, nX, nY, nRed, nGreen, nBlue int32) int32 {
	r, _, _ := skinH_AdjustAero.Call(uintptr(nAlpha), uintptr(nShwDark), uintptr(nShwSharp), uintptr(nShwSize), uintptr(nX), uintptr(nY), uintptr(nRed), uintptr(nGreen), uintptr(nBlue))
	return int32(r)
}

func AdjustAeroDefaultAll() int32 {
	return AdjustAero(0, 0, 0, 2, 0, 0, -1, -1, -1)
}

/*
function SkinH_NineBlt(hDtDC:HDC; left,top,right,bottom,nMRect:Integer):Integer; stdcall; external SkinH_DLL;
{
	错误代码

	SRET_OK             0 操作成功
	SRET_ERROR          1 操作失败
	SRET_ERROR_FILE     2 文件操作失败
	SRET_ERROR_PARAM    3 参数错误
	SRET_ERROR_CREATE   4 创建皮肤失败
	SRET_ERROR_FORMAT   5 皮肤格式错误
	SRET_ERROR_VERSION  6 皮肤版本错误
	SRET_ERROR_PASSWORD 7 皮肤密码错误
	SRET_ERROR_INVALID  8 授权失败
}
*/
func NineBlt(hDtDC types.HDC, left, top, right, bottom, nMRect int32) int32 {
	r, _, _ := skinH_NineBlt.Call(uintptr(hDtDC), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(nMRect))
	return int32(r)
}
