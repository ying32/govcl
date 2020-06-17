// main.c : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include "liblcl.h" 

 
#ifdef _WIN32
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

void onButton1Click(TObject sender) {
    ShowMessage("Hello world!");
}

void onOnDropFiles(TObject sender, void* aFileNames, intptr_t len) {
    printf("aFileNames: %p, len=%d\n", aFileNames, len);
    intptr_t i;
    for (i = 0; i < len; i++) {
        
#ifdef _WIN32
        char *filename = UTF8Decode(GetStringArrOf(aFileNames, i));
#else
        char *filename = GetStringArrOf(aFileNames, i);
#endif
        printf("file[%d]=%s\n", i+1, filename);
#ifdef _WIN32
        free((void*)filename);
#endif
    }
}

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
        Application_Initialize(Application);

        // 创建窗口
        TForm form = Application_CreateForm(Application, FALSE);
        Form_SetCaption(form, "LCL Form");
        Form_SetPosition(form, poScreenCenter);

        // 拖放文件测试
        Form_SetAllowDropFiles(form, TRUE); // 接受文件拖放
        Form_SetOnDropFiles(form, onOnDropFiles); // 事件

        // 窗口优先接受按键，不受其它影响
        Form_SetKeyPreview(form, TRUE);

        // 窗口按键事件
        Form_SetOnKeyDown(form, onFormKeyDown);
        
        // 从文件加载窗口设置
        // 从流加载
        //TMemoryStream mem = NewMemoryStream();
        //MemoryStream_Write(mem, data, datalen);
        //ResFormLoadFromStream(mem, form);
        //MemoryStream_Free(mem);
        // 从文件加载
        //ResFormLoadFromFile("./Form1.gfm", form);

        // --  动态创建 -- 
        // 创建按钮
        TButton btn = Button_Create(form);
        Button_SetParent(btn, form);
        Button_SetOnClick(btn, onButton1Click);
        Button_SetCaption(btn, "button1");
        Button_SetLeft(btn, 100);
        Button_SetTop(btn, 100);

        TEdit edit = Edit_Create(form);
        Edit_SetParent(edit, form);
        Edit_SetLeft(edit, 10);
        Edit_SetTop(edit, 10);
        Edit_SetOnChange(edit, onEditChange);

        // 运行app
        Application_Run(Application);

        // 释放
        close_liblcl();
    }
    return 0;
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
