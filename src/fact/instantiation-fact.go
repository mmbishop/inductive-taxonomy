package fact

import (
	. "object"
	. "taxonomy"
)

type InstantiationFact struct {
	prototypeName string
	objectName    string
}

func NewInstantiationFact(prototypeName string, objectName string) *InstantiationFact {
	return &InstantiationFact{prototypeName: prototypeName, objectName: objectName}
}

func (i InstantiationFact) PrototypeName() string {
	return i.prototypeName
}

func (i InstantiationFact) ObjectName() string {
	return i.objectName
}

func (i InstantiationFact) Apply(taxonomy *Taxonomy) *Taxonomy {
	prototype := taxonomy.GetObject(i.prototypeName)
	if prototype != nil {
		object := NewObject(i.objectName)
		object.SetPrototype(prototype)
		taxonomy.AddObject(object)
	}
	return taxonomy
}
