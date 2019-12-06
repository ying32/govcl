package api

import (
	"runtime"
	"strconv"
	"strings"
)

var (
	NULLChar = nullChar()
)

func nullChar() string {
	vers := strings.Split(runtime.Version(), ".")
	if len(vers) >= 2 {
		m1, _ := strconv.Atoi(vers[0])
		m2, _ := strconv.Atoi(vers[1])
		// version >= 1.13
		if m1 >= 1 && m2 >= 13 {
			return ""
		}
	}
	return "\x00"
}
