package fact

type Property struct {
	name  string
	value interface{}
}

func NewProperty(name string, value interface{}) Property {
	return Property{name, value}
}

func (property Property) Name() string {
	return property.name
}

func (property Property) Value() interface{} {
	return property.value
}
