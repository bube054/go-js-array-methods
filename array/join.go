package array

// The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings.
func Join(slice []string, separator *string) string {
	if separator == nil {
		defaultSeparator := ","
		separator = &defaultSeparator
	}

	sliceLength := len(slice)
	var joinedSlice string

	for i := 0; i < sliceLength; i++ {
		value := slice[i]
		if i == sliceLength-1 {
			joinedSlice += value
		} else {
			newValue := value + *separator
			joinedSlice += newValue
		}
	}

	return joinedSlice
}
