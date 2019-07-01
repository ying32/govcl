package msrdp

import "github.com/ying32/govcl/vcl"

// 初始
func init() {
    vcl.RegisterExtEventCallback(eventCallback)
}
