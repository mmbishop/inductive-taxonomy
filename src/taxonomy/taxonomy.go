package taxonomy

import (
	. "fact"
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

func (t *Taxonomy) AddObject(obj *Object) {
	t.objectMap[obj.Name()] = obj
}

func (t *Taxonomy) AcceptFact(fact Fact) {
	switch fact.FactType() {
	case "new":
		t.AddObject(createObject(fact, t))
	case "update":
		updateObject(fact, t)
	}
}

func createObject(fact Fact, taxonomy *Taxonomy) *Object {
	obj := NewObject(fact.Name())
	setObjectProperties(obj, fact.Properties())
	setPrototypeIfSpecified(obj, fact.SimilarTo(), taxonomy)
	return obj
}

func updateObject(fact Fact, taxonomy *Taxonomy) {
	obj := taxonomy.objectMap[fact.Name()]
	if obj != nil {
		setObjectProperties(obj, fact.Properties())
		setPrototypeIfSpecified(obj, fact.SimilarTo(), taxonomy)
	}
}

func setObjectProperties(obj *Object, properties []Property) {
	for _, property := range properties {
		obj.Set(property.Name(), property.Value())
	}
}

func setPrototypeIfSpecified(obj *Object, prototypeName string, taxonomy *Taxonomy) {
	if len(prototypeName) > 0 {
		obj.SetPrototype(taxonomy.objectMap[prototypeName])
	}
}
