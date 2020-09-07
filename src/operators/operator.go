package operators

import . "taxonomy"

type Operator interface {
	Apply(taxonomy *Taxonomy) *Taxonomy
}
