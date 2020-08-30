package tests

import (
	. "fact"
	"github.com/stretchr/testify/assert"
	"testing"
)

var property Property

func property_is_created(t *testing.T) {
	when_a_property_is_created()
	then_the_property_has_a_name_and_a_value(t)
}

func when_a_property_is_created() {
	property = NewProperty("length", 7.0)
}

func then_the_property_has_a_name_and_a_value(t *testing.T) {
	assert.Equal(t, "length", property.Name())
	assert.Equal(t, 7.0, property.Value())
}

func TestNewProperty(t *testing.T) {
	property_is_created(t)
}
