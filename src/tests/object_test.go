package tests

import (
	"github.com/stretchr/testify/assert"
	. "object"
	"testing"
)

var (
	instance *Object
	name     string
	obj      *Object
)

func object_returns_property_value_when_it_is_defined_locally(t *testing.T) {
	given_an_object()
	when_a_property_of_the_object_is_requested()
	then_the_property_value_is_returned(t)
}

func object_modifies_property_value(t *testing.T) {
	given_an_object()
	when_a_property_of_the_object_is_modified()
	then_the_property_has_a_new_value(t)
}

func object_delegates_property_value_request_when_property_value_is_not_defined_locally(t *testing.T) {
	given_an_object_with_a_prototype()
	when_a_property_of_the_object_is_requested()
	then_the_property_value_is_obtained_from_the_prototype(t)
}

func object_instantiates_itself(t *testing.T) {
	given_an_object()
	when_the_object_is_instantiated()
	then_the_object_becomes_the_prototype_of_the_instance(t)
}

func object_property_reverts_to_default_value(t *testing.T) {
	given_an_object_with_a_prototype_and_overridden_property()
	when_a_property_is_reset_to_its_default_value()
	then_the_property_value_reverts_to_its_default(t)
}

func change_to_prototype_property_affects_object_when_it_does_not_override_the_property(t *testing.T) {
	given_an_object_with_a_prototype()
	when_a_property_of_the_prototype_changes()
	then_the_object_returns_the_property_value_defined_in_the_property(t)
}

func given_an_object() {
	obj = NewObject()
	obj.Set("name", "Bugs Bunny")
}

func given_an_object_with_a_prototype() {
	obj = NewObject()
	prototype := NewObject()
	prototype.Set("name", "Character to be named later")
	obj.SetPrototype(prototype)
}

func given_an_object_with_a_prototype_and_overridden_property() {
	obj = NewObject()
	prototype := NewObject()
	prototype.Set("name", "Character to be named later")
	obj.SetPrototype(prototype)
	obj.Set("name", "Bugs Bunny")
}

func when_a_property_of_the_object_is_requested() {
	name = obj.Get("name").(string)
}

func when_a_property_of_the_object_is_modified() {
	obj.Set("name", "Yosemite Sam")
}

func when_the_object_is_instantiated() {
	instance = obj.Instantiate()
}

func when_a_property_is_reset_to_its_default_value() {
	obj.Unset("name")
}

func when_a_property_of_the_prototype_changes() {
	obj.Prototype().Set("name", "Tweety Bird")
}

func then_the_property_value_is_returned(t *testing.T) {
	assert.Equal(t, "Bugs Bunny", name)
}

func then_the_property_has_a_new_value(t *testing.T) {
	assert.Equal(t, "Yosemite Sam", obj.Get("name"))
}

func then_the_property_value_is_obtained_from_the_prototype(t *testing.T) {
	assert.Equal(t, "Character to be named later", obj.Get("name"))
}

func then_the_object_becomes_the_prototype_of_the_instance(t *testing.T) {
	assert.Equal(t, obj, instance.Prototype())
}

func then_the_property_value_reverts_to_its_default(t *testing.T) {
	assert.Equal(t, "Character to be named later", obj.Get("name"))
}

func then_the_object_returns_the_property_value_defined_in_the_property(t *testing.T) {
	assert.Equal(t, "Tweety Bird", obj.Prototype().Get("name"))
}

func TestGetProperty(t *testing.T) {
	object_returns_property_value_when_it_is_defined_locally(t)
}

func TestSetProperty(t *testing.T) {
	object_modifies_property_value(t)
}

func TestDelegateProperty(t *testing.T) {
	object_delegates_property_value_request_when_property_value_is_not_defined_locally(t)
}

func TestInstantiation(t *testing.T) {
	object_instantiates_itself(t)
}

func TestSetDefault(t *testing.T) {
	object_property_reverts_to_default_value(t)
}

func TestPrototypePropertyChange(t *testing.T) {
	change_to_prototype_property_affects_object_when_it_does_not_override_the_property(t)
}
