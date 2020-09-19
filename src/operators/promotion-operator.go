package operators

import (
	"github.com/deckarep/golang-set"
	. "object"
	. "taxonomy"
)

type PromotionOperator struct {
	prototypeName string
	objectNames   []string
}

func NewPromotionOperator(prototypeName string, objectNames []string) *PromotionOperator {
	return &PromotionOperator{prototypeName: prototypeName, objectNames: objectNames}
}

func (po PromotionOperator) PrototypeName() string {
	return po.prototypeName
}

func (po PromotionOperator) ObjectNames() []string {
	return po.objectNames
}

func (po PromotionOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	prototype := NewObject(po.prototypeName)
	objects := getObjects(po.objectNames, taxonomy)
	propertySet := getIntersectionOfProperties(objects)
	setProperties(prototype, propertySet)
	prototype.SetPrototype(objects[0].Prototype())
	for _, object := range objects {
		removeProperties(object, propertySet)
		object.SetPrototype(prototype)
	}
	taxonomy.AddObject(prototype)
	return taxonomy
}

func getIntersectionOfProperties(objects []*Object) mapset.Set {
	var propertySet mapset.Set = nil
	for _, object := range objects {
		if propertySet == nil {
			propertySet = mapset.NewSet()
			addProperties(propertySet, object)
		} else {
			propertySet = propertySet.Intersect(getPropertySet(object))
		}
	}
	return propertySet
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
