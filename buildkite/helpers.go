package buildkite

// String returns a pointer to the string value passed in
func String(v string) *string {
	return &v
}

// StringValue deferences a pointer to a string or returns "" if the pointer is nil
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
