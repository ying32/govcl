package vcl

// 从资源中创建TFrame
func CreateResFrame(owner IComponent, fields ...interface{}) {
	resObjtBuild(2, owner, 0, fields...)
}
