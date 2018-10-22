package rtl

import (
	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// GUIDToString
func GUIDToString(guid types.TGUID) string {
	return api.DGUIDToString(guid)
}

// StringToGUID
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// CreateGUID
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}
