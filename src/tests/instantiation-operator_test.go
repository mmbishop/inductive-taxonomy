package tests

import (
	"github.com/stretchr/testify/assert"
	. "object"
	. "operators"
	. "taxonomy"
	"testing"
)

type InstantiationOperatorTestData struct {
	instantiationOperator *InstantiationOperator
	taxonomy              *Taxonomy
}

var iotd InstantiationOperatorTestData

func instantiation_operator_is_applied_to_taxonomy(t *testing.T) {
	given_a_taxonomy_with_a_prototype()
	given_an_instantiation_operator()
	when_the_instantiation_operator_is_applied_to_the_taxonomy()
	then_the_object_instantiation_is_contained_in_the_taxonomy(t)
}

func given_a_taxonomy_with_a_prototype() {
	iotd.taxonomy = NewTaxonomy()
	iotd.taxonomy.AddObject(NewObject("APrototype"))
}

func given_an_instantiation_operator() {
	iotd.instantiationOperator = NewInstantiationOperator("APrototype", "AnObject")
}

func when_the_instantiation_operator_is_applied_to_the_taxonomy() {
	iotd.taxonomy = iotd.instantiationOperator.Apply(iotd.taxonomy)
}

func then_the_object_instantiation_is_contained_in_the_taxonomy(t *testing.T) {
	object := iotd.taxonomy.GetObject("AnObject")
	prototype := iotd.taxonomy.GetObject("APrototype")
	assert.Equal(t, prototype, object.Prototype())
}

func TestApply(t *testing.T) {
	instantiation_operator_is_applied_to_taxonomy(t)
}
