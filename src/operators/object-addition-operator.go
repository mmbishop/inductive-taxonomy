package operators

import (
	. "object"
	. "taxonomy"
)

type ObjectAdditionOperator struct {
	objectName    string
	prototypeName string
}

func NewObjectAdditionOperator(objectName string, prototypeName string) *ObjectAdditionOperator {
	return &ObjectAdditionOperator{objectName: objectName, prototypeName: prototypeName}
}

func (oao ObjectAdditionOperator) NewObjectName() string {
	return oao.objectName
}

func (oao ObjectAdditionOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	object := NewObject(oao.objectName)
	if len(oao.prototypeName) > 0 {
		prototype := taxonomy.GetObject(oao.prototypeName)
		if prototype != nil {
			object.SetPrototype(prototype)
		}
	}
	taxonomy.AddObject(object)
	return taxonomy
}
