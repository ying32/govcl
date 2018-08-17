package vcl

type IComponent interface {
	IObject
	Name() string
	SetName(string)
	FindComponent(string) *TComponent
	Tag() int
	SetTag(int)
	Components(int32) *TComponent
	Owner() *TComponent
	SetComponentIndex(int32)
	ComponentIndex() int32
	ComponentCount() int32
}
