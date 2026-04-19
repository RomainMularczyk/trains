package config

type Batching struct {
	TokenLimit int
	UnitLimit  int
}

type BatchingOverrides struct {
	TokenLimit *int
	UnitLimit  *int
}
