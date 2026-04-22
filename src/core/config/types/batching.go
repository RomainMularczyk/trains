package configTypes

type Batching struct {
	TokenLimit int
	UnitLimit  int
}

type BatchingOverrides struct {
	TokenLimit *int
	UnitLimit  *int
}

/*
Applies the given batching overrides to the batching configuration.
*/
func (b *Batching) Apply(o *BatchingOverrides) {
	if o == nil {
		return
	}
	if o.TokenLimit != nil {
		b.TokenLimit = *o.TokenLimit
	}
	if o.UnitLimit != nil {
		b.UnitLimit = *o.UnitLimit
	}
}
