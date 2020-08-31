package object

type Object struct {
	name        string
	prototype   *Object
	propertyMap map[string]interface{}
}

func NewObject(name string) *Object {
	return &Object{name: name, prototype: nil, propertyMap: make(map[string]interface{})}
}

func (obj Object) Name() string {
	return obj.name
}

func (obj *Object) SetName(newName string) {
	obj.name = newName
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

func (obj *Object) Unset(propertyName string) {
	delete(obj.propertyMap, propertyName)
}

func (obj Object) Prototype() *Object {
	return obj.prototype
}

func (obj *Object) SetPrototype(prototype *Object) {
	obj.prototype = prototype
}

func (obj *Object) Instantiate() *Object {
	instance := NewObject(obj.name)
	instance.SetPrototype(obj)
	return instance
}
