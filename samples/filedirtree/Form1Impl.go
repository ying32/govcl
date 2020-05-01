//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	paths map[uintptr]string
}

func (f *TForm1) OnFormCreate(object vcl.IObject) {

	f.TreeView1.Items().Clear()
	f.TreeView1.Items().BeginUpdate()
	node := f.TreeView1.Items().Add(nil, "Root")
	f.paths = make(map[uintptr]string, 0)
	f.paths[node.Instance()] = "."
	f.walkFile(node, ".", false)
	node.Expand(true)
	f.TreeView1.Items().EndUpdate()

}

func (f *TForm1) OnTreeView1Click(object vcl.IObject) {
	node := f.TreeView1.Selected()
	if node != nil {
		if path, ok := f.paths[node.Instance()]; ok {
			_, err := os.Stat(path)
			if err == nil {
				if !os.IsNotExist(err) {
					f.ListBox1.Clear()
					f.ListBox1.Items().BeginUpdate()
					f.walkFile(nil, path, true)
					f.ListBox1.Items().EndUpdate()
				}
			}
			fmt.Println(path)
		} else {
			fmt.Println("没有", node.Instance())
		}
	}
}

func (f *TForm1) walkFile(node *vcl.TTreeNode, path string, isFile bool) {
	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	defer fd.Close()
	for {
		files, err := fd.Readdir(100)
		for _, file := range files {
			if !isFile {
				if file.IsDir() {
					curPath := path + string(os.PathSeparator) + file.Name()
					subNode := f.TreeView1.Items().AddChild(node, file.Name())
					//subNode.SetData(unsafe.Pointer(uintptr(index)))
					f.paths[subNode.Instance()] = curPath
					f.walkFile(subNode, curPath, isFile)
				}
			} else {
				f.ListBox1.Items().Add(file.Name())
			}
		}
		if err == io.EOF {
			break
		}
		if len(files) == 0 {
			break
		}
	}
}
