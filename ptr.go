package ptr

// PStr returns a pointer of a given string.
func PStr(s string) *string {
	return &s
}

// PInt returns a pointer of a given int.
func PInt(s int) *int {
	return &s
}

// PBool returns a pointer of a given bool.
func PBool(s bool) *bool {
	return &s
}
