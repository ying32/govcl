package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStyleManager struct {
}

// StyleManager_IsValidStyle
func (s *TStyleManager) IsValidStyle(filename string) bool {
	return StyleManager_IsValidStyle(filename)
}

// StyleManager_IsValidStyle2
func (s *TStyleManager) IsValidStyle2(filename string) (bool, string) {
	return StyleManager_IsValidStyle2(filename)
}

// StyleManager_LoadFromFile
func (s *TStyleManager) LoadFromFile(filename string) uintptr {
	return StyleManager_LoadFromFile(filename)
}

// StyleManager_CheckSysClassName
func (s *TStyleManager) CheckSysClassName(className string) bool {
	return StyleManager_CheckSysClassName(className)
}

// StyleManager_TrySetStyle
func (s *TStyleManager) TrySetStyle(name string, showErrorDialog bool) bool {
	return StyleManager_TrySetStyle(name, showErrorDialog)
}

// StyleManager_SetStyle1
func (s *TStyleManager) SetStyle(handle uintptr) {
	StyleManager_SetStyle1(handle)
}

// StyleManager_SetStyle2
func (s *TStyleManager) SetStyle2(name string) {
	StyleManager_SetStyle2(name)
}

// StyleManager_TryLoadFromResource
func (s *TStyleManager) TryLoadFromResource(instance uintptr, resName, resType string, handle *uintptr) bool {
	return StyleManager_TryLoadFromResource(instance, resName, resType, handle)
}

// SetStyleFromFileName
func (s *TStyleManager) SetStyleFromFileName(filename string) {
	StyleManager.SetStyle(StyleManager.LoadFromFile(filename))
}

// ActiveStyle
func (s *TStyleManager) ActiveStyle() uintptr {
	return StyleManager_ActiveStyle()
}

// SystemStyle
func (s *TStyleManager) SystemStyle() uintptr {
	return StyleManager_SystemStyle()
}

// Enabled
func (s *TStyleManager) Enabled() bool {
	return StyleManager_Enabled()
}

// IsCustomStyleActive
func (s *TStyleManager) IsCustomStyleActive() bool {
	return StyleManager_IsCustomStyleActive()
}

// UnRegisterStyle
func (s *TStyleManager) UnRegisterStyle(style uintptr) {
	StyleManager_UnRegisterStyle(style)
}

// RegisterStyle
func (s *TStyleManager) RegisterStyle(style uintptr) {
	StyleManager_RegisterStyle(style)
}

// Style
func (s *TStyleManager) Style(name string) uintptr {
	return StyleManager_Style(name)
}

// StyleNames
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
