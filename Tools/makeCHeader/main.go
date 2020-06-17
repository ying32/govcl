//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	paramsExpr  = regexp.MustCompile(`\((.+?)\)`)
	incFileExpr = regexp.MustCompile(`\{\$I\s(MyLCL\_.+?\.inc)\}`)
	typeMap     = map[string]string{
		"LongBool":             "BOOL",
		"PChar":                "char*",
		"Pointer":              "void*",
		"Integer":              "int32_t",
		"Cardinal":             "uint32_t",
		"Double":               "double",
		"Single":               "float",
		"TThreadID":            "uintptr_t",
		"NativeInt":            "intptr_t",
		"NativeUInt":           "uintptr_t",
		"Boolean":              "BOOL",
		"Longint":              "int32_t",
		"UInt64":               "uint64_t",
		"IInt64":               "int64_t",
		"TBasicAction":         "TAction",
		"TPersistent":          "TObject",
		"Byte":                 "char",
		"TBorderWidth":         "int32_t",
		"TCustomImageList":     "TImageList",
		"TUnixDateTime":        "uint32_t",
		"TCustomListView":      "TListView",
		"TCustomTreeView":      "TTreeView",
		"TImageIndex":          "int32_t",
		"SmallInt":             "int16_t",
		"TCustomForm":          "TForm",
		"TWidth":               "int32_t",
		"uint32":               "uint32_t",
		"int32":                "int32_t",
		"Int64":                "int64_t",
		"LongWord":             "uint32_t",
		"Word":                 "uint16_t",
		"PPointerList":         "void*",
		"string":               "char*",
		"TCustomHeaderControl": "THeaderControl",
		"TCollectionItemClass": "TCollectionItem",
		"PPoint":               "TPoint*",
		"TOverlay":             "uint8_t",
		"TGoForm":              "TForm",
		"TLMessage":            "TMessage",

		//"TResItem":             "",
	}
	funcsMap   = make(map[string]string, 0)
	classArray = make([]string, 0)
)

func main() {
	//file := NewCFile("liblcl.h")
	file := NewCFile(`F:\Users\zhuying\Documents\Visual Studio 2017\Projects\ConsoleApplication3\ConsoleApplication3\liblcl.h`)
	file.WriteHeader()

	// UInt64 WINAPI MySyscall(void* AProc, intptr_t ALen, void* A1, void* A2, void* A3, void* A4, void* A5, void* A6, void* A7, void* A8, void* A9, void* A10, void* A11, void* A12);
	funcsMap["MySyscall"] = "" // 排除此函数，手动构建
	funcsMap["NSWindow_titleVisibility"] = ""
	funcsMap["NSWindow_setTitleVisibility"] = ""
	funcsMap["NSWindow_titlebarAppearsTransparent"] = ""
	funcsMap["NSWindow_setTitlebarAppearsTransparent"] = ""
	funcsMap["NSWindow_styleMask"] = ""
	funcsMap["NSWindow_setStyleMask"] = ""
	funcsMap["NSWindow_setRepresentedURL"] = ""
	funcsMap["NSWindow_release"] = ""

	file.WLn()
	file.WLn()

	// 自动生成的对象函数
	parseClassFiles(file, "uexport1.pas")
	parseClassFiles(file, "uexport2.pas")

	file.WLn()
	parseFile(file, "LazarusDef.inc")

	file.WLn()
	file.W("#ifdef __linux__\n")
	parseFile(file, "ulinuxpatchs.pas")
	file.WLn()
	file.W("#endif\n")
	file.WLn()

	file.WLn()
	file.W("#ifdef __APPLE__\n")
	parseFile(file, "umacospatchs.pas")
	file.WLn()
	file.W("#endif\n")
	file.WLn()

	parseFile(file, "uformdesignerfile.pas")

	// ulinuxpatchs.pas
	// umacospathcs.pas

	//fmt.Println(file.String())

	file.WriteFooter()
	// 生成枚举
	file.WLn()
	file.WComment("枚举定义\n")

	file.Save(classArray, parseEnums("../../vcl/types/enums.go"))

}

func parseFile(f *CFile, fileName string) error {
	bs, err := ioutil.ReadFile("../../UILibSources/liblcl/" + fileName)
	if err != nil {
		return err
	}
	f.WComment(fileName)
	// {无效参数}
	bs = bytes.Replace(bs, []byte("{无效参数}"), nil, -1)
	bs = bytes.Replace(bs, []byte("\r"), nil, -1)
	sps := bytes.Split(bs, []byte("\n"))
	for _, line := range sps {
		s := string(bytes.TrimSpace(line))
		if (strings.HasPrefix(strings.ToLower(s), "function") || strings.HasPrefix(strings.ToLower(s), "procedure")) && strings.HasSuffix(s, "extdecl;") {
			parseFunc(f, s)
		}
	}

	return nil
}

func parseClassFiles(f *CFile, fileName string) error {
	bs, err := ioutil.ReadFile("../../UILibSources/liblcl/" + fileName)
	if err != nil {
		return err
	}
	bs = bytes.Replace(bs, []byte("\r"), nil, -1)
	f.WComment(fileName)
	matchs := incFileExpr.FindAllStringSubmatch(string(bs), -1)
	for _, match := range matchs {
		if len(match) >= 2 {
			incFileName := strings.TrimSpace(match[1])
			className := strings.TrimPrefix(incFileName, "MyLCL_")
			className = strings.TrimSuffix(className, ".inc")
			if className != "Exception" {
				className = "T" + strings.Trim(className, " ")
			} else {
				className = strings.Trim(className, " ")
			}

			classArray = append(classArray, className)
			fmt.Println(incFileName)
			parseFile(f, incFileName)
		}
	}
	return nil
}

func parseFunc(f *CFile, s string) error {
	isFunc := strings.HasPrefix(strings.ToLower(s), "function")
	if isFunc {
		s = strings.TrimPrefix(s, "function")
	} else {
		s = strings.TrimPrefix(s, "procedure")
	}
	s = strings.TrimSpace(strings.TrimSuffix(s, "extdecl;"))

	haveParams := strings.Index(s, "(") != -1
	paramsStr := ""
	if haveParams {
		ms := paramsExpr.FindStringSubmatch(s)
		if len(ms) >= 2 {
			paramsStr = strings.TrimSpace(ms[1])

		}
	}

	funcName := ""
	n := strings.Index(s, "(")
	if n == -1 {
		n = strings.Index(s, ":")
	}
	if n != -1 {
		funcName = strings.TrimSpace(s[:n])
	}

	returnType := ""
	if isFunc {
		n := strings.Index(s, ")")
		if n == -1 {
			n = strings.Index(s, ":")
		}
		if n != -1 {
			returnType = strings.TrimSpace(s[n+1:])
			returnType = strings.TrimSuffix(returnType, ";")
			returnType = strings.TrimPrefix(returnType, ":")
			returnType = strings.TrimSpace(returnType)
		}
	}

	//fmt.Println("name: ", funcName, ", returnType:", returnType, ", isFunc: ", isFunc, ", haveParams: ", haveParams, ", params: ", paramsStr)
	params := ParseParams(paramsStr)
	//fmt.Println(params)
	//fmt.Println(s)
	if _, ok := funcsMap[funcName]; ok {
		return nil
	}
	funcsMap[funcName] = ""
	MakeCFunc(f, funcName, returnType, params)

	return nil
}

type Param struct {
	Name  string
	Type  string
	IsVar bool
}

func ParseParams(s string) []Param {
	if s == "" {
		return nil
	}
	ps := make([]Param, 0)
	pss := strings.Split(s, ";")
	for _, p := range pss {
		subps := strings.Split(strings.TrimSpace(p), ":")
		if len(subps) >= 2 {

			name := strings.TrimSpace(subps[0])
			typeStr := strings.TrimSpace(subps[1])
			isVar := strings.HasPrefix(name, "var ") || strings.HasPrefix(name, "out ")
			if isVar {
				name = strings.TrimPrefix(name, "var ")
				name = strings.TrimPrefix(name, "out ")

			}
			name = strings.TrimPrefix(name, "const ")
			name = strings.TrimSpace(name)
			// 共用类型的参数
			if strings.Index(name, ",") != -1 {
				ssubps := strings.Split(name, ",")
				for _, sps := range ssubps {
					var item Param
					item.IsVar = isVar
					item.Name = strings.TrimSpace(sps)
					item.Type = typeStr
					ps = append(ps, item)
				}

			} else {
				var item Param
				item.Name = name
				item.Type = typeStr
				item.IsVar = isVar
				ps = append(ps, item)
			}
		}
	}
	return ps
}

func MakeCFunc(f *CFile, name, returnType string, params []Param) error {

	if name == "DSendMessage" || name == "DCreateURLShortCut" {
		f.WLn()
		f.WLn()
		if name == "DSendMessage" {
			f.W("#ifndef _WIN32\n")
		} else if name == "DCreateURLShortCut" {
			f.W("#ifdef _WIN32\n")
		}

		f.WLn()
	}

	//f.W("  ")
	f.W(fmt.Sprintf("void* p%s; \n", name))
	if returnType != "" {
		f.W(TypeConvert(returnType))
	} else {
		f.W("void")
	}
	f.W(" ")
	//f.W("LCLAPI")
	//f.W(" ")
	f.W(name)
	f.W("(")
	for i, ps := range params {
		if i > 0 {
			f.W(", ")
		}
		//if ps.Type == "PChar" && !ps.IsVar {
		//	f.W("const ")
		//}
		f.W(TypeConvert(ps.Type))
		if ps.IsVar {
			f.W("*")
		}
		f.W(" ")
		f.W(ps.Name)
	}

	getParamNames := func() string {
		if len(params) > 0 {
			ns := ""
			for i, s := range params {
				if i >= 0 {
					ns += ", "
				}
				//ns += fmt.Sprintf("COV_PARAM(%s)", s.Name)
				ns += s.Name
			}
			for i := 0; i < 12-len(params); i++ {
				ns += " ,0"
			}
			return ns
		}
		return " ,0, 0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0"
	}

	f.W(") {\n")
	f.W(fmt.Sprintf("    GET_FUNC_ADDR(%s)\n", name))
	f.W("    ")
	if returnType != "" {
		f.W(fmt.Sprintf("return (%s)", TypeConvert(returnType)))
	}
	f.W(fmt.Sprintf("MySyscall(p%s, %d%s);", name, len(params), getParamNames()))
	f.WLn()
	f.W("}\n")
	f.WLn()

	// 添加结束
	if name == "DWindowFromPoint" || name == "DCreateShortCut" {
		f.WLn()
		f.W("#endif\n")
		f.WLn()
	}

	return nil
}

func TypeConvert(src string) string {
	if val, ok := typeMap[src]; ok {
		return val
	}
	return src
}

func parseEnums(fileName string) []byte {
	buff := bytes.NewBuffer(nil)
	buff.WriteString("//" + fileName + "\n")
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	lines := bytes.Split(bs, []byte("\n"))
	i := 0
	for i < len(lines) {
		s := string(bytes.TrimSpace(lines[i]))
		if strings.HasPrefix(s, "type") {
			//fmt.Println(s)
			sp := strings.Split(s, " ")
			if len(sp) >= 3 {
				switch strings.TrimSpace(sp[2]) { // 根据第三个成员判断是哪种

				case "=":
					// 集合类型
					switch sp[len(sp)-1] {
					case "TSet":
						buff.WriteString("typedef ")
						buff.WriteString("TSet ")
						buff.WriteString(strings.TrimSpace(sp[1]))
						buff.WriteString(";\n")
					default:
						buff.WriteString("typedef ")
						buff.WriteString(TypeConvert(strings.TrimSpace(sp[len(sp)-1])) + " ")
						buff.WriteString(strings.TrimSpace(sp[1]))
						buff.WriteString(";\n")
					}
				//case "TBorderStyle":
				//	buff.WriteString("typedef ")
				//	buff.WriteString(TypeConvert(strings.TrimSpace(sp[len(sp)-1])) + " ")
				//	buff.WriteString(strings.TrimSpace(sp[1]))
				//	buff.WriteString(";\n")

				default:

					if strings.TrimSpace(sp[1]) == "TCursor" ||
						strings.TrimSpace(sp[1]) == "TLeftRight" ||
						strings.TrimSpace(sp[1]) == "TFormBorderStyle" ||
						strings.TrimSpace(sp[1]) == "TColorBoxStyle" ||
						strings.TrimSpace(sp[1]) == "TLinkAlignment" ||
						strings.TrimSpace(sp[1]) == "TMenuAutoFlag" ||
						strings.TrimSpace(sp[1]) == "TNumGlyphs" ||
						strings.TrimSpace(sp[1]) == "TShortCut" ||
						strings.TrimSpace(sp[1]) == "TScrollBarInc" {
						i++
						continue
					}

					// 枚举类型
					buff.WriteString("typedef enum ")

					buff.WriteString(" {\n")

					i++
					find := false
					for i < len(lines) {
						s = string(bytes.TrimSpace(lines[i]))
						if find && s == ")" {
							break
						}
						if strings.HasPrefix(s, "type") {
							i--
							break
						} else if strings.HasPrefix(s, "const (") {
							i++
							find = true
							continue
						}
						if find {

							firstLowerChar := func(sx string) string {
								sx = strings.TrimSpace(sx)
								sx = strings.ToLower(string(sx[0])) + string(sx[1:])
								return sx
							}

							if s != "" && !strings.HasPrefix(s, "//") && !strings.HasPrefix(s, "{") {

								cc := strings.Split(s, "= iota")
								buff.WriteString("    ")
								if len(cc) >= 2 {
									buff.WriteString(firstLowerChar(cc[0]))
								} else if len(cc) == 1 {
									if strings.Index(s, "CrSizeNESW") != -1 {
										s = s[:strings.Index(s, "//")]
									}

									if strings.Index(s, "//") != -1 {
										//if strings.Index(s, ")") != -1 {
										//	n := strings.Index(s, " ")
										//	buff.WriteString(firstLowerChar(s[:n]))
										//	buff.WriteString(s[n:] + ",")
										//} else {
										n := strings.Index(s, " ")
										buff.WriteString(firstLowerChar(s[:n]) + ",")
										buff.WriteString(s[n:])
										//}

									} else {
										buff.WriteString(firstLowerChar(s))
									}

								}
								if len(cc) == 1 {
									if strings.Index(s, "//") == -1 {
										buff.WriteString(",")
									}

								} else {
									buff.WriteString(",")
								}
								buff.WriteString("\n")
							}

						}
						i++
					}

					buff.WriteString("}")
					buff.WriteString(strings.TrimSpace(sp[1]))
					buff.WriteString(";\n\n")
				}
			}
		}
		i++
	}

	return buff.Bytes()
}
