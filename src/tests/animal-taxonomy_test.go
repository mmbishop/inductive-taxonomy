package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "object"
	. "operators"
	. "taxonomy"
	"testing"
)

var animalTaxonomy *Taxonomy

func TestSampleAnimalTaxonomy(t *testing.T) {
	animalTaxonomy = createInitialTaxonomy()
	marisaSeesStanleyTheKomodoDragon()
	marisaSeesClydeTheElephant()
	marisaSeesJerryTheGiraffe()
	marisaLearnsThatStanleyIsAReptile()
	marisaLearnsThatClydeIsAnElephant()
	marisaLearnsThatJerryIsAGiraffe()
	marisaLearnsThatElephantsAndGiraffesAreMammals()
	marisaSeesChloeTheElephant()
	marisaSeesClaireTheElephant()
	marisaRealizesThatSomeElephantsHaveTusksAndSomeElephantsDoNot()

	printTaxonomy(animalTaxonomy)

	checkPrototype(t, animalTaxonomy, "Mammal", "Animal")
	checkPrototype(t, animalTaxonomy, "Reptile", "Animal")
	checkPrototype(t, animalTaxonomy, "Elephant", "Mammal")
	checkPrototype(t, animalTaxonomy, "Giraffe", "Mammal")
	checkPrototype(t, animalTaxonomy, "ElephantWithTusks", "Elephant")
	checkPrototype(t, animalTaxonomy, "ElephantWithoutTusks", "Elephant")
	checkPrototype(t, animalTaxonomy, "Stanley", "Reptile")
	checkPrototype(t, animalTaxonomy, "Clyde", "ElephantWithTusks")
	checkPrototype(t, animalTaxonomy, "Chloe", "ElephantWithoutTusks")
	checkPrototype(t, animalTaxonomy, "Claire", "ElephantWithoutTusks")
	checkPrototype(t, animalTaxonomy, "Jerry", "Giraffe")
	assert.True(t, animalTaxonomy.GetObject("Clyde").Get("hasTusks").(bool))
	assert.False(t, animalTaxonomy.GetObject("Chloe").Get("hasTusks").(bool))
	assert.Equal(t, "warm", animalTaxonomy.GetObject("Jerry").Get("blooded"))
	assert.True(t, animalTaxonomy.GetObject("Claire").Get("hasTrunk").(bool))
	assert.True(t, animalTaxonomy.GetObject("Clyde").Get("hasHair").(bool))
	assert.Equal(t, "cold", animalTaxonomy.GetObject("Stanley").Get("blooded"))
}

func createInitialTaxonomy() *Taxonomy {
	taxonomy := NewTaxonomy()
	oao := NewObjectAdditionOperator("Animal", "")
	return oao.Apply(taxonomy)
}

func marisaSeesStanleyTheKomodoDragon() {
	oao := NewObjectAdditionOperator("Stanley", "Animal")
	ouo := NewObjectUpdateOperator("Stanley", createProperties(map[string]interface{}{
		"blooded": "cold",
		"hasHair": false,
	}))
	applyOperatorsToTaxonomy([]Operator{oao, ouo}, animalTaxonomy)
}

func marisaLearnsThatStanleyIsAReptile() {
	po := NewPromotionOperator("Reptile", []string{"Stanley"})
	applyOperatorsToTaxonomy([]Operator{po}, animalTaxonomy)
}

func marisaSeesClydeTheElephant() {
	oao := NewObjectAdditionOperator("Clyde", "Animal")
	ouo := NewObjectUpdateOperator("Clyde", createProperties(map[string]interface{}{
		"blooded":  "warm",
		"hasHair":  true,
		"hasTrunk": true,
		"hasTusks": true,
	}))
	applyOperatorsToTaxonomy([]Operator{oao, ouo}, animalTaxonomy)
}

func marisaSeesJerryTheGiraffe() {
	oao := NewObjectAdditionOperator("Jerry", "Animal")
	ouo := NewObjectUpdateOperator("Jerry", createProperties(map[string]interface{}{
		"blooded":     "warm",
		"hasHair":     true,
		"hasLongNeck": true,
	}))
	applyOperatorsToTaxonomy([]Operator{oao, ouo}, animalTaxonomy)
}

func marisaLearnsThatClydeIsAnElephant() {
	po := NewPromotionOperator("Elephant", []string{"Clyde"})
	applyOperatorsToTaxonomy([]Operator{po}, animalTaxonomy)
}

func marisaLearnsThatJerryIsAGiraffe() {
	po := NewPromotionOperator("Giraffe", []string{"Jerry"})
	applyOperatorsToTaxonomy([]Operator{po}, animalTaxonomy)
}

func marisaLearnsThatElephantsAndGiraffesAreMammals() {
	po := NewPromotionOperator("Mammal", []string{"Elephant", "Giraffe"})
	applyOperatorsToTaxonomy([]Operator{po}, animalTaxonomy)
}

func marisaSeesChloeTheElephant() {
	addElephant("Chloe", map[string]interface{}{"hasTusks": false})
}

func marisaSeesClaireTheElephant() {
	addElephant("Claire", map[string]interface{}{"hasTusks": false})
}

func marisaRealizesThatSomeElephantsHaveTusksAndSomeElephantsDoNot() {
	po1 := NewPromotionOperator("ElephantWithoutTusks", []string{"Chloe", "Claire"})
	po2 := NewPromotionOperator("ElephantWithTusks", []string{"Clyde"})
	do := NewDistributionOperator("Elephant")
	applyOperatorsToTaxonomy([]Operator{po1, po2, do}, animalTaxonomy)
}

func addElephant(name string, propertyMap map[string]interface{}) {
	oao := NewObjectAdditionOperator(name, "Elephant")
	ouo := NewObjectUpdateOperator(name, createProperties(propertyMap))
	applyOperatorsToTaxonomy([]Operator{oao, ouo}, animalTaxonomy)
}

func createProperties(propertyMap map[string]interface{}) []Property {
	var properties []Property
	for key, val := range propertyMap {
		properties = append(properties, NewProperty(key, val))
	}
	return properties
}

func applyOperatorsToTaxonomy(operators []Operator, taxonomy *Taxonomy) {
	for _, operator := range operators {
		taxonomy = operator.Apply(taxonomy)
	}
}

func printTaxonomy(taxonomy *Taxonomy) {
	prototypeRelationships := taxonomy.GetPrototypeRelationships()
	for _, prototypeRelationship := range prototypeRelationships {
		printObjectString(prototypeRelationship.Instance())
		fmt.Print(" -> ")
		printObjectString(*prototypeRelationship.Prototype())
		fmt.Println()
	}
}

func printObjectString(object Object) {
	fmt.Print(object.Name())
	propertyCount := len(object.Properties())
	if len(object.Properties()) > 0 {
		fmt.Print("{")
		i := 0
		for key, val := range object.Properties() {
			fmt.Print(key, ": ", val)
			if i++; i < propertyCount {
				fmt.Print(", ")
			}
		}
		fmt.Print("}")
	}
}

func checkPrototype(t *testing.T, taxonomy *Taxonomy, instanceName string, prototypeName string) {
	assert.Equal(t, taxonomy.GetObject(prototypeName), taxonomy.GetObject(instanceName).Prototype())
}
