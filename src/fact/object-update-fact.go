package fact

import . "taxonomy"

type ObjectUpdateFact struct {
	objectName string
	properties []Property
}

func NewObjectUpdateFact(objectName string, properties []Property) *ObjectUpdateFact {
	return &ObjectUpdateFact{objectName: objectName, properties: properties}
}

func (ouf ObjectUpdateFact) ObjectName() string {
	return ouf.objectName
}

func (ouf ObjectUpdateFact) Properties() []Property {
	return ouf.properties
}

func (ouf ObjectUpdateFact) Apply(taxonomy *Taxonomy) *Taxonomy {
	object := taxonomy.GetObject(ouf.objectName)
	if object != nil {
		for _, property := range ouf.properties {
			object.Set(property.Name(), property.Value())
		}
	}
	return taxonomy
}
