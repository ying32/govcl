package api

// 字符串到UTF8指针
func StringToUTF8Ptr(s string) *uint8 {
	temp := []byte(s)
	utf8StrArr := make([]uint8, len(temp)+1) // +1是因为Lazarus中PChar为0结属尾
	copy(utf8StrArr, temp)
	return &utf8StrArr[0]
}
