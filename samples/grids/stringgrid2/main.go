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

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

// 例子改自Lazarus例子中的StringGrid Editor

type InputData struct {
	Auto         string
	Text         string
	EditMask     string
	Button       types.TColor
	ButtonColumn string
	CheckBox     types.TCheckBoxState
	Ellipsis     string
	None         string
	PickList     string
}

type TMainForm struct {
	*vcl.TForm
	grid     *vcl.TStringGrid
	edit     *vcl.TEdit
	pnl      *vcl.TPanel
	dlgColor *vcl.TColorDialog
	iData    []InputData
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetWidth(1000)
	f.SetHeight(570)
	f.ScreenCenter()

	f.dlgColor = vcl.NewColorDialog(f)

	f.pnl = vcl.NewPanel(f)
	f.pnl.SetParent(f)
	f.pnl.SetAlign(types.AlBottom)

	// 创建一个文本编辑框，不可视
	f.edit = vcl.NewEdit(f)
	f.edit.SetParent(f.pnl)
	f.edit.SetHint("改变“None”文本")
	f.edit.SetShowHint(true)
	f.edit.SetLeft(100)
	f.edit.SetTop(10)
	f.edit.SetWidth(300)
	f.edit.SetOnChange(f.onEditChange)

	// 这里用手动创建，主要是演示使用哪些属性来设置这些效果
	f.grid = vcl.NewStringGrid(f)
	f.grid.SetParent(f)
	f.grid.SetAlign(types.AlClient)
	// 事件设置
	f.grid.SetOnButtonClick(f.onGridButtonClick)
	f.grid.SetOnCheckboxToggled(f.onGridCheckboxToggled)
	f.grid.SetOnDrawCell(f.onGridDrawCell)
	f.grid.SetOnGetCheckboxState(f.onGridGetCheckboxState)
	f.grid.SetOnGetCellHint(f.onGridGetCellHint)
	f.grid.SetOnGetEditMask(f.onGridGetEditMask)
	f.grid.SetOnSetEditText(f.onGridSetEditText)
	f.grid.SetOnValidateEntry(f.onGridValidateEntry)
	f.grid.SetOnSelectEditor(f.onGridSelectEditor)

	// 这里设置不可操作的列和行数
	f.grid.SetFixedCols(0)
	//f.grid.SetFixedRows(0)
	// 设置一些选项
	f.grid.SetOptions(f.grid.Options().Include(types.GoAlwaysShowEditor, types.GoCellHints, types.GoEditing, types.GoTabs))

	// 设置不可操作列的背景颜色
	f.grid.SetFixedColor(colors.ClGreen)

	//f.grid.SetFixedHotColor(colors.ClBlue)
	// 表格边框样式，这里设置为没有边框
	f.grid.SetBorderStyle(types.BsNone)

	// 设置表格为平面样式
	f.grid.SetFlat(true)
	// 设置flat后可以用这个修改fixed区域的表格线
	f.grid.SetFixedGridLineColor(colors.ClRed)

	// 表格列宽，自动大小后填充区域
	f.grid.SetAutoFillColumns(true)

	// 表头的风格
	f.grid.SetTitleStyle(types.TsNative)

	// 单击表头排序
	f.grid.SetColumnClickSorts(true)
	// 自定义排序箭头的图标
	//f.grid.SetImageIndexSortAsc()
	//f.grid.SetImageIndexSortDesc()
	// 添加col
	col := f.grid.Columns().Add()
	title := col.Title()
	title.SetCaption("Auto")
	title.SetColor(colors.ClRed)
	//title.SetAlignment(types.TaCenter)
	//title.SetImageIndex()
	//title.SetImageLayout()
	//title.SetLayout()
	//title.SetMultiLine()

	col = f.grid.Columns().Add()
	col.Title().SetCaption("EditMask")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsButton)
	col.Title().SetCaption("Button")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsButtonColumn)
	col.Title().SetCaption("ButtonColumn")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsCheckboxColumn)
	col.Title().SetCaption("CheckBox")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsEllipsis)
	col.Title().SetCaption("Ellipsis")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsNone)
	col.Title().SetCaption("None")

	col = f.grid.Columns().Add()
	col.SetButtonStyle(types.CbsPickList)
	col.Title().SetCaption("PickList")
	col.PickList().Add("Cow")
	col.PickList().Add("Dog")
	col.PickList().Add("Pig")
	col.PickList().Add("Goat")
	col.PickList().Add("Elephant")

	// 设置表格的行数，上面已经通过Columns设置了列数了，所以这里只设置行数
	f.grid.SetRowCount(30)

	// 要填充的数据
	f.iData = make([]InputData, f.grid.RowCount()-1)
	for i := 0; i < len(f.iData); i++ {
		f.iData[i].Button = colors.ClWhite      //颜色
		f.iData[i].CheckBox = types.CbUnchecked //CheckBox的状态
		f.iData[i].None = "Not editable"        //'None' can only be changed in program
		f.grid.SetCells(6, int32(i)+1, f.iData[i].None)
	}
	f.edit.SetText(f.iData[0].None)
	f.grid.SetOptions(f.grid.Options().Include(types.GoCellHints))
	f.grid.SetShowHint(true)
	f.grid.Columns().Items(7).PickList().Add("Giraffe") //Add an item progamatically
	//The others are added in the Object Inspector
	vcl.Application.SetHintPause(1)

	// 表格可视行列
	fmt.Println("VisibleColCount: ", f.grid.VisibleColCount(), ", VisibleRowCount:", f.grid.VisibleRowCount())

	// 事件

}

func (f *TMainForm) onEditChange(sender vcl.IObject) {
	for i := int32(1); i < f.grid.RowCount(); i++ { //'None' can only be changed in program
		f.iData[i-1].None = f.edit.Text()
		f.grid.SetCells(6, i, f.iData[i-1].None)
	}
}

func (f *TMainForm) onGridButtonClick(sender vcl.IObject, aCol, aRow int32) {
	//For these columns there is no manual entry into the cell,
	//so use ButtonClick event
	fmt.Println("GridButtonClick")
	switch aCol {
	case 2:
		if f.dlgColor.Execute() {
			f.iData[aRow-1].Button = f.dlgColor.Color() //store cell colour in array
			f.grid.Invalidate()                         //Could also use 'Repaint' te force DrawCell event
		}
	case 3:
		f.grid.SetOptions(f.grid.Options().Exclude(types.GoEditing))
		//Prevent write to previous cell
		f.iData[aRow-1].ButtonColumn = fmt.Sprintf("%d, %d", aCol, aRow)
		//store as string
		f.grid.SetCells(aCol, aRow, f.iData[aRow-1].ButtonColumn)
		f.grid.SetOptions(f.grid.Options().Include(types.GoEditing)) //Turn cell editing back on
	case 5:
		if v := vcl.InputBox("输入", "值: ", ""); v != "" {
			// Click 'tick' sign on calculator to get result
			f.iData[aRow-1].Ellipsis = v
			//Store as string
			f.grid.SetCells(aCol, aRow, f.iData[aRow-1].Ellipsis)
		}
	}

}

func (f *TMainForm) onGridCheckboxToggled(sender vcl.IObject, aCol, aRow int32, aState types.TCheckBoxState) {
	fmt.Println("onGridCheckboxToggled")
	if (aRow > 0) && (aCol == 4) {
		f.iData[aRow-1].CheckBox = aState
	}
}

func (f *TMainForm) onGridDrawCell(sender vcl.IObject, aCol, aRow int32, aRect types.TRect, state types.TGridDrawState) {
	//Note in Col 2 'Button' the Repaint event takes place before the focus changes
	//to another cell.
	if aRow > 0 { //Use DrawCell to paint rectangle
		if aCol == 2 { //Get colour from array
			f.grid.Canvas().Brush().SetColor(f.iData[aRow-1].Button)
			f.grid.Canvas().FillRect(aRect) //Paint Cell
		}
	}
}

func (f *TMainForm) onGridGetCheckboxState(sender vcl.IObject, ACol, ARow int32, value *types.TCheckBoxState) {
	if ARow > 0 && (ACol == 4) {
		*value = f.iData[ARow-1].CheckBox
	}
}

func (f *TMainForm) onGridGetCellHint(sender vcl.IObject, ACol, ARow int32, hintText *string) {
	switch ACol {
	case 0:
		*hintText = "Button style cbsAuto sting grid column\r\n - enter any text."
	case 1:
		*hintText = "Button style cbsAuto, with basic Editmask for Time format.\r\nUses ValidateEntry as cell is exited to enforce  max of '23:59:59'"
	case 2:
		*hintText = "Button style cbsButton that shows colour dialog\r\nand changes cell colour."
	case 3:
		*hintText = "Button style cbsButtonColumn that shows cell position."
	case 4:
		*hintText = "Button style cbsCheckbox that toggles 'check' mark."
	case 5:
		*hintText = "Button style cbsEllipsis that opens calculator.\r\nClick 'tick' on calculator to send value to cell."
	case 6:
		*hintText = "Button style cbsNone that cannot be changed manually.\r\nChange Edit box contents to change displayed text."
	case 7:
		*hintText = "Button style cbsPicklist that offers a choice from\r\na list set in the Object Inspector."
	}

}

func (f *TMainForm) onGridGetEditMask(sender vcl.IObject, aCol, aRow int32, value *string) {
	//'!' = delete leading blanks. '0' = position must be a number.
	//'1' = keep formatting symbols. '_' =  trailing '0'.
	//Does not limit fields to 23:59:59.
	//Use ValidateEntry and Copy()to check and change each character as the cell is exited.
	if aRow > 0 && aCol == 1 {
		*value = "!00:00:00;1;_"
	}
}

func (f *TMainForm) onGridSetEditText(sender vcl.IObject, aCol, aRow int32, value string) {
	//Capture text from columns 0 and 1 to ayMyInput.Auto and .EditMask
	//SetEditText works for each keystroke
	if aRow > 0 {
		if aCol == 0 {
			f.iData[aRow-1].Auto = f.grid.Cells(aCol, aRow)
		} else if aCol == 1 {
			f.iData[aRow-1].EditMask = f.grid.Cells(aCol, aRow)
		}
		fmt.Println(value) //Show text as it is typed
	}
}

func (f *TMainForm) onGridValidateEntry(sender vcl.IObject, aCol, aRow int32, oldValue string, newValue *string) {
	//Constrain to '23:59:59'.
	//This only takes effect on leaving cell.
	//s := []byte(*newValue)
	//if aRow > 0 && aCol == 1 {
	//	if s[0] > '2' {
	//		s[0] = '2'
	//	} else if s[1] > '3' {
	//		s[1] = '3'
	//	} else if s[3] > '5' {
	//		s[3] = '5'
	//	} else if s[6] > '5' {
	//		s[6] = '3'
	//	}
	//	*newValue = string(s)
	//}
}

func (f *TMainForm) onGridSelectEditor(sender vcl.IObject, aCol, aRow int32, editor **vcl.TWinControl) {
	//if aCol == 1 && aRow > 0 && f.edit != nil {
	//	f.edit.SetBoundsRect(f.grid.CellRect(aCol, aRow))
	//	f.edit.SetText(f.grid.Cells(f.grid.Col(), f.grid.Row()) + "abc")
	//	fmt.Println("edit: ", f.edit.Instance())
	//	*editor = vcl.AsWinControl(f.edit)
	//}
}
