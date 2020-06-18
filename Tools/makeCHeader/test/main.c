// main.c : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include "liblcl.h" 

 
#ifdef _WIN32
// UTF8解码
char *UTF8Decode(char* str) {
    int len = MultiByteToWideChar(CP_UTF8, 0, str, -1, 0, 0);
    wchar_t* wCharBuffer = (wchar_t*)malloc(len * sizeof(wchar_t) + 1);
    MultiByteToWideChar(CP_UTF8, 0, str, -1, wCharBuffer, len);

    len = WideCharToMultiByte(CP_ACP, 0, wCharBuffer, -1, 0, 0, 0, NULL);
    char* aCharBuffer = (char*)malloc(len * sizeof(char) + 1);
    WideCharToMultiByte(CP_ACP, 0, wCharBuffer, -1, aCharBuffer, len, 0, NULL);
    free((void*)wCharBuffer);

    return aCharBuffer;
}
#endif

// 按钮单击事件
void onButton1Click(TObject sender) {
    ShowMessage("Hello world!");
}

// 文件拖放事件
void onOnDropFiles(TObject sender, void* aFileNames, intptr_t len) {
    printf("aFileNames: %p, len=%d\n", aFileNames, len);
    intptr_t i;
	// GetStringArrOf 为一个从Lazarus的string数组中获取成员的函数。
    for (i = 0; i < len; i++) {
        
#ifdef _WIN32
        // 由于liblcl使用的是UTF-8编码，所以获取或者传入的在Windows下都要经过UTF-8编/解码 
        char *filename = UTF8Decode(GetStringArrOf(aFileNames, i));
#else
	    // Linux与macOS默认都是UTF-8，则无需编/解码
        char *filename = GetStringArrOf(aFileNames, i);
#endif
        printf("file[%d]=%s\n", i+1, filename);
#ifdef _WIN32
        free((void*)filename);
#endif
    }
}

// 窗口键盘按下事件
void onFormKeyDown(TObject sender, Char* key, TShiftState shift) {
    printf("key=%d, shift=%d\n", *key, shift);
    if (*key == vkReturn) {
        ShowMessage("press Enter!");
    }

    TShiftState s = Include(0, ssAlt);
    if (InSet(s, ssAlt)) {
        printf("ssAlt1\n");
    }
    s = Exclude(s, ssAlt);
    if (!InSet(s, ssAlt)) {
        printf("ssAlt2\n");
    }
}

// 编辑框内容改变事件
void onEditChange(TObject sender) {
    printf("%s\n", Edit_GetText(sender));
}

int main()
{
    // 加载库
#ifdef _WIN32
    if (load_liblcl("liblcl.dll")) {
#endif
#ifdef __linux__
    if (load_liblcl("liblcl.so")) {
#endif
#ifdef __APPLE__
    if (load_liblcl("liblcl.dylib")) {
#endif

        // 主窗口显示在任务栏，仅Windows有效
		Application_SetMainFormOnTaskBar(Application, TRUE); 
		// 应用程序标题，影响到：比如ShowMessage的标题。
		Application_SetTitle(Application, "Hello LCL"); 
		// 初始化应用程序
        Application_Initialize(Application);

        // 创建窗口
        TForm form = Application_CreateForm(Application, FALSE);
		// 设置窗口标题
        Form_SetCaption(form, "LCL Form");
		// 设置窗口位置
        Form_SetPosition(form, poScreenCenter);

        // --- 拖放文件测试 ---
		// 接受文件拖放
        Form_SetAllowDropFiles(form, TRUE); 
	    // 拖放文件事件
        Form_SetOnDropFiles(form, onOnDropFiles);

        // 窗口优先接受按键，不受其它影响
        Form_SetKeyPreview(form, TRUE);

        // 窗口按键事件
        Form_SetOnKeyDown(form, onFormKeyDown);
        
		// ---------- 从内存流或者文件加载UI布局文件 ----------
        // 从文件加载窗口设置
        // 从流加载
        //TMemoryStream mem = NewMemoryStream();
        //MemoryStream_Write(mem, data, datalen);
		//MemoryStream_SetPosition(mem, 0); 
        //ResFormLoadFromStream(mem, form);
        //MemoryStream_Free(mem);
		
        // 从文件加载
        //ResFormLoadFromFile("./Form1.gfm", form);

        // ----------  动态创建控件 ---------- 
        // 创建一个按钮
        TButton btn = Button_Create(form);
		// 设置子父窗口
        Button_SetParent(btn, form);
		// 设置按钮单击事件
        Button_SetOnClick(btn, onButton1Click);
		// 设置按钮标题
        Button_SetCaption(btn, "button1");
		// 设置按钮在Parent的左边位置
        Button_SetLeft(btn, 100);
		// 设置按钮在Parent的顶边位置
        Button_SetTop(btn, 100);
		
        // 创建一个单行文件框（多行为TMemo）
        TEdit edit = Edit_Create(form);
		// 设置子父窗口
        Edit_SetParent(edit, form);
		// 设置左边
        Edit_SetLeft(edit, 10);
		// 设置顶边
        Edit_SetTop(edit, 10);
		// 设置编辑器内容改变事件
        Edit_SetOnChange(edit, onEditChange);

        // 运行app
        Application_Run(Application);

        // 释放liblcl库
        close_liblcl();
    }
    return 0;
}