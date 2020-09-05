package tests

import (
	. "fact"
	"github.com/stretchr/testify/assert"
	. "object"
	. "taxonomy"
	"testing"
)

type PromotionFactTestData struct {
	taxonomy      *Taxonomy
	promotionFact *PromotionFact
}

var pftd PromotionFactTestData

func promotion_fact_is_applied_to_one_object_without_properties(t *testing.T) {
	given_a_promotion_fact()
	when_the_promotion_fact_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy(t)
}

func promotion_fact_is_applied_to_one_object_with_properties(t *testing.T) {
	given_a_promotion_fact()
	given_an_object_to_be_promoted_with_properties()
	when_the_promotion_fact_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy_with_properties(t)
}

func promotion_fact_is_applied_to_multiple_objects_with_properties(t *testing.T) {
	given_a_promotion_fact_with_multiple_objects()
	given_multiple_objects_to_be_promoted_with_properties()
	when_the_promotion_fact_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy_with_an_intersection_of_the_instance_properties(t)
}

func given_a_promotion_fact() {
	pftd.taxonomy = NewTaxonomy()
	pftd.taxonomy.AddObject(NewObject("AnObject"))
	pftd.promotionFact = NewPromotionFact("APrototype", []string{"AnObject"})
}

func given_a_promotion_fact_with_multiple_objects() {
	pftd.taxonomy = NewTaxonomy()
	pftd.taxonomy.AddObject(NewObject("Object1"))
	pftd.taxonomy.AddObject(NewObject("Object2"))
	pftd.promotionFact = NewPromotionFact("APrototype", []string{"Object1", "Object2"})
}

func given_an_object_to_be_promoted_with_properties() {
	object := pftd.taxonomy.GetObject("AnObject")
	object.Set("prop1", "value1")
	object.Set("prop2", "value2")
}

func given_multiple_objects_to_be_promoted_with_properties() {
	object := pftd.taxonomy.GetObject("Object1")
	object.Set("prop1", "value1")
	object.Set("prop2", "value2")
	object = pftd.taxonomy.GetObject("Object2")
	object.Set("prop2", "value2")
	object.Set("prop3", "value3")
}

func when_the_promotion_fact_is_applied_to_the_taxonomy() {
	pftd.taxonomy = pftd.promotionFact.Apply(pftd.taxonomy)
}

func then_the_promoted_prototype_is_in_the_taxonomy(t *testing.T) {
	prototype := pftd.taxonomy.GetObject("APrototype")
	object := pftd.taxonomy.GetObject("AnObject")
	assert.NotNil(t, prototype)
	assert.NotNil(t, object)
	assert.Equal(t, prototype, object.Prototype())
}

func then_the_promoted_prototype_is_in_the_taxonomy_with_properties(t *testing.T) {
	prototype := pftd.taxonomy.GetObject("APrototype")
	object := pftd.taxonomy.GetObject("AnObject")
	assert.NotNil(t, prototype)
	assert.NotNil(t, object)
	assert.Equal(t, prototype, object.Prototype())
	assert.Equal(t, "value1", prototype.Get("prop1"))
	assert.Equal(t, "value2", prototype.Get("prop2"))
}

func then_the_promoted_prototype_is_in_the_taxonomy_with_an_intersection_of_the_instance_properties(t *testing.T) {
	prototype := pftd.taxonomy.GetObject("APrototype")
	object1 := pftd.taxonomy.GetObject("Object1")
	object2 := pftd.taxonomy.GetObject("Object2")
	assert.NotNil(t, prototype)
	assert.NotNil(t, object1)
	assert.NotNil(t, object2)
	assert.Equal(t, prototype, object1.Prototype())
	assert.Equal(t, prototype, object2.Prototype())
	assert.Equal(t, "value2", prototype.Get("prop2"))
	assert.Equal(t, "value1", object1.Get("prop1"))
	assert.Equal(t, "value3", object2.Get("prop3"))
}

func TestPromotionWithOneObjectAndNoProperties(t *testing.T) {
	promotion_fact_is_applied_to_one_object_without_properties(t)
}

func TestPromotionWithOneObjectAndProperties(t *testing.T) {
	promotion_fact_is_applied_to_one_object_with_properties(t)
}

func TestPromotionWithMultipleObjectsAndProperties(t *testing.T) {
	promotion_fact_is_applied_to_multiple_objects_with_properties(t)
}
