package configTypes

type Lock struct {
	Version int
	Path    string
}

type LockOverrides struct {
	Version *int    `validate:"required"`
	Path    *string `validate:"required"`
}

/*
Applies the given lock overrides to the lock configuration.
*/
func (l *Lock) Apply(o *LockOverrides) {
	if o == nil {
		return
	}
	if o.Version != nil {
		l.Version = *o.Version
	}
	if o.Path != nil {
		l.Path = *o.Path
	}
}
