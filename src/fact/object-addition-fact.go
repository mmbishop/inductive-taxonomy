package fact

import (
	. "object"
	. "taxonomy"
)

type ObjectAdditionFact struct {
	objectName    string
	prototypeName string
	properties    []Property
}

func NewObjectAdditionFact(objectName string, prototypeName string, properties []Property) *ObjectAdditionFact {
	return &ObjectAdditionFact{objectName: objectName, prototypeName: prototypeName, properties: properties}
}

func (oaf ObjectAdditionFact) NewObjectName() string {
	return oaf.objectName
}

func (oaf ObjectAdditionFact) Apply(taxonomy *Taxonomy) *Taxonomy {
	object := NewObject(oaf.objectName)
	for _, property := range oaf.properties {
		object.Set(property.Name(), property.Value())
	}
	if len(oaf.prototypeName) > 0 {
		prototype := taxonomy.GetObject(oaf.prototypeName)
		if prototype != nil {
			object.SetPrototype(prototype)
		}
	}
	taxonomy.AddObject(object)
	return taxonomy
}
