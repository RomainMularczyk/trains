package configTypes

type Lock struct {
	Version int
	Path    string
}

type LockOverrides struct {
	Version *int
	Path    *string
}
