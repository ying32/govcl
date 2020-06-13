package main

type TSci_NotifyHeader struct {
	/** Compatible with Windows NMHDR.
	* hwndFrom is really an environment specific window handle or pointer
	* but most clients of Scintilla.h do not have this type visible. */
	HwndFrom uint32 //hwndFrom : Pointer;  // void *hwndFrom
	IdFrom   uintptr
	Code     uint32
}

type TNotifyHeader = TSci_NotifyHeader

type TSCNotification struct {
	Nmhdr            TSci_NotifyHeader
	Position         int32   /* SCN_STYLENEEDED, SCN_MODIFIED, SCN_DWELLSTART, SCN_DWELLEND */
	Ch               int32   /* SCN_CHARADDED, SCN_KEY */
	Modifiers        int32   /* SCN_KEY */
	ModificationType int32   /* SCN_MODIFIED */
	Text             uintptr //{$IFDEF UNICODE}PWideChar{$ELSE}PAnsiChar{$ENDIF};/* SCN_MODIFIED, SCN_USERLISTSELECTION, SCN_AUTOCSELECTION */
	Length           int32   /* SCN_MODIFIED */
	LinesAdded       int32   /* SCN_MODIFIED */
	//{$IFDEF MACRO_RECORD_SUPPORT}
	//	Message               int32;	/* SCN_MACRORECORD */
	//	WParam                uptr_t;   /* SCN_MACRORECORD */
	//	LParam                sptr_t;   /* SCN_MACRORECORD */
	//{$ENDIF}
	Line                 int32 /* SCN_MODIFIED */
	FoldLevelNow         int32 /* SCN_MODIFIED */
	FoldLevelPrev        int32 /* SCN_MODIFIED */
	Margin               int32 /* SCN_MARGINCLICK */
	ListType             int32 /* SCN_USERLISTSELECTION */
	X                    int32 /* SCN_DWELLSTART, SCN_DWELLEND */
	Y                    int32 /* SCN_DWELLSTART, SCN_DWELLEND */
	Token                int32 /* SCN_MODIFIED with SC_MOD_CONTAINER */
	AnnotationLinesAdded int32 /* SC_MOD_CHANGEANNOTATION */
}
