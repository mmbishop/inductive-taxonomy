package operators

import . "taxonomy"

type ObjectUpdateOperator struct {
	objectName string
	properties []Property
}

func NewObjectUpdateOperator(objectName string, properties []Property) *ObjectUpdateOperator {
	return &ObjectUpdateOperator{objectName: objectName, properties: properties}
}

func (ouo ObjectUpdateOperator) ObjectName() string {
	return ouo.objectName
}

func (ouo ObjectUpdateOperator) Properties() []Property {
	return ouo.properties
}

func (ouo ObjectUpdateOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	object := taxonomy.GetObject(ouo.objectName)
	if object != nil {
		for _, property := range ouo.properties {
			object.Set(property.Name(), property.Value())
		}
	}
	return taxonomy
}
