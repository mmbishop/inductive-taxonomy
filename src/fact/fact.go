package fact

import . "taxonomy"

type Fact interface {
	ObjectName() string
	Properties() []Property
	Apply(taxonomy *Taxonomy) *Taxonomy
}
