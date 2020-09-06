package tests

import (
	"github.com/stretchr/testify/assert"
	. "object"
	. "operators"
	. "taxonomy"
	"testing"
)

type PromotionOperatorTestData struct {
	taxonomy          *Taxonomy
	promotionOperator *PromotionOperator
}

var potd PromotionOperatorTestData

func promotion_operator_is_applied_to_one_object_without_properties(t *testing.T) {
	given_a_promotion_operator()
	when_the_promotion_operator_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy(t)
}

func promotion_operator_is_applied_to_one_object_with_properties(t *testing.T) {
	given_a_promotion_operator()
	given_an_object_to_be_promoted_with_properties()
	when_the_promotion_operator_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy_with_properties(t)
}

func promotion_operator_is_applied_to_multiple_objects_with_properties(t *testing.T) {
	given_a_promotion_operator_with_multiple_objects()
	given_multiple_objects_to_be_promoted_with_properties()
	when_the_promotion_operator_is_applied_to_the_taxonomy()
	then_the_promoted_prototype_is_in_the_taxonomy_with_an_intersection_of_the_instance_properties(t)
}

func given_a_promotion_operator() {
	potd.taxonomy = NewTaxonomy()
	potd.taxonomy.AddObject(NewObject("AnObject"))
	potd.promotionOperator = NewPromotionOperator("APrototype", []string{"AnObject"})
}

func given_a_promotion_operator_with_multiple_objects() {
	potd.taxonomy = NewTaxonomy()
	potd.taxonomy.AddObject(NewObject("Object1"))
	potd.taxonomy.AddObject(NewObject("Object2"))
	potd.promotionOperator = NewPromotionOperator("APrototype", []string{"Object1", "Object2"})
}

func given_an_object_to_be_promoted_with_properties() {
	object := potd.taxonomy.GetObject("AnObject")
	object.Set("prop1", "value1")
	object.Set("prop2", "value2")
}

func given_multiple_objects_to_be_promoted_with_properties() {
	object := potd.taxonomy.GetObject("Object1")
	object.Set("prop1", "value1")
	object.Set("prop2", "value2")
	object = potd.taxonomy.GetObject("Object2")
	object.Set("prop2", "value2")
	object.Set("prop3", "value3")
}

func when_the_promotion_operator_is_applied_to_the_taxonomy() {
	potd.taxonomy = potd.promotionOperator.Apply(potd.taxonomy)
}

func then_the_promoted_prototype_is_in_the_taxonomy(t *testing.T) {
	prototype := potd.taxonomy.GetObject("APrototype")
	object := potd.taxonomy.GetObject("AnObject")
	assert.NotNil(t, prototype)
	assert.NotNil(t, object)
	assert.Equal(t, prototype, object.Prototype())
}

func then_the_promoted_prototype_is_in_the_taxonomy_with_properties(t *testing.T) {
	prototype := potd.taxonomy.GetObject("APrototype")
	object := potd.taxonomy.GetObject("AnObject")
	assert.NotNil(t, prototype)
	assert.NotNil(t, object)
	assert.Equal(t, prototype, object.Prototype())
	assert.Equal(t, "value1", prototype.Get("prop1"))
	assert.Equal(t, "value2", prototype.Get("prop2"))
}

func then_the_promoted_prototype_is_in_the_taxonomy_with_an_intersection_of_the_instance_properties(t *testing.T) {
	prototype := potd.taxonomy.GetObject("APrototype")
	object1 := potd.taxonomy.GetObject("Object1")
	object2 := potd.taxonomy.GetObject("Object2")
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
	promotion_operator_is_applied_to_one_object_without_properties(t)
}

func TestPromotionWithOneObjectAndProperties(t *testing.T) {
	promotion_operator_is_applied_to_one_object_with_properties(t)
}

func TestPromotionWithMultipleObjectsAndProperties(t *testing.T) {
	promotion_operator_is_applied_to_multiple_objects_with_properties(t)
}
