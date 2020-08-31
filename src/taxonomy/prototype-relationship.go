package taxonomy

import . "object"

type PrototypeRelationship struct {
	instance  Object
	prototype *Object
}

func NewPrototypeRelationship(instance Object, prototype *Object) *PrototypeRelationship {
	return &PrototypeRelationship{instance: instance, prototype: prototype}
}

func (p PrototypeRelationship) Instance() Object {
	return p.instance
}

func (p PrototypeRelationship) Prototype() *Object {
	return p.prototype
}
