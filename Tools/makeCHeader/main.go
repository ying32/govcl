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
		"IObject":              "TObject",
		"IWinControl":          "TWinControl",
		"IComponent":           "TComponent",
		"IControl":             "TControl",
		"bool":                 "BOOL",
		"uint16":               "uint16_t",

		//"TResItem":             "",
	}
	funcsMap   = make(map[string]string, 0)
	classArray = make([]string, 0)
	objsMap    = make(map[string]string)
)

func main() {

	file := NewCFile("./test/liblcl.h")
	file.WriteHeader()

	// UInt64 WINAPI MySyscall(void* AProc, intptr_t ALen, void* A1, void* A2, void* A3, void* A4, void* A5, void* A6, void* A7, void* A8, void* A9, void* A10, void* A11, void* A12);
	funcsMap["MySyscall"] = ""     // 排除此函数，手动构建
	funcsMap["DGetClassName"] = "" // 已经废弃，govcl中也没引用
	// 已经是c语言的了，则不需要这些了
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
	parseFile(file, "LazarusDef.inc", false)

	file.WLn()
	file.W("#ifdef __linux__\n")
	parseFile(file, "ulinuxpatchs.pas", true)
	file.WLn()
	file.W("#endif\n")
	file.WLn()

	file.WLn()
	file.W("#ifdef __APPLE__\n")
	parseFile(file, "umacospatchs.pas", true)
	file.WLn()
	file.W("#endif\n")
	file.WLn()

	parseFile(file, "uformdesignerfile.pas", true)

	// ulinuxpatchs.pas
	// umacospathcs.pas

	//fmt.Println(file.String())

	file.WriteFooter()
	// 生成枚举
	file.WLn()
	//file.WComment("枚举定义\n")

	file.AddReplaceFlag("typedefs", getClassDefs())
	file.AddReplaceFlag("enumdefs", parseEnums("../../vcl/types/enums.go"))
	file.AddReplaceFlag("eventdefs", parseEvents("../../vcl/events.go"))
	file.AddReplaceFlag("colorconsts", parseConst("../../vcl/types/colors/colors.go"))
	file.AddReplaceFlag("keyconsts", parseConst("../../vcl/types/keys/keys.go"))
	file.AddReplaceFlag("typeconsts", parseConst("../../vcl/types/consts.go"))

	file.Save()

}

func ReadFile(fileName string) ([]byte, error) {
	bs, err := ioutil.ReadFile("../../UILibSources/liblcl/" + fileName)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func parseFile(f *CFile, fileName string, isclass bool) {
	bs, err := ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	f.WComment(fileName)
	// {无效参数}
	bs = bytes.Replace(bs, []byte("{无效参数}"), nil, -1)
	bs = bytes.Replace(bs, []byte("\r"), nil, -1)
	sps := bytes.Split(bs, []byte("\n"))
	for i, line := range sps {
		s := string(bytes.TrimSpace(line))
		if (strings.HasPrefix(strings.ToLower(s), "function") || strings.HasPrefix(strings.ToLower(s), "procedure")) && strings.HasSuffix(s, "extdecl;") {
			eventType := ""
			if i > 0 {
				prevLine := string(bytes.TrimSpace(sps[i-1]))
				if strings.HasPrefix(prevLine, "//EVENT_TYPE:") {
					//fmt.Println("事件：", prevLine)
					eventType = strings.TrimSpace(strings.TrimPrefix(prevLine, "//EVENT_TYPE:"))
				}
			}
			parseFunc(f, s, isclass, eventType)
		}
	}

}

func parseClassFiles(f *CFile, fileName string) {
	bs, err := ReadFile(fileName)
	if err != nil {
		panic(err)
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
			// 后面事件判断是否为类的
			objsMap[className] = className
			// 生成类型定义用的
			classArray = append(classArray, className)
			fmt.Println(incFileName)
			parseFile(f, incFileName, true)
		}
	}

}

func parseFunc(f *CFile, s string, isclass bool, eventType string) error {
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
	params := ParseParams(paramsStr, eventType)
	//fmt.Println(params)
	//fmt.Println(s)
	if _, ok := funcsMap[funcName]; ok {
		return nil
	}
	funcsMap[funcName] = ""
	MakeCFunc(f, funcName, returnType, params, isclass)

	return nil
}

type Param struct {
	Name  string
	Type  string
	IsVar bool
	IsArr bool // 事件那
}

func ParseParams(s string, eventType string) []Param {
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
				if eventType != "" && name == "AEventId" {
					item.Type = eventType
				} else {
					item.Type = typeStr
				}
				item.IsVar = isVar
				ps = append(ps, item)
			}
		}
	}
	return ps
}

func MakeCFunc(f *CFile, name, returnType string, params []Param, isclass bool) error {

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

	//f.W(fmt.Sprintf("static void* p%s; \n", name))
	f.W(fmt.Sprintf("DEFINE_FUNC_PTR(%s)\n", name))
	if returnType != "" {
		f.W(TypeConvert(returnType))
	} else {
		f.W("void")
	}
	f.W(" ")
	//f.W("LCLAPI")
	//f.W(" ")
	// 不是类的成员，不要那个D开头
	tempName := name
	if !isclass {
		if tempName[0] == 'D' {
			tempName = tempName[1:]
		}
	}
	f.W(tempName)

	f.W("(")
	for i, ps := range params {
		if i > 0 {
			f.W(", ")
		}
		if ps.Type == "PChar" && !ps.IsVar {
			//f.W("const ")
			f.W("CChar ")
		}
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

	buff, lines, err := ReadFileLines(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

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

func parseEvents(fileName string) []byte {
	buff, lines, err := ReadFileLines(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, line := range lines {
		s := string(bytes.TrimSpace(line))
		if strings.HasPrefix(s, "type ") {
			s = strings.TrimPrefix(s, "type ")
			// 干掉注释
			if i := strings.Index(s, "//"); i != -1 {
				s = s[:i]
			}
			if i := strings.Index(s, "/*"); i != -1 {
				s = s[:i]
			}
			s = strings.TrimSpace(s)
			// 检测是否有返回值，有则移到参数后面
			if i := strings.Index(s, ")"); i != -1 {
				retVal := strings.TrimSpace(s[i+1:])
				if len(retVal) > 3 {
					s = s[:i] + ", result " + retVal + ")" // 重新处理这个返回值
				}
			}

			// type TLVOwnerDataHintEvent = TLVDataHintEvent
			// 处理上面这种情况
			if strings.Index(s, "=") != -1 {
				fmt.Println("------------", s)
				sp := strings.Split(s, "=")
				buff.WriteString(fmt.Sprintf("typedef %s %s;\n\n", strings.TrimSpace(sp[1]), strings.TrimSpace(sp[0])))
				continue
			}

			idx := strings.Index(s, " ")
			name := s[:idx]
			body := strings.TrimRight(strings.TrimPrefix(s[idx+1:], "func("), ")")
			//fmt.Println(name, "\n", body)

			params := make([]Param, 0)
			for _, ps := range strings.Split(body, ",") {
				ps = strings.TrimSpace(ps)
				subps := strings.Split(ps, " ")
				item := Param{}
				if len(subps) >= 1 {
					item.Name = strings.TrimSpace(subps[0])
				}
				if len(subps) >= 2 {
					item.Type = strings.Trim(strings.TrimSpace(subps[1]), "*")
					item.IsVar = strings.HasPrefix(strings.TrimSpace(subps[1]), "*")
					item.IsArr = strings.HasPrefix(strings.TrimSpace(subps[1]), "[]")
					if item.IsArr {
						item.Type = "void*" // 所有数组都变成一个指针类型 //strings.TrimPrefix(item.Type, "[]")
					}
				}
				params = append(params, item)
				// 如果上个参数是一个数组，则添加一个数组长度参数
				if item.IsArr {
					item := Param{}
					item.Name = "len"
					item.Type = "intptr_t"
					params = append(params, item)
				}
			}
			// 处理参数，将所有参数的类型都补上
			lastType := ""
			lastIsVar := false
			for i := len(params) - 1; i >= 0; i-- {
				item := params[i]
				if item.Type == "" {
					item.IsVar = lastIsVar
					item.Type = lastType
					params[i] = item
				}
				lastType = item.Type
				lastIsVar = item.IsVar
			}

			getcppComment := func() string {
				s := "void ("
				if len(params) > 0 && params[0].Name != "" {
					for i, ps := range params {
						if i > 0 {
							s += ", "
						}
						if _, ok := objsMap[ps.Type]; ok {
							s += ps.Type
						} else {
							s += TypeConvert(ps.Type)
							if ps.IsVar {
								s += "*"
							}
						}
						s += " " + ps.Name
					}
				}
				return s + ")\n"
			}

			// 注释要修改
			buff.WriteString("// " + getcppComment())
			buff.WriteString(fmt.Sprintf("typedef void(*%s)(", name))
			// 写参数
			if len(params) > 0 && params[0].Name != "" {
				for i, ps := range params {
					if i > 0 {
						buff.WriteString(", ")
					}
					if _, ok := objsMap[ps.Type]; ok {
						buff.WriteString(ps.Type)
					} else {
						buff.WriteString(TypeConvert(ps.Type))
						if ps.IsVar {
							buff.WriteString("*")
						}
					}
				}
			} else {
				buff.WriteString("void")
			}
			buff.WriteString(");\n\n")
		}
	}
	//fmt.Println(buff.String())
	return buff.Bytes()
}

func getClassDefs() []byte {
	buf := bytes.NewBuffer(nil)
	for _, class := range classArray {
		buf.WriteString(fmt.Sprintf("typedef void* %s;\n", class))
	}
	return buf.Bytes()
}

func parseConst(filename string) []byte {

	buff, lines, err := ReadFileLines(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	i := 0
	for i < len(lines) {
		s := string(bytes.TrimSpace(lines[i]))
		if strings.HasPrefix(s, "const (") {
			i++
			for s != ")" {
				s = string(bytes.TrimSpace(lines[i]))
				// 注释也收集
				if strings.HasPrefix(s, "//") || strings.HasPrefix(s, "/*") {
					buff.WriteString("\n")
					buff.WriteString(s)
					buff.WriteString("\n\n")
				} else {
					ss := strings.Split(s, "=")
					if len(ss) == 2 {
						buff.WriteString("#define ")
						buff.WriteString(firstLowerChar(strings.TrimSpace(ss[0])))
						buff.WriteString(strings.Repeat(" ", strings.Count(ss[0], " ")))

						buff.WriteString(strings.TrimSpace(ss[1]))
						buff.WriteString("\n")
					}
				}
				i++
			}
			buff.WriteString("\n")
			continue
		}
		i++
	}
	return buff.Bytes()
}

func firstLowerChar(sx string) string {
	sx = strings.TrimSpace(sx)
	sx = strings.ToLower(string(sx[0])) + string(sx[1:])
	return sx
}

func ReadFileLines(filename string) (*bytes.Buffer, [][]byte, error) {
	buff := bytes.NewBuffer(nil)
	buff.WriteString("//" + filename)
	buff.WriteString("\n")
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	return buff, bytes.Split(bs, []byte("\n")), nil
}
