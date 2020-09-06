package taxonomy

import (
	. "object"
)

type Taxonomy struct {
	objectMap map[string]*Object
}

func NewTaxonomy() *Taxonomy {
	return &Taxonomy{make(map[string]*Object)}
}

func (t Taxonomy) GetObject(name string) *Object {
	return t.objectMap[name]
}

func (t Taxonomy) GetObjects() []*Object {
	objects := make([]*Object, 0, len(t.objectMap))
	for _, o := range t.objectMap {
		objects = append(objects, o)
	}
	return objects
}

func (t Taxonomy) GetPrototypeRelationships() []*PrototypeRelationship {
	prototypeRelationships := make([]*PrototypeRelationship, 0)
	for _, o := range t.objectMap {
		if o.Prototype() != nil {
			prototypeRelationships = append(prototypeRelationships, NewPrototypeRelationship(*o, o.Prototype()))
		}
	}
	return prototypeRelationships
}

func (t Taxonomy) GetInstances(prototypeName string) []*Object {
	var instances []*Object
	prototype := t.GetObject(prototypeName)
	if prototype != nil {
		for _, object := range t.objectMap {
			if object.Prototype() == prototype {
				instances = append(instances, object)
			}
		}
	}
	return instances
}

func (t *Taxonomy) AddObject(obj *Object) {
	t.objectMap[obj.Name()] = obj
}
