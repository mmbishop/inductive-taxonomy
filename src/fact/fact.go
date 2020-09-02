package fact

type FactType int

const (
	NEW FactType = 1 + iota
	UPDATE
	INSTANTIATE
	PROMOTE
	DELETE
)

type Fact struct {
	factType   FactType
	name       string
	similarTo  string
	properties []Property
}

func NewFact(factType FactType, name string, similarTo string) *Fact {
	return &Fact{factType, name, similarTo, []Property{}}
}

func (fact Fact) Properties() []Property {
	return fact.properties
}

func (fact *Fact) AddProperty(property Property) {
	fact.properties = append(fact.properties, property)
}

func (fact *Fact) SetProperties(properties []Property) {
	fact.properties = properties
}

func (fact Fact) FactType() FactType {
	return fact.factType
}

func (fact Fact) Name() string {
	return fact.name
}

func (fact Fact) SimilarTo() string {
	return fact.similarTo
}
