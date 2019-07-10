package api

var (
	// 全局导入库
	libvcl = loadUILib()

	// 判断是否加载lcl库，一般用于windows上
	IsloadedLcl = false
)
