package main

import (
	"fmt"

	"strings"

	"os"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/win"
)

// windows下关联文件类型， 需要windows管理权限才能写入

func regAllFileType() {
	if !win.IsAdministrator() {
		fmt.Println("需要管理员权限")
		return
	}

	reg := vcl.NewRegistryAllAccess()
	defer reg.Free()
	reg.SetRootKey(win.HKEY_CLASSES_ROOT)
	menuTitle := "打开(&O)"
	exeFileName := os.Args[0]
	regFileType(reg, "jpg", "JPEG图片", menuTitle, exeFileName, 1)
	regFileType(reg, "bmp", "BMP图片", menuTitle, exeFileName, 2)
	regFileType(reg, "png", "PNG图片", menuTitle, exeFileName, 2)
	// 通知更新
	win.SHChangeNotify(win.SHCNE_ASSOCCHANGED, win.SHCNF_IDLIST, 0, 0)
}

//
func unRegAllFileType() {
	if !win.IsAdministrator() {
		fmt.Println("需要管理员权限")
		return
	}

	reg := vcl.NewRegistryAllAccess()
	defer reg.Free()
	reg.SetRootKey(win.HKEY_CLASSES_ROOT)
	unRegFileType(reg, "jpg")
	unRegFileType(reg, "bmp")
	unRegFileType(reg, "png")
	win.SHChangeNotify(win.SHCNE_ASSOCCHANGED, win.SHCNF_IDLIST, 0, 0)
}

// 注册文件类型
func regFileType(reg *vcl.TRegistry, ext, description, menuTitle, executeFileName string, iconIndex int) {
	wStr := func(aPath, aKey, aVal string) {
		if reg.OpenKey(aPath, true) {
			reg.WriteString(aKey, aVal)
			reg.CloseKey()
		}
	}
	if strings.HasPrefix(ext, ".") {
		ext = ext[1:]
	}
	wStr("."+ext, "", ext+"_file")
	wStr(ext+"_file", "", description)
	wStr(ext+"_file\\DefaultIcon", "", fmt.Sprintf("%s,%d", executeFileName, iconIndex))
	wStr(ext+"_file\\shell\\open", "", menuTitle)
	wStr(ext+"_file\\shell\\open\\command", "", fmt.Sprintf("\"%s\" \"%%1\"", executeFileName))
}

// 解除注册类型
func unRegFileType(reg *vcl.TRegistry, ext string) {
	if strings.HasPrefix(ext, "") {
		ext = ext[1:]
	}
	reg.DeleteKey("." + ext)
	reg.DeleteKey(ext + "_file")
}
