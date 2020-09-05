package fact

import (
	"github.com/deckarep/golang-set"
	. "object"
	. "taxonomy"
)

type PromotionFact struct {
	prototypeName string
	objectNames   []string
}

func NewPromotionFact(prototypeName string, objectNames []string) *PromotionFact {
	return &PromotionFact{prototypeName: prototypeName, objectNames: objectNames}
}

func (pf PromotionFact) PrototypeName() string {
	return pf.prototypeName
}

func (pf PromotionFact) ObjectNames() []string {
	return pf.objectNames
}

func (pf PromotionFact) Apply(taxonomy *Taxonomy) *Taxonomy {
	prototype := NewObject(pf.prototypeName)
	var propertySet mapset.Set = nil
	objects := getObjects(pf.objectNames, taxonomy)
	for _, object := range objects {
		if propertySet == nil {
			propertySet = mapset.NewSet()
			addProperties(propertySet, object)
		} else {
			propertySet = propertySet.Intersect(getPropertySet(object))
		}
	}
	setProperties(prototype, propertySet)
	prototype.SetPrototype(objects[0].Prototype())
	for _, object := range objects {
		removeProperties(object, propertySet)
		object.SetPrototype(prototype)
	}
	taxonomy.AddObject(prototype)
	return taxonomy
}

func getObjects(objectNames []string, taxonomy *Taxonomy) []*Object {
	var objects []*Object
	for _, objectName := range objectNames {
		object := taxonomy.GetObject(objectName)
		if object != nil {
			objects = append(objects, object)
		}
	}
	return objects
}

func addProperties(propertySet mapset.Set, object *Object) {
	for key, val := range object.Properties() {
		propertySet.Add(NewProperty(key, val))
	}
}

func getPropertySet(object *Object) mapset.Set {
	propertySet := mapset.NewSet()
	addProperties(propertySet, object)
	return propertySet
}

func setProperties(object *Object, propertySet mapset.Set) {
	for element := range propertySet.Iter() {
		property := element.(Property)
		object.Set(property.Name(), property.Value())
	}
}

func removeProperties(object *Object, propertySet mapset.Set) {
	for element := range propertySet.Iter() {
		property := element.(Property)
		object.Unset(property.Name())
	}
}
