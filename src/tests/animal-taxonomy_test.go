package tests

import (
	. "fact"
	. "taxonomy"
	"testing"
)

func TestSampleAnimalTaxonomy(t *testing.T) {
	taxonomy := NewTaxonomy()
	oaf := NewObjectAdditionFact("Clyde", "", createProperties([]string{"blooded", "hair", "trunk", "tusks"}, []interface{}{"warm", true, true, true}))
	taxonomy = oaf.Apply(taxonomy)
	oaf = NewObjectAdditionFact("Stanley", "", createProperties([]string{"blooded", "hair"}, []interface{}{"cold", false}))
	taxonomy = oaf.Apply(taxonomy)
	pf := NewPromotionFact("Elephant", []string{"Clyde"})
	taxonomy = pf.Apply(taxonomy)
	pf = NewPromotionFact("Snake", []string{"Stanley"})
	taxonomy = pf.Apply(taxonomy)
	oaf = NewObjectAdditionFact("Indira", "Elephant", createProperties([]string{"tusks"}, []interface{}{false}))
	taxonomy = oaf.Apply(taxonomy)
	oaf = NewObjectAdditionFact("Chloe", "Elephant", createProperties([]string{"tusks"}, []interface{}{false}))
	taxonomy = oaf.Apply(taxonomy)
	pf = NewPromotionFact("ElephantWithTusks", []string{"Clyde"})
	taxonomy = pf.Apply(taxonomy)
	pf = NewPromotionFact("ElephantWithoutTusks", []string{"Indira", "Chloe"})
	taxonomy = pf.Apply(taxonomy)
}

func createProperties(keys []string, values []interface{}) []Property {
	properties := make([]Property, len(keys))
	for i := 0; i < len(keys); i++ {
		properties = append(properties, NewProperty(keys[i], values[i]))
	}
	return properties
}
