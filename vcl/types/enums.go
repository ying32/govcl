//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

/*
  注意：Free Pascal中所有集合这里全部使用TSet(uint32)表示，也就是说最多32个元素

*/

// TAlign = (alNone, alTop, alBottom, alLeft, alRight, alClient, alCustom);
type TAlign int32

const (
	AlNone = iota + 0
	AlTop
	AlBottom
	AlLeft
	AlRight
	AlClient
	AlCustom
)

type TAlignSet = TSet

//  TFormBorderStyle = (bsNone, bsSingle, bsSizeable, bsDialog, bsToolWindow, bsSizeToolWin);
//  TBorderStyle = bsNone..bsSingle;
type TBorderStyle int32

const (
	BsNone = iota + 0
	BsSingle
	BsSizeable
	BsDialog
	BsToolWindow
	BsSizeToolWin
)

type TFormBorderStyle TBorderStyle

// vcl TFormStyle = (fsNormal, fsMDIChild, fsMDIForm, fsStayOnTop);
// lcl TFormStyle = (fsNormal, fsMDIChild, fsMDIForm, fsStayOnTop, fsSplash, fsSystemStayOnTop);
type TFormStyle int32

const (
	FsNormal = iota + 0
	FsMDIChild
	FsMDIForm
	FsStayOnTop

	// lcl
	fsSplash
	fsSystemStayOnTop
)

//  TPosition = (poDesigned, poDefault, poDefaultPosOnly, poDefaultSizeOnly, poScreenCenter, poDesktopCenter, poMainFormCenter, poOwnerFormCenter);
type TPosition int32

const (
	PoDesigned        = iota + 0 // use bounds from the designer (read from stream)
	PoDefault                    // LCL decision (normally window manager decides)
	PoDefaultPosOnly             // designed size and LCL position
	PoDefaultSizeOnly            // designed position and LCL size
	PoScreenCenter               // center form on screen (depends on DefaultMonitor)
	PoDesktopCenter              // center form on desktop (total of all screens)
	PoMainFormCenter             // center form on main form (depends on DefaultMonitor)
	PoOwnerFormCenter            // center form on owner form (depends on DefaultMonitor)
	PoWorkAreaCenter             // center form on working area (depends on DefaultMonitor)
)

//  TCursor = -32768..32767;
type TCursor int16

const (
	CrHigh = TCursor(0)

	CrDefault   = TCursor(0)
	CrNone      = TCursor(-1)
	CrArrow     = TCursor(-2)
	CrCross     = TCursor(-3)
	CrIBeam     = TCursor(-4)
	CrSize      = TCursor(-22)
	CrSizeNESW  = TCursor(-6) // diagonal north east - south west
	CrSizeNS    = TCursor(-7)
	CrSizeNWSE  = TCursor(-8)
	CrSizeWE    = TCursor(-9)
	CrSizeNW    = TCursor(-23)
	CrSizeN     = TCursor(-24)
	CrSizeNE    = TCursor(-25)
	CrSizeW     = TCursor(-26)
	CrSizeE     = TCursor(-27)
	CrSizeSW    = TCursor(-28)
	CrSizeS     = TCursor(-29)
	CrSizeSE    = TCursor(-30)
	CrUpArrow   = TCursor(-10)
	CrHourGlass = TCursor(-11)
	CrDrag      = TCursor(-12)
	CrNoDrop    = TCursor(-13)
	CrHSplit    = TCursor(-14)
	CrVSplit    = TCursor(-15)
	CrMultiDrag = TCursor(-16)
	CrSQLWait   = TCursor(-17)
	CrNo        = TCursor(-18)
	CrAppStart  = TCursor(-19)
	CrHelp      = TCursor(-20)
	CrHandPoint = TCursor(-21)
	CrSizeAll   = TCursor(-22)

	CrLow = -30
)

// TSeekOrigin = (soBeginning, soCurrent, soEnd);
type TSeekOrigin int32

const (
	SoBeginning = iota + 0
	SoCurrent
	SoEnd
)

// TPixelFormat = (pfDevice, pf1bit, pf4bit, pf8bit, pf15bit, pf16bit, pf24bit, pf32bit, pfCustom);
type TPixelFormat int32

const (
	PfDevice = iota + 0
	Pf1bit
	Pf4bit
	Pf8bit
	Pf15bit
	Pf16bit
	Pf24bit
	Pf32bit
	PfCustom
)

// TBalloonHintStyle = (bhsStandard, bhsBalloon);
//type TBalloonHintStyle int32
//
//const (
//	bhsStandard = iota + 0
//	bhsBalloon
//)

//TAlignment = (taLeftJustify, taRightJustify, taCenter);
type TAlignment int32

const (
	TaLeftJustify = iota + 0
	TaRightJustify
	TaCenter
)

//  TLeftRight = TAlignment.taLeftJustify..TAlignment.taRightJustify;
type TLeftRight int32

//  TBiDiMode = (bdLeftToRight, bdRightToLeft, bdRightToLeftNoAlign, bdRightToLeftReadingOnly);
type TBiDiMode int32

const (
	BdLeftToRight = iota + 0
	BdRightToLeft
	BdRightToLeftNoAlign
	BdRightToLeftReadingOnly
)

//  TVerticalAlignment = (taAlignTop, taAlignBottom, taVerticalCenter);
type TVerticalAlignment int32

const (
	TaAlignTop = iota + 0
	TaAlignBottom
	TaVerticalCenter
)

// TButtonStyle = (bsPushButton, bsCommandLink, bsSplitButton);
type TButtonStyle int32

const (
	BsPushButton = iota + 0
	BsCommandLink
	BsSplitButton
)

type TColorBoxStyle int32

type TComboBoxStyle int32

const (
	CsDropDown                  = iota + 0 // like an TEdit plus a button to drop down the list, default
	CsSimple                               // like an TEdit plus a TListBox
	CsDropDownList                         // like TLabel plus a button to drop down the list
	CsOwnerDrawFixed                       // like csDropDownList, but custom drawn
	CsOwnerDrawVariable                    // like csDropDownList, but custom drawn and with each item can have another height
	CsOwnerDrawEditableFixed               // like csOwnerDrawFixed, but with TEdit
	CsOwnerDrawEditableVariable            // like csOwnerDrawVariable, but with TEdit
)

//  TWindowState = (wsNormal, wsMinimized, wsMaximized);
type TWindowState int32

const (
	WsNormal = iota + 0
	WsMinimized
	WsMaximized

	// LCL
	WsFullScreen
)

//  TTextLayout = (tlTop, tlCenter, tlBottom);
type TTextLayout int32

const (
	TlTop = iota + 0
	TlCenter
	TlBottom
)

//  TEllipsisPosition = (epNone, epPathEllipsis, epEndEllipsis, epWordEllipsis);
type TEllipsisPosition int32

const (
	EpNone = iota + 0
	EpPathEllipsis
	EpEndEllipsis
	EpWordEllipsis
)

type TLinkAlignment TAlignment

// TListBoxStyle = (lbStandard, lbOwnerDrawFixed, lbOwnerDrawVariable, lbVirtual);
type TListBoxStyle int32

const (
	LbStandard = iota + 0
	LbOwnerDrawFixed
	LbOwnerDrawVariable
	LbVirtual
	//LbVirtualOwnerDraw
)

//TMenuItemAutoFlag = (maAutomatic, maManual, maParent);
type TMenuItemAutoFlag int32

const (
	MaAutomatic = iota + 0
	MaManual
	MaParent
)

//  TMenuAutoFlag = maAutomatic..maManual;
type TMenuAutoFlag TMenuItemAutoFlag

//TPopupAlignment = (paLeft, paRight, paCenter);
type TPopupAlignment int32

const (
	PaLeft = iota + 0
	PaRight
	PaCenter
)

//  TTrackButton = (tbRightButton, tbLeftButton);
type TTrackButton int32

const (
	TbRightButton = iota + 0
	TbLeftButton
)

// TProgressBarOrientation = (pbHorizontal, pbVertical, pbRightToLeft, pbTopDown);
type TProgressBarOrientation int32

const (
	PbHorizontal = iota + 0
	PbVertical
	// lcl
	PbRightToLeft
	PbTopDown
)

//  TProgressBarStyle = (pbstNormal, pbstMarquee);
type TProgressBarStyle int32

const (
	PbstNormal = iota + 0
	PbstMarquee
)

//  TProgressBarState = (pbsNormal, pbsError, pbsPaused);
type TProgressBarState int32

const (
	PbsNormal = iota + 0
	PbsError
	PbsPaused
)

//TButtonLayout = (blGlyphLeft, blGlyphRight, blGlyphTop, blGlyphBottom);
type TButtonLayout int32

const (
	BlGlyphLeft = iota + 0
	BlGlyphRight
	BlGlyphTop
	BlGlyphBottom
)

//  TButtonState = (bsUp, bsDisabled, bsDown, bsExclusive);
type TButtonState int32

const (
	BsUp        = iota + 0 // button is up
	BsDisabled             // button disabled (grayed)
	BsDown                 // button is down
	BsExclusive            // button is the only down in his group
	// lcl
	BsHot // button is under mouse
)

// TButtonStyle = (bsAutoDetect, bsWin31, bsNew);
//type TButtonStyle int32

const (
	BsAutoDetect = iota + 0
	BsWin31
	BsNew
)

//  TNumGlyphs = 1..4;
type TNumGlyphs int32

// TStaticBorderStyle = (sbsNone, sbsSingle, sbsSunken);
type TStaticBorderStyle int32

const (
	SbsNone = iota + 0
	sbsSingle
	sbsSunken
)

// TFontStyle = (fsBold, fsItalic, fsUnderline, fsStrikeOut);
type TFontStyle int32

const (
	FsBold = iota + 0
	FsItalic
	FsUnderline
	FsStrikeOut
)

// TFontStyles = set of TFontStyle
type TFontStyles = TSet

// TScrollStyle = (ssNone, ssHorizontal, ssVertical, ssBoth);
type TScrollStyle int32

const (
	SsNone = iota + 0
	SsHorizontal
	SsVertical
	SsBoth
	// lcl
	SsAutoHorizontal
	SsAutoVertical
	SsAutoBoth
)

//TSortType = (stNone, stData, stText, stBoth);
type TSortType int32

const (
	StNone = iota + 0
	StData
	StText
	StBoth
)

//  TMultiSelectStyles = (msControlSelect, msShiftSelect, msVisibleOnly, msSiblingOnly);
type TMultiSelectStyles int32

const (
	MsControlSelect = iota + 0
	MsShiftSelect
	MsVisibleOnly
	MsSiblingOnly
)

//  TListArrangement = (arAlignBottom, arAlignLeft, arAlignRight, arAlignTop, arDefault, arSnapToGrid);
type TListArrangement int32

const (
	ArAlignBottom = iota + 0
	ArAlignLeft
	ArAlignRight
	ArAlignTop
	ArDefault
	ArSnapToGrid
)

//  TViewStyle = (vsIcon, vsSmallIcon, vsList, vsReport);
type TViewStyle int32

const (
	VsIcon = iota + 0
	VsSmallIcon
	VsList
	VsReport
)

//  TItemState = (isNone, isCut, isDropHilited, isFocused, isSelected, isActivating);
type TItemState int32

const (
	IsNone = iota + 0
	IsCut
	IsDropHilited
	IsFocused
	IsSelected
	IsActivating
)

//  TItemStates = set of TItemState;
type TItemStates = TSet

//  TItemChange = (ctText, ctImage, ctState);
type TItemChange int32

const (
	CtText = iota + 0
	CtImage
	CtState
)

//  TItemFind = (ifData, ifPartialString, ifExactString, ifNearest);
type TItemFind int32

const (
	IfData = iota + 0
	IfPartialString
	IfExactString
	IfNearest
)

//  TSearchDirection = (sdLeft, sdRight, sdAbove, sdBelow, sdAll);
type TSearchDirection int32

const (
	SdLeft = iota + 0
	SdRight
	SdAbove
	SdBelow
	SdAll
)

//  TListHotTrackStyle = (htHandPoint, htUnderlineCold, htUnderlineHot);
type TListHotTrackStyle int32

const (
	HtHandPoint = iota + 0
	HtUnderlineCold
	HtUnderlineHot
)

//  TListHotTrackStyles = set of TListHotTrackStyle;
type TListHotTrackStyles = TSet

//  TItemRequests = (irText, irImage, irParam, irState, irIndent);
type TItemRequests int32

const (
	IrText = iota + 0
	IrImage
	IrParam
	IrState
	IrIndent
)

//TBrushStyle = (bsSolid, bsClear, bsHorizontal, bsVertical,
//    bsFDiagonal, bsBDiagonal, bsCross, bsDiagCross);
type TBrushStyle int32

const (
	BsSolid = iota + 0
	BsClear
	BsHorizontal
	BsVertical
	BsFDiagonal
	BsBDiagonal
	BsCross
	BsDiagCross

	// lcl
	BsImage
	BsPattern
)

type TPenStyle int32

const (
	PsSolid = iota + 0
	PsDash
	PsDot
	PsDashDot
	PsDashDotDot
	PsinsideFrame
	PsPattern
	PsClear
)

// TUDBtnType = (btNext, btPrev);
type TUDBtnType int32

const (
	btNext = iota + 0
	btPrev
)

//  TTabPosition = (tpTop, tpBottom, tpLeft, tpRight);
type TTabPosition int32

const (
	TpTop = iota + 0
	TpBottom
	TpLeft
	TpRight
)

//  TTabStyle = (tsTabs, tsButtons, tsFlatButtons);
type TTabStyle int32

const (
	TsTabs = iota + 0
	TsButtons
	TsFlatButtons
)

// TFontPitch = (fpDefault, fpVariable, fpFixed);
type TFontPitch int32

const (
	FpDefault = iota + 0
	FpVariable
	FpFixed
)

// TPenMode = (pmBlack, pmWhite, pmNop, pmNot, pmCopy, pmNotCopy,
//    pmMergePenNot, pmMaskPenNot, pmMergeNotPen, pmMaskNotPen, pmMerge,
//    pmNotMerge, pmMask, pmNotMask, pmXor, pmNotXor);
type TPenMode int32

const (
	PmBlack = iota + 0
	PmWhite
	PmNop
	PmNot
	PmCopy
	PmNotCopy
	PmMergePenNot
	PmMaskPenNot
	PmMergeNotPen
	PmMaskNotPen
	PmMerge
	PmNotMerge
	PmMask
	PmNotMask
	PmXor
	PmNotXor
)

// TTrackBarOrientation = (trHorizontal, trVertical);
type TTrackBarOrientation int32

const (
	TrHorizontal = iota + 0
	TrVertical
)

// TUDOrientation = (udHorizontal, udVertical);
type TUDOrientation int32

const (
	UdHorizontal = iota + 0
	UdVertical
)

//  TFontQuality = (fqDefault, fqDraft, fqProof, fqNonAntialiased, fqAntialiased,
//    fqClearType, fqClearTypeNatural);
type TFontQuality int32

const (
	FqDefault = iota + 0
	FqDraft
	FqProof
	FqNonAntialiased
	FqAntialiased
	FqClearType
	FqClearTypeNatural
)

// TCloseAction = (caNone, caHide, caFree, caMinimize);
type TCloseAction int32

const (
	CaNone = iota + 0
	CaHide
	CaFree
	CaMinimize
)

type TBalloonFlags int32

const (
	BfNone = iota + 0
	BfInfo
	BfWarning
	BfError
)

//  TMsgDlgType = (mtWarning, mtError, mtInformation, mtConfirmation, mtCustom);
type TMsgDlgType int32

const (
	MtWarning = iota + 0
	MtError
	MtInformation
	MtConfirmation
	MtCustom
)

//  TMsgDlgBtn = (mbYes, mbNo, mbOK, mbCancel, mbAbort, mbRetry, mbIgnore,
//    mbAll, mbNoToAll, mbYesToAll, mbHelp, mbClose);
type TMsgDlgBtn int32

const (
	MbYes = iota + 0
	MbNo
	MbOK
	MbCancel
	MbAbort
	MbRetry
	MbIgnore
	MbAll
	MbNoToAll
	MbYesToAll
	MbHelp
	MbClose
)

//  TMsgDlgButtons = set of TMsgDlgBtn;
type TMsgDlgButtons = TSet

// TSysLinkType = (sltURL, sltID);
type TSysLinkType int32

const (
	SltURL = iota + 0
	SltID
)

//  TStatusPanelStyle = (psText, psOwnerDraw);
type TStatusPanelStyle int32

const (
	PsText = iota + 0
	PsOwnerDraw
)

//  TStatusPanelBevel = (pbNone, pbLowered, pbRaised);
type TStatusPanelBevel int32

const (
	PbNone = iota + 0
	PbLowered
	PbRaised
)

//  TJPEGQualityRange = 1..100;   // 100 = best quality, 25 = pretty awful
type TJPEGQualityRange = uint32

//  TJPEGPerformance = (jpBestQuality, jpBestSpeed);
type TJPEGPerformance int32

const (
	JpBestQuality = iota + 0
	JpBestSpeed
)

//  TJPEGScale = (jsFullSize, jsHalf, jsQuarter, jsEighth);
type TJPEGScale int32

const (
	JsFullSize = iota + 0
	JsHalf
	JsQuarter
	JsEighth
)

//  TJPEGPixelFormat = (jf24Bit, jf8Bit);
type TJPEGPixelFormat = TPixelFormat

//const (
//	Jf24Bit = iota + 0
//	Jf8Bit
//)

//  TGIFVersion = (gvUnknown, gv87a, gv89a);
type TGIFVersion int32

const (
	GvUnknown = iota + 0
	Gv87a
	Gv89a
)

// Animation loop behaviour
//TGIFAnimationLoop = (
//  glDisabled,                 // Never loop
//  glEnabled,                  // Loop is specified in GIF
//  glContinously               // Loop continously regardless of GIF
//  );
type TGIFAnimationLoop int32

const (
	GlDisabled = iota + 0
	GlEnabled
	GlContinously
)

// Auto dithering of GIF output to Netscape 216 color palette
//TGIFDithering = (
//  gdDisabled,                 // Never dither
//  gdEnabled,                  // Always dither
//  gdAuto                      // Dither if Desktop DC supports <= 256 colors.
//  );
type TGIFDithering int32

const (
	GdDisabled = iota + 0
	GdEnabled
	GdAuto
)

// TCompressionLevel = 0..9;
type TCompressionLevel = uint32

type TShortCut uint16

type TNodeState int32

const (
	NsCut           = iota + 0 // = Node.Cut
	NsDropHilite               // = Node.DropTarget
	NsFocused                  // = Node.Focused
	NsSelected                 // = Node.Selected
	NsMultiSelected            // = Node.MultiSelected
	NsExpanded                 // = Node.Expanded
	NsHasChildren              // = Node.HasChildren
	NsDeleting                 // = Node.Deleting, set on Destroy
	NsVisible                  // = Node.Visible
	NsBound                    // bound to a tree, e.g. has Parent or is top lvl node
)

type TNodeAttachMode int32

const (
	NaAdd           = iota + 0 // add as last sibling of Destination
	NaAddFirst                 // add as first sibling of Destination
	NaAddChild                 // add as last child of Destination
	NaAddChildFirst            // add as first child of Destination
	NaInsert                   // insert in front of Destination
	NaInsertBehind             // insert behind Destination
)

//  TAddMode = (taAddFirst, taAdd, taInsert);
type TAddMode int32

const (
	TaAddFirst = iota + 0
	TaAdd
	TaInsert
)

//type TMultiSelectStyles int32
//
//const (
//	MsControlSelect = iota + 0
//	MsShiftSelect
//	MsVisibleOnly
//	MsSiblingOnly
//)

type TMultiSelectStyle = TSet

// TActionListState = (asNormal, asSuspended, asSuspendedEnabled);
type TActionListState int32

const (
	AsNormal = iota + 0
	AsSuspended
	AsSuspendedEnabled
)

// TGradientDirection = (gdHorizontal, gdVertical);
type TGradientDirection int32

const (
	GdHorizontal = iota + 0
	GdVertical
)

// TDrawingStyle = (dsFocus, dsSelected, dsNormal, dsTransparent);
type TDrawingStyle int32

const (
	DSFocus = iota + 0
	DSSelected
	DSNormal
	DSTransparent
)

// TImageType = (itImage, itMask);
type TImageType int32

const (
	ItImage = iota + 0
	ItMask
)

// TResType = (rtBitmap, rtCursor, rtIcon);
type TResType int32

const (
	RtBitmap = iota + 0
	RtCursor
	RtIcon
)

// TOverlay = 0..14;

// TLoadResource = (lrDefaultColor, lrDefaultSize, lrFromFile, lrMap3DColors, lrTransparent, lrMonoChrome);
type TLoadResource int32

const (
	LrDefaultColor = iota + 0
	LrDefaultSize
	LrFromFile
	LrMap3DColors
	LrTransparent
	LrMonoChrome
)

//   TLoadResources = set of TLoadResource;
type TLoadResources = TSet

// TColorDepth = (cdDefault, cdDeviceDependent, cd4Bit, cd8Bit, cd16Bit, cd24Bit, cd32Bit);
type TColorDepth int32

const (
	CdDefault = iota + 0
	CdDeviceDependent
	Cd4Bit
	Cd8Bit
	Cd16Bit
	Cd24Bit
	Cd32Bit
)

// TCheckBoxState = (cbUnchecked, cbChecked, cbGrayed);
type TCheckBoxState int32

const (
	CbUnchecked = iota + 0
	CbChecked
	CbGrayed
)

//  TListGroupState = (
//    lgsNormal,
//    lgsHidden,
//    lgsCollapsed,
//    lgsNoHeader,
//    lgsCollapsible,
//    lgsFocused,
//    lgsSelected,
//    lgsSubseted,
//    lgsSubSetLinkFocused
//  );
//type TListGroupState int32
//
//const (
//	LgsNormal = iota + 0
//	LgsHidden
//	LgsCollapsed
//	LgsNoHeader
//	LgsCollapsible
//	LgsFocused
//	LgsSelected
//	LgsSubseted
//	LgsSubSetLinkFocused
//)

//  TListGroupStateSet = set of TListGroupState;
//type TListGroupStateSet = TSet

// TTBDrawingStyle = (dsNormal, dsGradient);
type TTBDrawingStyle int32

const (
	DsNormal = iota + 0
	DsGradient
)

type TToolButtonStyle int32

const (
	TbsButton     = iota + 0 // button (can be clicked)
	TbsCheck                 // check item (click to toggle state, can be grouped)
	TbsDropDown              // button with dropdown button to show a popup menu
	TbsSeparator             // space holder
	TbsDivider               // space holder with line
	TbsButtonDrop            // button with arrow (not separated from each other)
)

//  TToolButtonState = (tbsChecked, tbsPressed, tbsEnabled, tbsHidden, tbsIndeterminate, tbsWrap, tbsEllipses, tbsMarked);
//type TToolButtonState int32
//
//const (
//	TbsChecked = iota + 0
//	TbsPressed
//	TbsEnabled
//	TbsHidden
//	TbsIndeterminate
//	TbsWrap
//	TbsEllipses
//	TbsMarked
//)

// TTBGradientDrawingOption = (gdoHotTrack, gdoGradient)
type TTBGradientDrawingOption int32

const (
	GdoHotTrack = iota + 0
	GdoGradient
)

// TTBGradientDrawingOptions = set of (gdoHotTrack, gdoGradient);
type TTBGradientDrawingOptions = TSet

// TColorDialogOption = (cdFullOpen, cdPreventFullOpen, cdShowHelp,
// cdSolidColor, cdAnyColor);
type TColorDialogOption int32

const (
	CdFullOpen = iota + 0
	CdPreventFullOpen
	CdShowHelp
	CdSolidColor
	CdAnyColor
)

type TColorDialogOptions = TSet

// TBorderIcon = (biSystemMenu, biMinimize, biMaximize, biHelp);
type TBorderIcon int32

const (
	BiSystemMenu = iota + 0
	BiMinimize
	BiMaximize
	BiHelp
)

// TBorderIcons = set of TBorderIcon;
type TBorderIcons = TSet

// TFontDialogOption = (fdAnsiOnly, fdTrueTypeOnly, fdEffects,
//     fdFixedPitchOnly, fdForceFontExist, fdNoFaceSel, fdNoOEMFonts,
//     fdNoSimulations, fdNoSizeSel, fdNoStyleSel,  fdNoVectorFonts,
//     fdShowHelp, fdWysiwyg, fdLimitSize, fdScalableOnly, fdApplyButton);
type TFontDialogOption int32

const (
	FdAnsiOnly = iota + 0
	FdTrueTypeOnly
	FdEffects
	FdFixedPitchOnly
	FdForceFontExist
	FdNoFaceSel
	FdNoOEMFonts
	FdNoSimulations
	FdNoSizeSel
	FdNoStyleSel
	FdNoVectorFonts
	FdShowHelp
	FdWysiwyg
	FdLimitSize
	FdScalableOnly
	FdApplyButton
)

//   TFontDialogOptions = set of TFontDialogOption;
type TFontDialogOptions = TSet

// { TOpenDialog }
// TOpenOption = (ofReadOnly, ofOverwritePrompt, ofHideReadOnly,
//   ofNoChangeDir, ofShowHelp, ofNoValidate, ofAllowMultiSelect,
//   ofExtensionDifferent, ofPathMustExist, ofFileMustExist, ofCreatePrompt,
//   ofShareAware, ofNoReadOnlyReturn, ofNoTestFileCreate, ofNoNetworkButton,
//   ofNoLongNames, ofOldStyleDialog, ofNoDereferenceLinks, ofEnableIncludeNotify,
//   ofEnableSizing, ofDontAddToRecent, ofForceShowHidden);
type TOpenOption int32

const (
	OfReadOnly        = iota + 0
	OfOverwritePrompt // if selected file exists shows a message, that file
	// will be overwritten
	OfHideReadOnly // hide read only file
	OfNoChangeDir  // do not change current directory
	OfShowHelp     // show a help button
	OfNoValidate
	OfAllowMultiSelect // allow multiselection
	OfExtensionDifferent
	OfPathMustExist // shows an error message if selected path does not exist
	OfFileMustExist // shows an error message if selected file does not exist
	OfCreatePrompt
	OfShareAware
	OfNoReadOnlyReturn // do not return filenames that are readonly
	OfNoTestFileCreate
	OfNoNetworkButton
	OfNoLongNames
	OfOldStyleDialog
	OfNoDereferenceLinks // do not resolve links while dialog is shown (only on Windows, see OFN_NODEREFERENCELINKS)
	OfNoResolveLinks     // do not resolve links after Execute
	OfEnableIncludeNotify
	OfEnableSizing    // dialog can be resized, e.g. via the mouse
	OfDontAddToRecent // do not add the path to the history list
	OfForceShowHidden // show hidden files
	OfViewDetail      // details are OS and interface dependent
	OfAutoPreview     // details are OS and interface dependent
)

// TOpenOptions = set of TOpenOption;
type TOpenOptions = TSet

// TOpenOptionEx = (ofExNoPlacesBar);
type TOpenOptionEx int32

const (
	OfExNoPlacesBar = iota + 0
)

// TOpenOptionsEx = set of TOpenOptionEx;
type TOpenOptionsEx = TSet

// { TPrintDialog }
// TPrintRange = (prAllPages, prSelection, prPageNums);
type TPrintRange int32

const (
	PrAllPages = iota + 0
	PrSelection
	PrPageNums
	// lcl
	PrCurrentPage
)

// TPrintDialogOption = (poPrintToFile, poPageNums, poSelection, poWarning,
//   poHelp, poDisablePrintToFile);
type TPrintDialogOption int32

const (
	PoPrintToFile = iota + 0
	PoPageNums
	PoSelection
	PoWarning
	PoHelp
	PoDisablePrintToFile
	//lcl
	PoBeforeBeginDoc
)

// TPrintDialogOptions = set of TPrintDialogOption;
type TPrintDialogOptions = TSet

// TPageSetupDialogOption = (psoDefaultMinMargins, psoDisableMargins,
//   psoDisableOrientation, psoDisablePagePainting, psoDisablePaper, psoDisablePrinter,
//   psoMargins, psoMinMargins, psoShowHelp, psoWarning, psoNoNetworkButton);
type TPageSetupDialogOption int32

const (
	PsoDefaultMinMargins = iota + 0
	PsoDisableMargins
	PsoDisableOrientation
	PsoDisablePagePainting
	PsoDisablePaper
	PsoDisablePrinter
	PsoMargins
	PsoMinMargins
	PsoShowHelp
	PsoWarning
	PsoNoNetworkButton
)

// TPageSetupDialogOptions = set of TPageSetupDialogOption;
type TPageSetupDialogOptions = TSet

// TPrinterKind = (pkDotMatrix, pkHPPCL);
type TPrinterKind int32

const (
	PkDotMatrix = iota + 0
	PkHPPCL
)

// TPageType = (ptEnvelope, ptPaper);
type TPageType int32

const (
	PtEnvelope = iota + 0
	PtPaper
)

// TPageMeasureUnits = (pmDefault, pmMillimeters, pmInches);
type TPageMeasureUnits int32

const (
	PmDefault = iota + 0
	PmMillimeters
	PmInches
)

//   TStringsOption = (soStrictDelimiter, soWriteBOM, soTrailingLineBreak, soUseLocale);
type TStringsOption = uint32

const (
	SoStrictDelimiter = iota + 0
	SoWriteBOM
	SoTrailingLineBreak
	SoUseLocale
)

//   TStringsOptions = set of TStringsOption;
type TStringsOptions = TSet

type TShiftState = TSet

const (
	SsShift = iota + 0
	SsAlt
	SsCtrl
	SsLeft
	SsRight
	SsMiddle
	SsDouble
	// Extra additions
	SsMeta
	SsSuper
	SsHyper
	SsAltGr
	SsCaps
	SsNum
	SsScroll
	SsTriple
	SsQuad
	SsExtra1
	SsExtra2
)

// vcl TMouseButton = (mbLeft, mbRight, mbMiddle);
// lcl TMouseButton = (mbLeft, mbRight, mbMiddle, mbExtra1, mbExtra2);
type TMouseButton int32

const (
	MbLeft = iota + 0
	MbRight
	MbMiddle

	// lcl
	mbExtra1
	mbExtra2
)

// TFillStyle = (fsSurface, fsBorder);
type TFillStyle int32

const (
	FsSurface = iota + 0
	FsBorder
)

// TFillMode = (fmAlternate, fmWinding);
type TFillMode int32

const (
	FmAlternate = iota + 0
	FmWinding
)

// TCopyMode = Longint;

// TCanvasStates = (csHandleValid, csFontValid, csPenValid, csBrushValid);
type TCanvasStates int32

const (
	CsHandleValid = iota + 0
	CsFontValid
	CsPenValid
	CsBrushValid
	//lcl
	CsRegionValid
)

// TCanvasState = set of TCanvasStates;
type TCanvasState = TSet

// TCanvasOrientation = (coLeftToRight, coRightToLeft);
type TCanvasOrientation int32

const (
	CoLeftToRight = iota + 0
	CoRightToLeft
)

// // Note: tfComposited only supported by ThemeServices.DrawText
// TTextFormats = (tfBottom, tfCalcRect, tfCenter, tfEditControl, tfEndEllipsis,
//   tfPathEllipsis, tfExpandTabs, tfExternalLeading, tfLeft, tfModifyString,
//   tfNoClip, tfNoPrefix, tfRight, tfRtlReading, tfSingleLine, tfTop,
//   tfVerticalCenter, tfWordBreak, tfHidePrefix, tfNoFullWidthCharBreak,
//   tfPrefixOnly, tfTabStop, tfWordEllipsis, tfComposited);
type TTextFormats int32

const (
	TfBottom = iota + 0
	TfCalcRect
	TfCenter
	TfEditControl
	TfEndEllipsis
	TfPathEllipsis
	TfExpandTabs
	TfExternalLeading
	TfLeft
	TfModifyString
	TfNoClip
	TfNoPrefix
	TfRight
	TfRtlReading
	TfSingleLine
	TfTop
	TfVerticalCenter
	TfWordBreak
	TfHidePrefix
	TfNoFullWidthCharBreak
	TfPrefixOnly
	TfTabStop
	TfWordEllipsis
	TfComposited
)

// TTextFormat = set of TTextFormats;
type TTextFormat = TSet

// TStyleElements = set of (seFont, seClient, seBorder);
//type TStyleElements = TSet
//
//const (
//	SeFont = iota + 0
//	SeClient
//	SeBorder
//)

// TBevelCut = (bvNone, bvLowered, bvRaised, bvSpace);
type TBevelCut int32

const (
	BvNone = iota + 0
	BvLowered
	BvRaised
	BvSpace
)

// TBevelEdge = (beLeft, beTop, beRight, beBottom);
type TBevelEdge int32

const (
	BeLeft = iota + 0
	BeTop
	BeRight
	BeBottom
)

// TBevelEdges = set of TBevelEdge;
type TBevelEdges = TSet

// TBevelKind = (bkNone, bkTile, bkSoft, bkFlat);
type TBevelKind int32

const (
	BkNone = iota + 0
	BkTile
	BkSoft
	BkFlat
)

// TTickMark = (tmBottomRight, tmTopLeft, tmBoth);
type TTickMark int32

const (
	TmBottomRight = iota + 0
	TmTopLeft
	TmBoth
)

// TTickStyle = (tsNone, tsAuto, tsManual);
type TTickStyle int32

const (
	TsNone = iota + 0
	TsAuto
	TsManual
)

// TPositionToolTip = (ptNone, ptTop, ptLeft, ptBottom, ptRight);
type TPositionToolTip int32

const (
	PtNone = iota + 0
	PtTop
	PtLeft
	PtBottom
	PtRight
)

// TDateTimeKind = (dtkDate, dtkTime);
type TDateTimeKind int32

const (
	DtkDate = iota + 0
	DtkTime
	// lcl
	DtkDateTime
)

// TDTDateMode = (dmComboBox, dmUpDown);
type TDTDateMode int32

const (
	DmComboBox = iota + 0
	DmUpDown
	// lcl
	DmNone
)

// TDTDateFormat = (dfShort, dfLong);
type TDTDateFormat int32

const (
	DfShort = iota + 0
	DfLong
)

// TDTCalAlignment = (dtaLeft, dtaRight);
type TDTCalAlignment int32

const (
	DtaLeft = iota + 0
	DtaRight
	// lcl
	DtaDefault
)

// { Calendar common control support }
// TCalDayOfWeek = (dowMonday, dowTuesday, dowWednesday, dowThursday,
//   dowFriday, dowSaturday, dowSunday, dowLocaleDefault);
type TCalDayOfWeek int32

const (
	dowMonday = iota + 0
	dowTuesday
	dowWednesday
	dowThursday
	dowFriday
	dowSaturday
	dowSunday
	dowLocaleDefault
)

// TSearchType = (stWholeWord, stMatchCase);
type TSearchType int32

const (
	StWholeWord = iota + 0
	StMatchCase
)

// TSearchTypes = set of TSearchType;

type TSearchTypes = TSet

// TNumberingStyle = (nsNone, nsBullet);
type TNumberingStyle int32

const (
	NsNone = iota + 0
	NsBullte
)

// TAttributeType = (atSelected, atDefaultText);
type TAttributeType int32

const (
	AtSelected = iota + 0
	AtDefaultText
)

// TConsistentAttribute = (caBold, caColor, caFace, caItalic,
//   caSize, caStrikeOut, caUnderline, caProtected);
type TConsistentAttribute int32

const (
	CaBold = iota + 0
	CaColor
	CaFace
	CaItalic
	CaSize
	CaStrikeOut
	CaUnderline
	CaProtected
)

// TConsistentAttributes = set of TConsistentAttribute;
type TConsistentAttributes = TSet

// TIconArrangement = (iaTop, iaLeft);
type TIconArrangement int32

const (
	IaTop = iota + 0
	IaLeft
)

// THeaderStyle = (hsGradient, hsImage, hsThemed);
type THeaderStyle int32

const (
	HsGradient = iota + 0
	HsImage
	HsThemed
)

// TImageAlignment = (iaLeft, iaRight, iaTop, iaBottom, iaCenter);
type TImageAlignment int32

const (
	IiaLeft = iota + 0 // IaTop有冲突，所以增加一个i
	IiaRight
	IiaTop
	IiaBottom
	IiaCenter
)

// vcl TAnchorKind = (akLeft, akTop, akRight, akBottom);
// lcl TAnchorKind = (akTop, akLeft, akRight, akBottom);
type TAnchorKind int32

const (
	AkTop = iota + 0
	AkLeft
	AkRight
	AkBottom
)

//  TAnchors = set of TAnchorKind;
type TAnchors = TSet

// TOwnerDrawState = set of (odSelected, odGrayed, odDisabled, odChecked,
//    odFocused, odDefault, odHotLight, odInactive, odNoAccel, odNoFocusRect,
//    odReserved1, odReserved2, odComboBoxEdit);
type TOwnerDrawState = TSet

type TOwnerDrawStateType int32

const (
	OdSelected = iota + 0
	OdGrayed
	OdDisabled
	OdChecked
	OdFocused
	OdDefault
	OdHotLight
	OdInactive
	OdNoAccel
	OdNoFocusRect
	OdReserved1
	OdReserved2
	OdComboBoxEdit
	OdBackgroundPainted // item background already painted
)

//   TBitBtnKind = (bkCustom, bkOK, bkCancel, bkHelp, bkYes, bkNo, bkClose,
// bkAbort, bkRetry, bkIgnore, bkAll);
type TBitBtnKind int32

const (
	BkCustom = iota + 0
	BkOK
	BkCancel
	BkHelp
	BkYes
	BkNo
	BkClose
	BkAbort
	BkRetry
	BkIgnore
	BkAll
	BkNoToAll
	BkYesToAll
)

// TScrollBarKind = (sbHorizontal, sbVertical);
type TScrollBarKind int32

const (
	SbHorizontal = iota + 0
	SbVertical
)

// TScrollBarInc = 1..32767;
type TScrollBarInc int16

// TScrollBarStyle = (ssRegular, ssFlat, ssHotTrack);
type TScrollBarStyle int32

const (
	SsRegular = iota + 0
	SsFlat
	ssHotTrack
)

type TShapeType int32

const (
	StRectangle = iota + 0
	StSquare
	StRoundRect
	StRoundSquare
	StEllipse
	StCircle
	StSquaredDiamond
	StDiamond
	StTriangle
	StTriangleLeft
	StTriangleRight
	StTriangleDown
	StStar
	StStarDown
)

// TBevelStyle = (bsLowered, bsRaised);
type TBevelStyle int32

const (
	BsLowered = iota + 0
	BsRaised
)

// TBevelShape = (bsBox, bsFrame, bsTopLine, bsBottomLine, bsLeftLine,
// bsRightLine, bsSpacer);
type TBevelShape int32

const (
	BsBox = iota + 0
	BsFrame
	BsTopLine
	BsBottomLine
	BsLeftLine
	BsRightLine
	BsSpacer
)

// TGaugeKind = (gkText, gkHorizontalBar, gkVerticalBar, gkPie, gkNeedle);
type TGaugeKind int32

const (
	GkText = iota + 0
	GkHorizontalBar
	GkVerticalBar
	GkPie
	GkNeedle
)

//TCustomDrawTarget = (dtControl, dtItem, dtSubItem);
type TCustomDrawTarget int32

const (
	DtControl = iota + 0
	DtItem
	DtSubItem
)

//TCustomDrawStage = (cdPrePaint, cdPostPaint, cdPreErase, cdPostErase);
type TCustomDrawStage int32

const (
	CdPrePaint = iota + 0
	CdPostPaint
	CdPreErase
	cdPostErase
)

type TCustomDrawState = TSet

const (
	CdsSelected = iota + 0
	CdsGrayed
	CdsDisabled
	CdsChecked
	CdsFocused
	CdsDefault
	CdsHot
	CdsMarked
	CdsIndeterminate

	//CdsSelected = iota + 0
	//CdsGrayed
	//CdsDisabled
	//CdsChecked
	//CdsFocused
	//CdsDefault
	//CdsHot
	//CdsMarked
	//CdsIndeterminate
	//CdsShowKeyboardCues
	//CdsNearHot
	//CdsOtherSideHot
	//CdsDropHilited
)

//TTBCustomDrawFlags = set of (tbNoEdges, tbHiliteHotTrack, tbNoOffset, tbNoMark, tbNoEtchedEffect);
type TTBCustomDrawFlags = TSet

const (
	TbNoEdges = iota + 0
	TbHiliteHotTrack
	TbNoOffset
	TbNoMark
	TbNoEtchedEffect
)

//  TDisplayCode = (drBounds, drIcon, drLabel, drSelectBounds);
type TDisplayCode int32

const (
	DrBounds = iota + 0
	DrIcon
	DrLabel
	DrSelectBounds
)

//TSelectDirOpt = (sdAllowCreate, sdPerformCreate, sdPrompt);
type TSelectDirOpt int32

const (
	SdAllowCreate = iota + 0
	SdPerformCreate
	SdPrompt
)

//TSelectDirOpts = set of TSelectDirOpt;
type TSelectDirOpts = TSet

//TSelectDirExtOpt = (sdNewFolder, sdShowEdit, sdShowShares, sdNewUI, sdShowFiles,
//sdValidateDir);
type TSelectDirExtOpt int32

const (
	SdNewFolder = iota + 0
	SdShowEdit
	SdShowShares
	SdNewUI
	SdShowFiles
	SdValidateDir
)

//TSelectDirExtOpts = set of TSelectDirExtOpt;
type TSelectDirExtOpts = TSet

// TFindOption
type TFindOption = uint32

const (
	FrDown = iota + 0
	FrFindNext
	FrHideMatchCase
	FrHideWholeWord
	FrHideUpDown
	FrMatchCase
	FrDisableMatchCase
	FrDisableUpDown
	FrDisableWholeWord
	FrReplace
	FrReplaceAll
	FrWholeWord
	FrShowHelp

	// LCL
	FrEntireScope
	FrHideEntireScope
	FrPromptOnReplace
	FrHidePromptOnReplace
	FrButtonsAtBottom
)

// TFindOptions = set of TFindOption
type TFindOptions = TSet

type TDragMode int32

const (
	DmManual = iota + 0
	DmAutomatic
)

type TDragState int32

const (
	DsDragEnter = iota + 0
	DsDragLeave
	DsDragMove
)

type TDragKind int32

const (
	DkDrag = iota + 0
	DkDock
)

// Editors common support=
type TEditCharCase int32

const (
	EcNormal = iota + 0
	EcUpperCase
	EcLowerCase
)

type TEdgeBorder int32

const (
	EbLeft = iota + 0
	EbTop
	EbRight
	EbBottom
)

// set of TEdgeBorder
type TEdgeBorders = TSet

type TEdgeStyle int32

const (
	EsNone = iota + 0
	EsRaised
	EsLowered
)

type TGridDrawingStyle int32

const (
	GdsClassic = iota + 0
	GdsThemed
	GdsGradient
)

// Lazarus的Grids选项，跟Delphi有点不一样。
type TGridOption = int32

const (
	GoFixedVertLine = iota + 0
	GoFixedHorzLine
	GoVertLine
	GoHorzLine
	GoRangeSelect
	GoDrawFocusSelected
	GoRowSizing
	GoColSizing
	GoRowMoving
	GoColMoving
	GoEditing
	GoAutoAddRows
	GoTabs
	GoRowSelect
	GoAlwaysShowEditor
	GoThumbTracking
	// Additional Options
	GoColSpanning                 // Enable cellextent calcs
	GoRelaxedRowSelect            // User can see focused cell on goRowSelect
	GoDblClickAutoSize            // dblclicking columns borders (on hdrs) resize col.
	GoSmoothScroll                // Switch scrolling mode (pixel scroll is by default)
	GoFixedRowNumbering           // Ya
	GoScrollKeepVisible           // keeps focused cell visible while scrolling
	GoHeaderHotTracking           // Header cells change look when mouse is over them
	GoHeaderPushedLook            // Header cells looks pushed when clicked
	GoSelectionActive             // Setting grid.Selection moves also cell cursor
	GoFixedColSizing              // Allow to resize fixed columns
	GoDontScrollPartCell          // clicking partially visible cells will not scroll
	GoCellHints                   // show individual cell hints
	GoTruncCellHints              // show cell hints if cell text is too long
	GoCellEllipsis                // show "..." if cell text is too long
	GoAutoAddRowsSkipContentCheck //BB Also add a row (if AutoAddRows in Options) if last row is empty
	GoRowHighlight                // Highlight the current Row
)

// Delphi set of TGridOption,  Lazarus set of TGridOptionLz
type TGridOptions = TSet

type TGridDrawState = TSet

const (
	GdSelected = iota + 0
	GdFocused
	GdFixed
	GdHot
	GdPushed
	GdRowHighlight
)

type TGridScrollDirection = uint32

const (
	//SdLeft = iota + 0
	//SdRight
	SdUp = iota + SdRight
	SdDown
)

type THeaderSectionStyle int32

const (
	HsText = iota + 0
	HsOwnerDraw
)

type TLabelPosition int32

const (
	LpAbove = iota + 0
	LpBelow
	LpLeft
	LpRight
)

type TFlowStyle int32

const (
	FsLeftRightTopBottom = iota + 0
	FsRightLeftTopBottom
	FsLeftRightBottomTop
	FsRightLeftBottomTop
	FsTopBottomLeftRight
	FsBottomTopLeftRight
	FsTopBottomRightLeft
	FsBottomTopRightLeft
)

type TCoolBandMaximize int32

const (
	BmNone = iota + 0
	BmClick
	BmDblClick
)

type TMenuBreak int32

const (
	MbNone = iota + 0
	MbBreak
	MbBarBreak
)

type TSectionTrackState int32

const (
	TsTrackBegin = iota + 0
	TsTrackMove
	TsTrackEnd
)

// TControlState = set of (csLButtonDown, csClicked, csPalette,
//  csReadingState, csAlignmentNeeded, csFocusing, csCreating,
//  csPaintCopy, csCustomPaint, csDestroyingHandle, csDocking,
//  csDesignerHide, csPanning, csRecreating, csAligning, csGlassPaint,
//  csPrintClient);
type TControlState = TSet

const (
	CsLButtonDown = iota + 0
	CsClicked
	CsPalette
	CsReadingState
	CsFocusing
	CsCreating // not used, exists for Delphi compatibility
	CsPaintCopy
	CsCustomPaint
	CsDestroyingHandle
	CsDocking
	CsVisibleSetInLoading
)

type TControlStyleType int32

const (
	CsAcceptsControls            = iota + 0 // can have children in the designer
	CsCaptureMouse                          // auto capture mouse when clicked
	CsDesignInteractive                     // wants mouse events in design mode
	CsClickEvents                           // handles mouse events
	CsFramed                                // not implemented, has 3d frame
	CsSetCaption                            // if Name=Caption, changing the Name changes the Caption
	CsOpaque                                // the control paints its area completely
	CsDoubleClicks                          // understands mouse double clicks
	CsTripleClicks                          // understands mouse triple clicks
	CsQuadClicks                            // understands mouse quad clicks
	CsFixedWidth                            // cannot change its width
	CsFixedHeight                           // cannot change its height (for example combobox)
	CsNoDesignVisible                       // is invisible in the designer
	CsReplicatable                          // PaintTo works
	CsNoStdEvents                           // standard events such as mouse, key, and click events are ignored.
	CsDisplayDragImage                      // display images from dragimagelist during drag operation over control
	CsReflector                             // not implemented, the controls respond to size, focus and dlg messages - it can be used as ActiveX control under Windows
	CsActionClient                          // Action is set
	CsMenuEvents                            // not implemented
	CsNoFocus                               // control will not take focus when clicked with mouse.
	CsNeedsBorderPaint                      // not implemented
	CsParentBackground                      // tells WinXP to paint the theme background of parent on controls background
	CsDesignNoSmoothResize                  // when resizing control in the designer do not SetBounds while dragging
	CsDesignFixedBounds                     // can not be moved nor resized in designer
	CsHasDefaultAction                      // implements useful ExecuteDefaultAction
	CsHasCancelAction                       // implements useful ExecuteCancelAction
	CsNoDesignSelectable                    // can not be selected at design time
	CsOwnedChildrenNotSelectable            // child controls owned by this control are NOT selectable in the designer
	CsAutoSize0x0                           // if the preferred size is 0x0 then control is shrinked ot 0x0
	CsAutoSizeKeepChildLeft                 // when AutoSize=true do not move children horizontally
	CsAutoSizeKeepChildTop                  // when AutoSize=true do not move children vertically
	CsRequiresKeyboardInput                 // If the device has no physical keyboard then show the virtual keyboard when this control gets focus (therefore available only to TWinControl descendents)
)

type TControlStyle = TSet

type TMouseActivate int32

const (
	MaDefault = iota + 0
	MaActivate
	MaActivateAndEat
	MaNoActivate
	MaNoActivateAndEat
)

type TTaskBarProgressState int32

const (
	None = iota + 0
	Indeterminate
	Normal
	Error
	Paused
)

//TThumbButtonState = (Enabled, DismissOnClick, NoBackground, Hidden, NonInteractive);
type TThumbButtonState int32

const (
	Enabled = iota + 0
	DismissOnClick
	NoBackground
	Hidden
	NonInteractive
)

//TThumbButtonStates = set of TThumbButtonState;
type TThumbButtonStates = TSet

//TThumbTabProperty = (AppThumbAlways, AppThumbWhenActive, AppPeekAlways, AppPeekWhenActive, CustomizedPreview);
type TThumbTabProperty int32

const (
	AppThumbAlways = iota + 0
	AppThumbWhenActive
	AppPeekAlways
	AppPeekWhenActive
	CustomizedPreview
)

//TThumbTabProperties = set of TThumbTabProperty;
type TThumbTabProperties = TSet

// TBitmapHandleType
type TBitmapHandleType int32

const (
	BmDIB = iota + 0
	BmDDB
)

// TPrinterState = (psNoHandle, psHandleIC, psHandleDC);
type TPrinterState int32

const (
	PsNoDefine = iota + 0
	PsReady
	PsPrinting
	PsStopped
)

// TPrinterOrientation = (poPortrait, poLandscape);
type TPrinterOrientation int32

const (
	PoPortrait = iota + 0
	PoLandscape
	PoReverseLandscape
	PoReversePortrait
)

//TPrinterCapability = (pcCopies, pcOrientation, pcCollation);
type TPrinterCapability int32

const (
	PcCopies = iota + 0
	PcOrientation
	PcCollation
)

// Set of TPrinterCapability
type TPrinterCapabilities = TSet

type TPrinterType int32

const (
	PtLocal = iota + 0
	PtNetWork
)

type TReadyState int32

const (
	RsUninitialized = iota + 0
	RsLoading
	RsLoaded
	RsInterActive
	RsComplete
)

type TStringEncoding int32

const (
	SeUnknown = iota + 0
	SeANSI
	SeUnicode
	SeUTF8
)

type TShowInTaskbar int32

const (
	StDefault = iota + 0 // use default rules for showing taskbar item
	StAlways             // always show taskbar item for the form
	StNever              // never show taskbar item for the form
)

// Set of TTaskDialogCommonButton
type TTaskDialogCommonButtons = TSet

type TTaskDialogCommonButton int32

const (
	TcbOk = iota + 0
	TcbYes
	TcbNo
	TcbCancel
	TcbRetry
	TcbClose
)

// set of TTaskDialogFlag;
type TTaskDialogFlags = TSet

type TTaskDialogFlag int32

const (
	TfEnableHyperlinks = iota + 0
	TfUseHiconMain
	TfUseHiconFooter
	TfAllowDialogCancellation
	TfUseCommandLinks
	TfUseCommandLinksNoIcon
	TfExpandFooterArea
	TfExpandedByDefault
	TfVerificationFlagChecked
	TfShowProgressBar
	TfShowMarqueeProgressBar
	TfCallbackTimer
	TfPositionRelativeToWindow
	TfRtlLayout
	TfNoDefaultRadioButton
	TfCanBeMinimized
)

// TTaskDialogIcon = Low(Integer)..High(Integer);
type TTaskDialogIcon int32

const (
	TdiNone = iota + 0
	TdiWarning
	TdiError
	TdiInformation
	TdiShield
	TdiQuestion // Lazarus所有
)

// TProgressBarState = Vcl.ComCtrls.TProgressBarState; // for compatibility
//type TProgressBarState = TProgressBarState

type TComboBoxExStyle int32

const (
	CsExDropDown = iota + 0
	CsExSimple
	CsExDropDownList
)

type TComboBoxExStyleEx int32

const (
	CsExCaseSensitive = iota + 0
	CsExNoEditImage
	CsExNoEditImageIndent
	CsExNoSizeLimit
	CsExPathWordBreak
)

// set of TComboBoxExStyleEx;
type TComboBoxExStyles = TSet

type TAutoCompleteOption int32

const (
	AcoAutoSuggest = iota + 0
	AcoAutoAppend
	AcoSearch
	AcoFilterPrefixes
	AcoUseTab
	AcoUpDownKeyDropsList
	AcoRtlReading
)

// set of TAutoCompleteOption;
type TAutoCompleteOptions = TSet

type TDefaultMonitor int32

const (
	DmDesktop = iota + 0
	DmPrimary
	DmMainForm
	DmActiveForm
)

// TTransparentMode = (tmAuto, tmFixed);
type TTransparentMode int32

const (
	TmAuto = iota + 0
	TmFixed
)

//// libvcl
//// TAlphaFormat = (afIgnored, afDefined, afPremultiplied);
//type TAlphaFormat int32
//
//// afIgnored  The Reserved byte in the TRGBQuad is ignored.
//// afDefined  The reserved byte in the TRGBQuad contains an alpha value.
//// afPremultiplied The reserved byte in the TRGBQuad contains an alpha value. The red, green, and blue values have been premultiplied with the alpha value.
//const (
//	AfIgnored = iota + 0
//	AfDefined
//	AfPremultiplied
//)

// TDrawImageMode = (dimNormal, dimCenter, dimStretch);
type TDrawImageMode int32

const (
	DimNormal = iota + 0
	DimCenter
	DimStretch
)

type TListBoxOption int32

const (
	LboDrawFocusRect = iota + 0 // draw focus rect in case of owner drawing
)

type TListBoxOptions = TSet

type TAntialiasingMode int32

const (
	AmDontCare = iota + 0 // default antialiasing
	AmOn                  // enabled
	AmOff                 // disabled
)

type TSortDirection int32

const (
	SdAscending = iota + 0
	SdDescending
)

type TTreeViewExpandSignType int32

const (
	TvestTheme     = iota + 0 // use themed sign
	TvestPlusMinus            // use +/- sign
	TvestArrow                // use blank arrow
	TvestArrowFill            // use filled arrow
)

type TTreeViewOption int32

const (
	TvoAllowMultiselect = iota + 0
	TvoAutoExpand
	TvoAutoInsertMark
	TvoAutoItemHeight
	TvoHideSelection
	TvoHotTrack
	TvoKeepCollapsedNodes
	TvoReadOnly
	TvoRightClickSelect
	TvoRowSelect
	TvoShowButtons
	TvoShowLines
	TvoShowRoot
	TvoShowSeparators
	TvoToolTips
	TvoNoDoubleClickExpand
	TvoThemedDraw
)

type TTreeViewOptions = TSet

type TGlyphShowMode int32

const (
	GsmAlways      = iota + 0 // always show
	GsmNever                  // never show
	GsmApplication            // depends on application settings
	GsmSystem                 // depends on system settings
)

// These are LCL additions
type TCTabControlOption int32

const (
	NboShowCloseButtons = iota + 0
	NboMultiLine
	NboHidePageListPopup
	NboKeyboardTabSwitch
	NboShowAddTabButton
	NboDoChangeOnSetIndex
)

type TCTabControlOptions = TSet

type TAnchorSideReference int32

const (
	AsrTop = iota + 0
	AsrBottom
	AsrCenter
)

type TControlCellAlign int32

const (
	CcaFill = iota + 0
	CcaLeftTop
	CcaRightBottom
	CcaCenter
)

type TControlCellAligns = TSet

type TChildControlResizeStyle int32

const (
	CrsAnchorAligning        = iota + 0 // (like Delphi)
	CrsScaleChilds                      // scale children equally, keep space between children fixed
	CrsHomogenousChildResize            // enlarge children equally (i.e. by the same amount of pixel)
	CrsHomogenousSpaceResize            // enlarge space between children equally
//{$IFDEF EnablecrsSameSize}
//,CrsSameSize  // each child gets the same size (maybe one pixel difference)
//{$ENDIF}
)

type TControlChildrenLayout int32

const (
	CclNone                       = iota + 0
	CclLeftToRightThenTopToBottom // if BiDiMode <> bdLeftToRight then it becomes RightToLeft
	CclTopToBottomThenLeftToRight
)
