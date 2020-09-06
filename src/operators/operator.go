package operators

import . "taxonomy"

type Operator interface {
	ObjectName() string
	Properties() []Property
	Apply(taxonomy *Taxonomy) *Taxonomy
}
