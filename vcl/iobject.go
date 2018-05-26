package vcl

// IObject 共公的对象接口
type IObject interface {
	Instance() uintptr
	ClassName() string
	Free()
	GetHashCode() int32
	Equals(IObject) bool
	IsValid() bool
}
