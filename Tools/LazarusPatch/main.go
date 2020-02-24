package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Lazarus的补丁工具

var (
	lazarusPath = ""
)

func main() {

	// lazarus的路径
	lazarusPath = GetLazarusPath()
	checkLazarusPath()

	// 开始进行文的补丁操作

	patchFile("lcl/comctrls.pp", ComCtrlsPPData)
	patchFile("lcl/include/listitems.inc", listItemsIncData)
	patchFile("lcl/include/treeview.inc", treeViewIncData)
}

func checkLazarusPath() {
	if lazarusPath == "" {
		panic("未找到Lazarus。")
	}
	if !strings.HasSuffix(lazarusPath, string(os.PathSeparator)) {
		lazarusPath += string(os.PathSeparator)
	}
	if !checkFileExists(lazarusPath) {
		panic("不存在路径：" + lazarusPath)
	}
	fmt.Println("Lazarus Root: ", lazarusPath)
}

func checkFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func copyFile(src, dest string) error {
	filedest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer filedest.Close()
	filesrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer filesrc.Close()
	_, err = io.Copy(filedest, filesrc)
	return err
}

type PatchItem struct {
	Find1   []byte
	Find2   []byte
	Content []byte
	IsMod   bool
	IsFind1 bool
}

var (
	ying32PatchFlag = []byte("ying32 patch")
)

func cmpFile(file1, file2 string) bool {
	f1, _ := ioutil.ReadFile(file1)
	f2, _ := ioutil.ReadFile(file2)
	return sha1.Sum(f1) == sha1.Sum(f2)
}

func restoreBackupFile(fileName string) bool {
	return false
}

func patchFile(fileName string, datas []PatchItem) {

	fn := lazarusPath + fileName
	fmt.Println("patch File: ", fn)

	//if restoreBackupFile(fn) {
	//
	//}
	if len(datas) == 0 {
		fmt.Println("find1和content不能为空。")
		return
	}

	if !checkFileExists(fn) {
		fmt.Println("文件：", fn, "不存在。")
		return
	}
	fmt.Println("备份文件检查")
	// 检查备份文件
	bakupFile := fn + ".ying32bak"
	if !checkFileExists(bakupFile) {
		fmt.Println("备文件不存在，创建。。。")
		if err := copyFile(fn, bakupFile); err == nil {
			fmt.Println("备份文件成功")
		} else {
			fmt.Println("备份文件失败：", err)
			return
		}
	} else {
		// 如果存在，检查md5值是否相等，不相等的重新备份
		bV := cmpFile(fn, bakupFile)
		fmt.Println("检查备份文件与源文件内容是否一致:", bV)
		if !bV {
			temp, _ := ioutil.ReadFile(fn)
			if !bytes.Contains(temp, ying32PatchFlag) {
				fmt.Println("重新备份文件")
				if err := copyFile(fn, bakupFile); err == nil {
					fmt.Println("重新备份文件成功")
				} else {
					fmt.Println("重新备份文件失败：", err)
					return
				}
			} else {
				fmt.Println("不需要重新备份。")
			}
		}
	}

	// 开始修改内容
	fileContent, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println(err)
		return
	}

	if bytes.Contains(fileContent, ying32PatchFlag) {
		fmt.Println("文件已经修补过。")
		return
	}

	lines := bytes.Split(fileContent, []byte{'\n'})
	newBuffer := bytes.NewBuffer([]byte("// ying32 patch " + time.Now().String() + "\n"))

	for _, line := range lines {
		for j := 0; j < len(datas); j++ {
			// 这个项目已经修改完成了
			item := datas[j]
			if item.IsMod {
				continue
			}

			if bytes.HasPrefix(bytes.TrimSpace(line), item.Find1) {
				if len(item.Find2) > 0 {
					item.IsFind1 = true
					datas[j] = item
				} else {
					newBuffer.Write(item.Content)
					newBuffer.WriteByte('\n')
					item.IsMod = true
					datas[j] = item
				}
			}
			if item.IsFind1 && len(item.Find2) > 0 {
				if bytes.HasPrefix(bytes.TrimSpace(line), item.Find2) {
					// 找到第二个项目了，添加进去
					newBuffer.Write(item.Content)
					newBuffer.WriteByte('\n')
					item.IsFind1 = false
					item.IsMod = true
					datas[j] = item
				}
			}
		}
		newBuffer.Write(line)
		newBuffer.WriteByte('\n')
	}

	if newBuffer.Len() == 0 {
		fmt.Println("修改后的文件不能为空。")
		return
	}

	//err = ioutil.WriteFile("F:\\Users\\zhuying\\Desktop\\test\\"+filepath.Base(fileName), newBuffer.Bytes(), os.ModePerm)
	err = ioutil.WriteFile(fn, newBuffer.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println("修改失败：", err)
		return
	}
	fmt.Println("修改成功。")
}
