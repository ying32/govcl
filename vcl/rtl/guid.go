package rtl

import (
	"fmt"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// GUIDToString
func GUIDToString(guid types.TGUID) string {
	return fmt.Sprintf("{%.8X-%.4X-%.4X-%.2X%.2X-%.2X%.2X%.2X%.2X%.2X%.2X}",
		guid.D1, guid.D2, guid.D3, guid.D4[0], guid.D4[1], guid.D4[2], guid.D4[3], guid.D4[4], guid.D4[5], guid.D4[6], guid.D4[7])
	//return api.DGUIDToString(guid)
}

// StringToGUID
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// CreateGUID
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}
