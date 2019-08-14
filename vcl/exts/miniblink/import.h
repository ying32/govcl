

#include "wke.h"

inline jsValue cgowkeRunJS(wkeWebView webView, const utf8* script,) {
    return wkeRunJS(webView, script);
}