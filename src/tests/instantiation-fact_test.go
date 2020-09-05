package tests

import (
	. "fact"
	"github.com/stretchr/testify/assert"
	. "object"
	. "taxonomy"
	"testing"
)

type InstantiationFactTestData struct {
	instantiationFact *InstantiationFact
	taxonomy          *Taxonomy
}

var iftd InstantiationFactTestData

func instantiation_fact_is_applied_to_taxonomy(t *testing.T) {
	given_a_taxonomy_with_a_prototype()
	given_an_instantiation_fact()
	when_the_instantiation_fact_is_applied_to_the_taxonomy()
	then_the_object_instantiation_is_contained_in_the_taxonomy(t)
}

func given_a_taxonomy_with_a_prototype() {
	iftd.taxonomy = NewTaxonomy()
	iftd.taxonomy.AddObject(NewObject("APrototype"))
}

func given_an_instantiation_fact() {
	iftd.instantiationFact = NewInstantiationFact("APrototype", "AnObject")
}

func when_the_instantiation_fact_is_applied_to_the_taxonomy() {
	iftd.taxonomy = iftd.instantiationFact.Apply(iftd.taxonomy)
}

func then_the_object_instantiation_is_contained_in_the_taxonomy(t *testing.T) {
	object := iftd.taxonomy.GetObject("AnObject")
	prototype := iftd.taxonomy.GetObject("APrototype")
	assert.Equal(t, prototype, object.Prototype())
}

func TestApply(t *testing.T) {
	instantiation_fact_is_applied_to_taxonomy(t)
}
