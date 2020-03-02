package main

import (
	"fmt"
	"time"

	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {

	iniFile := vcl.NewIniFile(".\\test.ini")
	defer iniFile.Free()

	iniFile.WriteBool("First", "Bool", true)
	iniFile.WriteString("First", "String", "这是字符串")
	iniFile.WriteDateTime("First", "Time", time.Now())
	iniFile.WriteInteger("First", "Integer", 123456)
	iniFile.WriteFloat("First", "Float", 1.2555)

	fmt.Println("Bool:", iniFile.ReadBool("First", "Bool", false))
	fmt.Println("String:", iniFile.ReadString("First", "String", ""))
	fmt.Println("Time:", iniFile.ReadDate("First", "Time", time.Now()))
	fmt.Println("Integer:", iniFile.ReadInteger("First", "Integer", 0))
	fmt.Println("Float:", iniFile.ReadFloat("First", "Float", 0.0))
}
