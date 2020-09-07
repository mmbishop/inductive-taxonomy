package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	. "operators"
	"os"
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

	generatePlantUML(*animalTaxonomy, "/home/mike/IdeaProjects/inductive-taxonomy/src/resources/taxonomy.puml")

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

func checkPrototype(t *testing.T, taxonomy *Taxonomy, instanceName string, prototypeName string) {
	assert.Equal(t, taxonomy.GetObject(prototypeName), taxonomy.GetObject(instanceName).Prototype())
}

func generatePlantUML(taxonomy Taxonomy, filePath string) {
	file, err := os.Create(filePath)
	checkFileError(err)
	defer file.Close()

	_, err = file.WriteString("@startuml\n")
	checkFileError(err)

	for _, prototypeRelationship := range taxonomy.GetPrototypeRelationships() {
		_, err := file.WriteString(prototypeRelationship.Prototype().Name() + " <|-- " + prototypeRelationship.Instance().Name() + "\n")
		checkFileError(err)
	}

	for _, object := range taxonomy.GetObjects() {
		for key, val := range object.Properties() {
			valueString := fmt.Sprintf("%v", val)
			_, err := file.WriteString(object.Name() + " : " + key + " = " + valueString + "\n")
			checkFileError(err)
		}
	}

	_, err = file.WriteString("@enduml\n")
	checkFileError(err)

	file.Sync()
}

func checkFileError(err error) {
	if err != nil {
		panic(err)
	}
}
