package main

import (
	"fmt"

	. "github.com/ying32/govcl/vcl"
	. "github.com/ying32/govcl/vcl/rtl"
	. "github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TPlayListItem struct {
	Caption  string // 歌曲名
	Singer   string // 歌手
	Length   int32  // 歌曲长度
	Lyric    string // 歌词文件
	FileName string // 歌曲文件名
}

type TPlayControl struct {
	*TDrawGrid
	datas          []TPlayListItem
	focusedColor   TColor
	playColor      TColor
	mouseMoveColor TColor
	mouseMoveIndex int32
	playingIndex   int32
	singerPicR     TRect
	singerPic      *TBitmap

	OnSelect func(sender IObject, item TPlayListItem)
}

func NewPlayControl(owner IComponent) *TPlayControl {
	m := new(TPlayControl)
	m.TDrawGrid = NewDrawGrid(owner)

	m.TDrawGrid.SetDefaultDrawing(false)
	m.TDrawGrid.SetDefaultRowHeight(24)
	if LcLLoaded() {
		m.TDrawGrid.SetOptions(Include(0, GoLzRangeSelect, GoLzRowSelect))
	} else {
		m.TDrawGrid.SetOptions(Include(0, GoRangeSelect, GoRowSelect))
	}
	m.TDrawGrid.SetRowCount(1)
	m.TDrawGrid.SetColCount(4)
	m.TDrawGrid.SetFixedRows(0)
	m.TDrawGrid.SetFixedCols(0)
	m.TDrawGrid.SetGridLineWidth(0)
	m.TDrawGrid.SetBorderStyle(BsNone)
	m.TDrawGrid.SetScrollBars(SsVertical)

	m.TDrawGrid.SetWidth(536)
	m.TDrawGrid.SetHeight(397)

	m.TDrawGrid.SetDrawingStyle(GdsThemed)
	// 加载时取消第一行永远被选中
	m.TDrawGrid.SetSelection(TGridRect{24, 24, 24, 24})

	m.TDrawGrid.SetColWidths(0, 60)
	m.TDrawGrid.SetColWidths(1, 230)
	m.TDrawGrid.SetColWidths(2, 100)
	m.TDrawGrid.SetColWidths(3, 80)

	//m.TDrawGrid.SetColor(0x39302c) //0x00EDEEF9)
	m.TDrawGrid.SetDoubleBuffered(true)

	m.focusedColor = colors.ClMoneyGreen //0x20302c //0x00C8CBEB
	m.playColor = m.focusedColor + 12
	m.mouseMoveColor = m.focusedColor - 12

	m.mouseMoveIndex = -1
	m.playingIndex = -1

	m.TDrawGrid.SetOnDblClick(m.onDblClick)
	m.TDrawGrid.SetOnMouseMove(m.onMouseMove)
	//m.TDrawGrid.SetOnMouseDown(m.onMouseDown)
	m.TDrawGrid.SetOnDrawCell(m.onDrawCell)
	//m.TDrawGrid.SetOnMouseEnter(m.onMouseEnter)
	m.TDrawGrid.SetOnMouseLeave(m.onMouseLeave)

	return m
}

func (p *TPlayControl) Add(item TPlayListItem) int32 {
	p.datas = append(p.datas, item)
	p.SetRowCount(int32(len(p.datas)))
	return int32(len(p.datas)) - 1
}

func (p *TPlayControl) onDrawCell(sender IObject, aCol, aRow int32, rect TRect, state TGridDrawState) {
	canvas := p.Canvas()
	//canvas.Brush().SetColor(0x39302c)
	//canvas.FillRect(p.ClientRect())

	if len(p.datas) > 0 {

		if aRow < int32(len(p.datas)) {
			drawFlags := Include(0, TfVerticalCenter, TfSingleLine, TfEndEllipsis)
			item := p.datas[int(aRow)]
			if p.mouseMoveIndex == aRow && p.playingIndex != aRow && !InSets(state, GdFocused) && !InSets(state, GdSelected) {
				canvas.Brush().SetColor(p.focusedColor - 12)
			} else if InSets(state, GdFocused) || InSets(state, GdSelected) {
				canvas.Brush().SetColor(p.focusedColor)
			} else {
				canvas.Brush().SetColor(p.Color())
			}

			if p.playingIndex == aRow {
				canvas.Brush().SetColor(p.focusedColor + 12)
				p.SetRowHeights(aRow, 50)
				if p.RowHeights(aRow) < 50 {
					p.ScrollBy(0, aRow+2)
				}
			} else {
				p.SetRowHeights(aRow, 24)
			}
			canvas.FillRect(rect)
			r := p.CellRect(aCol, aRow)
			switch aCol {
			case 0:

				if aRow == p.playingIndex {
					if !p.singerPicR.IsEmpty() {
						r.Left += 1
						r.Top += +1
						r.Bottom -= -1
						//canvas.StretchDraw(r, FSingerPic);
					}
				} else {
					r.Inflate(-10, 0)
					s := fmt.Sprintf("%d.", aRow+1)
					canvas.TextRect3(&r, s, Include(drawFlags, TfRight))
				}

			case 1:
				if aRow == p.playingIndex {
					r.Inflate(-10, 0)
					canvas.Font().SetSize(11)
					canvas.Font().SetStyle(Include(0, FsBold))
					canvas.TextRect3(&r, item.Caption, drawFlags)
				} else {
					r.Inflate(-5, 0)
					canvas.TextRect3(&r, item.Caption, drawFlags)
				}
				canvas.Font().SetSize(8)
				canvas.Font().SetStyle(0)
			case 2:
				r.Inflate(-5, 0)
				canvas.TextRect3(&r, item.Singer, drawFlags)
			case 3:
				r.Inflate(-5, 0)
				canvas.TextRect3(&r, p.mediaLengthToTimeStr(item.Length), Include(drawFlags, TfRight))
			}
		}

	} else {
		//p.Canvas().Brush().SetColor(p.Color())
		//p.Canvas().FillRect(rect)
	}
}

func (p *TPlayControl) onMouseMove(sender IObject, shift TShiftState, x, y int32) {
	if !p.Enabled() {
		return
	}
	var col, row int32
	p.MouseToCell(x, y, &col, &row)
	p.mouseMoveIndex = row
	if p.mouseMoveIndex == -1 {
		return
	}
	p.Invalidate()
}

func (p *TPlayControl) onMouseDown(sender IObject, button TMouseButton, shift TShiftState, x, y int32) {

}

func (p *TPlayControl) onDblClick(sender IObject) {
	if !p.Enabled() {
		return
	}
	row := p.Row()
	if row == -1 {
		return
	}
	p.playingIndex = row
	p.Invalidate()
	if p.OnSelect != nil {
		p.OnSelect(p, p.datas[row])
	}
}

func (p *TPlayControl) onMouseEnter(sender IObject) {

}

func (p *TPlayControl) onMouseLeave(sender IObject) {
	if !p.Enabled() {
		return
	}
	p.mouseMoveIndex = -1
	p.Invalidate()
}

func (p *TPlayControl) mediaLengthToTimeStr(alen int32) string {
	return fmt.Sprintf("%.2d:%.2d", int(float32(alen)/1000.0)/60, int(float32(alen)/1000.0)%60)
}
