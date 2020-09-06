package operators

import (
	. "object"
	. "taxonomy"
)

type DistributionOperator struct {
	objectName string
}

func NewDistributionOperator(objectName string) *DistributionOperator {
	return &DistributionOperator{objectName: objectName}
}

func (do DistributionOperator) ObjectName() string {
	return do.objectName
}

func (do DistributionOperator) Apply(taxonomy *Taxonomy) *Taxonomy {
	prototype := taxonomy.GetObject(do.objectName)
	if prototype != nil {
		instances := taxonomy.GetInstances(do.objectName)
		distributeProperties(prototype, instances)
	}
	return taxonomy
}

func distributeProperties(prototype *Object, instances []*Object) {
	for _, property := range getPropertiesToDistribute(prototype, instances) {
		for _, instance := range instances {
			if instance.Properties()[property] == nil {
				prototypeVal := prototype.Properties()[property]
				instance.Set(property, prototypeVal)
				prototype.Unset(property)
			}
		}
	}
}

func getPropertiesToDistribute(prototype *Object, instances []*Object) []string {
	var propertiesToDistribute []string
	prototypePropertyKeys := getPropertyKeys(prototype)
	for _, key := range prototypePropertyKeys {
		for _, instance := range instances {
			if instance.Properties()[key] != nil {
				propertiesToDistribute = append(propertiesToDistribute, key)
				break
			}
		}
	}
	return propertiesToDistribute
}

func getPropertyKeys(object *Object) []string {
	propertyKeys := make([]string, len(object.Properties()))
	for key := range object.Properties() {
		propertyKeys = append(propertyKeys, key)
	}
	return propertyKeys
}
