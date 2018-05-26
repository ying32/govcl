package vcl

type IComponent interface {
	IObject
	Name() string
	SetName(string)
}
