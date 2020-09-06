package tests

import (
	"fmt"
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
	marisaLearnsThatClydeAndJerryAreMammals()
	marisaLearnsThatClydeIsAnElephant()
	marisaSeesChloeTheElephant()
	marisaSeesClaireTheElephant()
	marisaRealizesThatSomeElephantsHaveTusksAndSomeElephantsDoNot()
	clyde := animalTaxonomy.GetObject("Clyde")
	chloe := animalTaxonomy.GetObject("Chloe")
	fmt.Println("Clyde has tusks:", clyde.Get("hasTusks"))
	fmt.Println("Chloe has tusks:", chloe.Get("hasTusks"))
}

func createInitialTaxonomy() *Taxonomy {
	taxonomy := NewTaxonomy()
	oao := NewObjectAdditionOperator("Animal", "", nil)
	return oao.Apply(taxonomy)
}

func marisaSeesStanleyTheKomodoDragon() {
	oao := NewObjectAdditionOperator("Stanley", "Animal", createProperties(map[string]interface{}{"blooded": "cold", "hasHair": false}))
	animalTaxonomy = oao.Apply(animalTaxonomy)
}

func marisaLearnsThatStanleyIsAReptile() {
	po := NewPromotionOperator("Reptile", []string{"Stanley"})
	animalTaxonomy = po.Apply(animalTaxonomy)
}

func marisaSeesClydeTheElephant() {
	oao := NewObjectAdditionOperator("Clyde", "Animal", createProperties(map[string]interface{}{
		"blooded":  "warm",
		"hasHair":  true,
		"hasTrunk": true,
		"hasTusks": true,
	}))
	animalTaxonomy = oao.Apply(animalTaxonomy)
}

func marisaSeesJerryTheGiraffe() {
	oao := NewObjectAdditionOperator("Jerry", "Animal", createProperties(map[string]interface{}{
		"blooded":     "warm",
		"hasHair":     true,
		"hasLongNeck": true,
	}))
	animalTaxonomy = oao.Apply(animalTaxonomy)
}

func marisaLearnsThatClydeAndJerryAreMammals() {
	po := NewPromotionOperator("Mammal", []string{"Clyde", "Jerry"})
	animalTaxonomy = po.Apply(animalTaxonomy)
}

func marisaLearnsThatClydeIsAnElephant() {
	po := NewPromotionOperator("Elephant", []string{"Clyde"})
	animalTaxonomy = po.Apply(animalTaxonomy)
}

func marisaSeesChloeTheElephant() {
	addElephant("Chloe", map[string]interface{}{"hasTusks": false})
}

func marisaSeesClaireTheElephant() {
	addElephant("Claire", map[string]interface{}{"hasTusks": false})
}

func marisaRealizesThatSomeElephantsHaveTusksAndSomeElephantsDoNot() {
	po := NewPromotionOperator("ElephantWithoutTusks", []string{"Chloe", "Claire"})
	animalTaxonomy = po.Apply(animalTaxonomy)
	po = NewPromotionOperator("ElephantWithTusks", []string{"Clyde"})
	animalTaxonomy = po.Apply(animalTaxonomy)
	do := NewDistributionOperator("Elephant")
	animalTaxonomy = do.Apply(animalTaxonomy)
}

func addElephant(name string, propertyMap map[string]interface{}) {
	oao := NewObjectAdditionOperator(name, "Elephant", createProperties(propertyMap))
	animalTaxonomy = oao.Apply(animalTaxonomy)
}

func createProperties(propertyMap map[string]interface{}) []Property {
	properties := make([]Property, len(propertyMap))
	for key, val := range propertyMap {
		properties = append(properties, NewProperty(key, val))
	}
	return properties
}
