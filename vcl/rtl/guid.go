package rtl

import (
	"fmt"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// GUIDToString 将GUID转为字符形式
func GUIDToString(guid types.TGUID) string {
	return fmt.Sprintf("{%.8X-%.4X-%.4X-%.2X%.2X-%.2X%.2X%.2X%.2X%.2X%.2X}",
		guid.D1, guid.D2, guid.D3, guid.D4[0], guid.D4[1], guid.D4[2], guid.D4[3], guid.D4[4], guid.D4[5], guid.D4[6], guid.D4[7])
	//return api.DGUIDToString(guid)
}

// StringToGUID 将字符形式的GUID转为TGUID结构
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// CreateGUID 创建一个新的GUID
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}
