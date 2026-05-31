package utils

// Replace empty strings with nil values
func EmptyToNil(src *string) *string {
	if *src == "" {
		return nil
	}
	return src
}
