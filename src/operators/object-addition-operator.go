package operators

import (
	. "object"
	. "taxonomy"
)

type ObjectAdditionOperator struct {
	objectName    string
	prototypeName string
	properties    []Property
}

func NewObjectAdditionOperator(objectName string, prototypeName string, properties []Property) *ObjectAdditionOperator {
	return &ObjectAdditionOperator{objectName: objectName, prototypeName: prototypeName, properties: properties}
}

func (oao ObjectAdditionOperator) NewObjectName() string {
	return oao.objectName
}

func (oao ObjectAdditionOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	object := NewObject(oao.objectName)
	for _, property := range oao.properties {
		object.Set(property.Name(), property.Value())
	}
	if len(oao.prototypeName) > 0 {
		prototype := taxonomy.GetObject(oao.prototypeName)
		if prototype != nil {
			object.SetPrototype(prototype)
		}
	}
	taxonomy.AddObject(object)
	return taxonomy
}
