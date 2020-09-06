package operators

import (
	. "object"
	. "taxonomy"
)

type InstantiationOperator struct {
	prototypeName string
	objectName    string
}

func NewInstantiationOperator(prototypeName string, objectName string) *InstantiationOperator {
	return &InstantiationOperator{prototypeName: prototypeName, objectName: objectName}
}

func (i InstantiationOperator) PrototypeName() string {
	return i.prototypeName
}

func (i InstantiationOperator) ObjectName() string {
	return i.objectName
}

func (i InstantiationOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	prototype := taxonomy.GetObject(i.prototypeName)
	if prototype != nil {
		object := NewObject(i.objectName)
		object.SetPrototype(prototype)
		taxonomy.AddObject(object)
	}
	return taxonomy
}
