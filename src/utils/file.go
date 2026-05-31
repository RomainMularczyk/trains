package utils

/*
Check if the file is empty.
*/
func IsFileEmpty(fileContent []byte) bool {
	if fileContent == nil {
		return true
	}
	return false
}
