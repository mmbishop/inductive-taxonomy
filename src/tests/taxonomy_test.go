package tests

import (
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

func given_a_taxonomy_with_objects() {
	taxonomy = NewTaxonomy()
	obj1 := NewObject("Obj 1")
	obj2 := NewObject("Obj 2")
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
