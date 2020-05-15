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
	"path/filepath"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	paths map[uintptr]string
}

func (f *TForm1) OnFormCreate(object vcl.IObject) {

	f.TreeView1.Items().Clear()
	f.TreeView1.Items().BeginUpdate()
	f.paths = make(map[uintptr]string, 0)
	f.walkFile(nil, ".", true)

	f.TreeView1.Items().EndUpdate()
}

func (f *TForm1) OnTreeView1Click(object vcl.IObject) {
	node := f.TreeView1.Selected()
	if node != nil {
		if path, ok := f.paths[node.Instance()]; ok {
			_, err := os.Stat(path)
			if err == nil {
				if !os.IsNotExist(err) {
					f.ListView1.Clear()
					f.ListView1.Items().BeginUpdate()
					f.walkFile(nil, path, false)
					f.ListView1.Items().EndUpdate()
				}
			}
			fmt.Println(path)
		} else {
			fmt.Println("没有", node.Instance())
		}
	}
}

func (f *TForm1) walkFile(node *vcl.TTreeNode, path string, recurse bool) {
	fd, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		return
	}
	defer fd.Close()

	// 添加当前目录
	if recurse && (node == nil) {
		subNode := f.TreeView1.Items().Add(nil, ".")
		f.paths[subNode.Instance()] = path
		subNode.SetSelectedIndex(1)
		subNode.SetStateIndex(0)
	}

	for {
		files, err := fd.Readdir(100)
		for _, file := range files {
			if recurse {
				if file.IsDir() {
					curPath := path + string(os.PathSeparator) + file.Name()

					var subNode *vcl.TTreeNode
					if node == nil {
						subNode = f.TreeView1.Items().Add(nil, file.Name())
					} else {
						subNode = f.TreeView1.Items().AddChild(node, file.Name())
					}

					//subNode.SetImageIndex(0)
					subNode.SetSelectedIndex(1)
					subNode.SetStateIndex(0)

					//subNode.SetData(unsafe.Pointer(uintptr(index)))
					f.paths[subNode.Instance()] = curPath
					f.walkFile(subNode, curPath, recurse)
					subNode.Expand(false)

				}
			} else {
				item := f.ListView1.Items().Add()
				item.SetCaption(file.Name())
				item.SubItems().Add(fmt.Sprintf("%d byte", file.Size()))
				item.SubItems().Add(filepath.Ext(file.Name()))
				item.SetImageIndex(2)
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
