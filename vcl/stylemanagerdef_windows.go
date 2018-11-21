package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

type TStyleManager struct {
}

// IsValidStyle 检测指定的样式文件是否有效
func (s *TStyleManager) IsValidStyle(filename string) bool {
	return StyleManager_IsValidStyle(filename)
}

// IsValidStyle2 检测指定的样式文件是否有效，并返回样式名
func (s *TStyleManager) IsValidStyle2(filename string) (bool, string) {
	return StyleManager_IsValidStyle2(filename)
}

// LoadFromFile 从文件加载样式
func (s *TStyleManager) LoadFromFile(filename string) uintptr {
	return StyleManager_LoadFromFile(filename)
}

// CheckSysClassName
func (s *TStyleManager) CheckSysClassName(className string) bool {
	return StyleManager_CheckSysClassName(className)
}

// TrySetStyle 尝试设置一个样式
func (s *TStyleManager) TrySetStyle(name string, showErrorDialog bool) bool {
	return StyleManager_TrySetStyle(name, showErrorDialog)
}

// SetStyle1 从指定句柄设置一个样式
func (s *TStyleManager) SetStyle(handle uintptr) {
	StyleManager_SetStyle1(handle)
}

// SetStyle2 从名称设置一个样式，已经注册中的
func (s *TStyleManager) SetStyle2(name string) {
	StyleManager_SetStyle2(name)
}

// TryLoadFromResource 尝试从资源中设置指定样式
func (s *TStyleManager) TryLoadFromResource(instance uintptr, resName, resType string, handle *uintptr) bool {
	return StyleManager_TryLoadFromResource(instance, resName, resType, handle)
}

// SetStyleFromFileName 从文件设置一个样式
func (s *TStyleManager) SetStyleFromFileName(filename string) {
	StyleManager.SetStyle(StyleManager.LoadFromFile(filename))
}

// ActiveStyle 激样式句柄
func (s *TStyleManager) ActiveStyle() uintptr {
	return StyleManager_ActiveStyle()
}

// SystemStyle 系统样式句柄
func (s *TStyleManager) SystemStyle() uintptr {
	return StyleManager_SystemStyle()
}

// Enabled 样式是否启用
func (s *TStyleManager) Enabled() bool {
	return StyleManager_Enabled()
}

// IsCustomStyleActive 是否自定义的样式
func (s *TStyleManager) IsCustomStyleActive() bool {
	return StyleManager_IsCustomStyleActive()
}

// UnRegisterStyle 解除已注册样式
func (s *TStyleManager) UnRegisterStyle(style uintptr) {
	StyleManager_UnRegisterStyle(style)
}

// RegisterStyle 注册一个样式
func (s *TStyleManager) RegisterStyle(style uintptr) {
	StyleManager_RegisterStyle(style)
}

// Style 从已注册样式中返回指定名称的样式句柄
func (s *TStyleManager) Style(name string) uintptr {
	return StyleManager_Style(name)
}

// StyleNames 当前已注册样式名称数组
func (s *TStyleManager) StyleNames() (arr []string) {
	var i int32
	for {
		s := StyleManager_StyleNamesOf(i)
		if s == "" {
			return
		}
		arr = append(arr, s)
		i++
	}
	return
}
