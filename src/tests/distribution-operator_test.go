package tests

import (
	"github.com/stretchr/testify/assert"
	. "object"
	. "operators"
	. "taxonomy"
	"testing"
)

type DistributionOperatorTestData struct {
	distributionOperator *DistributionOperator
	instances            []*Object
	prototype            *Object
	taxonomy             *Taxonomy
}

var dotd DistributionOperatorTestData

func properties_are_distributed_to_instances_when_a_subset_of_them_have_overridden_the_properties(t *testing.T) {
	given_a_distribution_operator()
	given_instances_with_overridden_properties()
	when_the_distribution_operator_is_applied()
	then_the_properties_are_distributed_to_the_instances_that_do_not_have_them(t)
}

func properties_are_not_distributed_to_instances_when_they_have_no_properties(t *testing.T) {
	given_a_distribution_operator()
	given_instances_with_no_properties()
	when_the_distribution_operator_is_applied()
	then_all_of_the_properties_are_distributed_to_the_instance(t)
}

func given_a_distribution_operator() {
	dotd.taxonomy = NewTaxonomy()
	dotd.prototype = NewObject("Prototype")
	dotd.prototype.Set("prop1", "value1")
	dotd.prototype.Set("prop2", "value2")
	dotd.taxonomy.AddObject(dotd.prototype)
	dotd.distributionOperator = NewDistributionOperator("Prototype")
}

func given_instances_with_no_properties() {
	dotd.instances = make([]*Object, 2)
	dotd.instances[0] = NewObject("Instance1")
	dotd.instances[1] = NewObject("Instance2")
	dotd.instances[0].SetPrototype(dotd.prototype)
	dotd.instances[1].SetPrototype(dotd.prototype)
	dotd.taxonomy.AddObject(dotd.instances[0])
	dotd.taxonomy.AddObject(dotd.instances[1])
}

func given_instances_with_overridden_properties() {
	dotd.instances = make([]*Object, 2)
	dotd.instances[0] = NewObject("Instance1")
	dotd.instances[0].Set("prop1", "value1a")
	dotd.instances[1] = NewObject("Instance2")
	dotd.instances[0].SetPrototype(dotd.prototype)
	dotd.instances[1].SetPrototype(dotd.prototype)
	dotd.taxonomy.AddObject(dotd.instances[0])
	dotd.taxonomy.AddObject(dotd.instances[1])
}

func when_the_distribution_operator_is_applied() {
	dotd.taxonomy = dotd.distributionOperator.Apply(dotd.taxonomy)
}

func then_all_of_the_properties_are_distributed_to_the_instance(t *testing.T) {
	assert.Equal(t, 2, len(dotd.prototype.Properties()))
	assert.Equal(t, 0, len(dotd.instances[0].Properties()))
	assert.Equal(t, 0, len(dotd.instances[1].Properties()))
	assert.Equal(t, "value1", dotd.prototype.Get("prop1"))
	assert.Equal(t, "value2", dotd.prototype.Get("prop2"))
}

func then_the_properties_are_distributed_to_the_instances_that_do_not_have_them(t *testing.T) {
	assert.Equal(t, 1, len(dotd.prototype.Properties()))
	assert.Equal(t, 1, len(dotd.instances[0].Properties()))
	assert.Equal(t, 1, len(dotd.instances[1].Properties()))
	assert.Nil(t, dotd.prototype.Get("prop1"))
	assert.Equal(t, "value2", dotd.prototype.Get("prop2"))
	assert.Equal(t, "value1a", dotd.instances[0].Get("prop1"))
	assert.Equal(t, "value1", dotd.instances[1].Get("prop1"))
}

func TestDistributionWithInstancesHavingNoProperties(t *testing.T) {
	properties_are_not_distributed_to_instances_when_they_have_no_properties(t)
}

func TestDistributionWithInstancesOverridingProperties(t *testing.T) {
	properties_are_distributed_to_instances_when_a_subset_of_them_have_overridden_the_properties(t)
}
