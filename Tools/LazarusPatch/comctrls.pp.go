package main

var (
	ComCtrlsPPData = []PatchItem{
		{Find1: []byte("TTreeNodes = class(TPersistent)"), Find2: []byte("procedure ReadData(Stream: TStream);"), Content: []byte("    procedure ReadDelphiNodeData(Stream: TStream);")},
		{Find1: []byte("TListItems = class(TPersistent)"), Find2: []byte("procedure ReadData(Stream: TStream);"), Content: []byte("    procedure ReadItemData(Stream: TStream);")},
	}
)

// 查找 TTreeNodes = class(TPersistent)

// 查找 procedure ReadData(Stream: TStream);

// 添加  procedure ReadDelphiNodeData(Stream: TStream);

//  -------TListView

// 查找 TListItems = class(TPersistent)

// 查找 procedure ReadData(Stream: TStream);

// 添加 procedure ReadItemData(Stream: TStream);
