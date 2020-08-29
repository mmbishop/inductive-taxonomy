package object

type Object struct {
	prototype   *Object
	propertyMap map[string]interface{}
}

func NewObject() *Object {
	return &Object{prototype: nil, propertyMap: make(map[string]interface{})}
}

func (obj Object) Get(propertyName string) interface{} {
	value := obj.propertyMap[propertyName]
	if value == nil {
		value = obj.prototype.Get(propertyName)
	}
	return value
}

func (obj *Object) Set(propertyName string, value interface{}) {
	obj.propertyMap[propertyName] = value
}

func (obj *Object) Prototype() *Object {
	return obj.prototype
}

func (obj *Object) SetPrototype(prototype *Object) {
	obj.prototype = prototype
}

func (obj *Object) Instantiate() *Object {
	instance := NewObject()
	instance.SetPrototype(obj)
	return instance
}
