package fact

import (
	. "object"
	. "taxonomy"
)

type ObjectAdditionFact struct {
	object *Object
}

func NewObjectAdditionFact(object *Object) *ObjectAdditionFact {
	return &ObjectAdditionFact{object: object}
}

func (oaf ObjectAdditionFact) Object() *Object {
	return oaf.object
}

func (oaf ObjectAdditionFact) Apply(taxonomy *Taxonomy) *Taxonomy {
	taxonomy.AddObject(oaf.object)
	return taxonomy
}
