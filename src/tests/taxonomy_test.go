package tests

import (
	. "fact"
	"github.com/stretchr/testify/assert"
	. "object"
	. "taxonomy"
	"testing"
)

var (
	object                 *Object
	objects                []*Object
	prototypeRelationships []*PrototypeRelationship
	taxonomy               *Taxonomy
)

func taxonomy_returns_all_objects(t *testing.T) {
	given_a_taxonomy_with_objects()
	when_all_objects_are_requested()
	then_all_objects_are_returned(t)
}

func taxonomy_returns_object_by_name(t *testing.T) {
	given_a_taxonomy_with_objects()
	when_a_particular_object_is_requested()
	then_that_object_is_returned(t)
}

func object_is_added_to_taxonomy(t *testing.T) {
	given_a_taxonomy_with_objects()
	when_an_object_is_added_to_the_taxonony()
	then_the_new_object_is_included_in_the_taxonomy(t)
}

func taxonomy_returns_prototype_relationships(t *testing.T) {
	given_a_taxonomy_with_objects()
	when_its_prototype_relationships_are_requested()
	then_the_prototype_relationships_are_returned(t)
}

func taxonomy_accepts_facts_adding_new_objects(t *testing.T) {
	given_an_empty_taxonomy()
	when_facts_representing_new_objects_are_asserted()
	then_the_taxonomy_contains_the_objects_added_by_the_facts(t)
}

func taxonomy_is_modified_by_asserted_facts(t *testing.T) {
	given_a_taxonomy_with_objects()
	when_facts_modifying_the_taxonomy_are_asserted()
	then_the_taxonomy_is_modified_by_the_facts(t)
}

func given_an_empty_taxonomy() {
	taxonomy = NewTaxonomy()
}

func given_a_taxonomy_with_objects() {
	taxonomy = NewTaxonomy()
	obj1 := NewObject("Obj 1")
	obj1.Set("prop1", "value1")
	obj2 := NewObject("Obj 2")
	obj2.Set("prop2", "value2")
	obj2.SetPrototype(obj1)
	taxonomy.AddObject(obj1)
	taxonomy.AddObject(obj2)
}

func when_all_objects_are_requested() {
	objects = taxonomy.GetObjects()
}

func when_a_particular_object_is_requested() {
	object = taxonomy.GetObject("Obj 2")
}

func when_an_object_is_added_to_the_taxonony() {
	taxonomy.AddObject(NewObject("Obj 3"))
}

func when_its_prototype_relationships_are_requested() {
	prototypeRelationships = taxonomy.GetPrototypeRelationships()
}

func when_facts_representing_new_objects_are_asserted() {
	fact := NewFact(NEW, "Animal", "")
	taxonomy.AcceptFact(*fact)
	fact = NewFact(UPDATE, "Mammal", "Animal")
	fact.AddProperty(NewProperty("blooded", "warm"))
	fact.AddProperty(NewProperty("hasHair", true))
	taxonomy.AcceptFact(*fact)
}

func when_facts_modifying_the_taxonomy_are_asserted() {
	fact := NewFact(NEW, "Obj 3", "Obj 1")
	fact.AddProperty(NewProperty("prop3", "value3"))
	taxonomy.AcceptFact(*fact)
	fact = NewFact(UPDATE, "Obj 2", "Obj 3")
	taxonomy.AcceptFact(*fact)
}

func then_all_objects_are_returned(t *testing.T) {
	assert.Equal(t, 2, len(objects))
	assert.Equal(t, "Obj 1", objects[0].Name())
	assert.Equal(t, "Obj 2", objects[1].Name())
}

func then_that_object_is_returned(t *testing.T) {
	assert.NotNil(t, object)
	assert.Equal(t, "Obj 2", object.Name())
}

func then_the_new_object_is_included_in_the_taxonomy(t *testing.T) {
	assert.Equal(t, 3, len(taxonomy.GetObjects()))
}

func then_the_prototype_relationships_are_returned(t *testing.T) {
	assert.Equal(t, 1, len(prototypeRelationships))
	assert.Equal(t, "Obj 2", prototypeRelationships[0].Instance().Name())
	assert.Equal(t, "Obj 1", prototypeRelationships[0].Prototype().Name())
}

func then_the_taxonomy_contains_the_objects_added_by_the_facts(t *testing.T) {
	objects := taxonomy.GetObjects()
	assert.Equal(t, 2, len(objects))
	obj := taxonomy.GetObject("Animal")
	assert.NotNil(t, obj)
	obj = taxonomy.GetObject("Mammal")
	assert.NotNil(t, obj)
	assert.Equal(t, "warm", obj.Get("blooded"))
	assert.Equal(t, true, obj.Get("hasHair"))
}

func then_the_taxonomy_is_modified_by_the_facts(t *testing.T) {
	objects := taxonomy.GetObjects()
	assert.Equal(t, 3, len(objects))
	obj1 := taxonomy.GetObject("Obj 1")
	obj2 := taxonomy.GetObject("Obj 2")
	obj3 := taxonomy.GetObject("Obj 3")
	assert.Equal(t, obj1, obj3.Prototype())
	assert.Equal(t, obj3, obj2.Prototype())
	assert.Equal(t, "value3", obj3.Get("prop3"))
}

func TestGetObjects(t *testing.T) {
	taxonomy_returns_all_objects(t)
}

func TestGetObject(t *testing.T) {
	taxonomy_returns_object_by_name(t)
}

func TestAddObject(t *testing.T) {
	object_is_added_to_taxonomy(t)
}

func TestGetPrototypeRelationships(t *testing.T) {
	taxonomy_returns_prototype_relationships(t)
}

func TestAcceptFactsWithNewObjects(t *testing.T) {
	taxonomy_accepts_facts_adding_new_objects(t)
}

func TestAcceptFactsThatModifyTheTaxonomy(t *testing.T) {
	taxonomy_is_modified_by_asserted_facts(t)
}
