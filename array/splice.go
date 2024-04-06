package array

// The Splice() function adds and/or removes slice elements. The Splice() function does not overwrites the original slice.
func Splice[T comparable](slice []T, index int, howMany *int, elements ...T) ([]T, error) {
	sliceLength := len(slice)
	newIndex, err := ConvertIndex(slice, index, "index")

	if err != nil {
		return slice, err
	}

	if howMany == nil {
		defaultHowMany := 0
		howMany = &defaultHowMany
	}

	removalRange := sliceLength - newIndex
	rightSlicePartitionBeginning := *howMany + newIndex

	if rightSlicePartitionBeginning > removalRange {
		rightSlicePartitionBeginning = removalRange
	}

	leftSlicePartition := slice[:newIndex]
	newMiddleSlicePartition := elements
	rightSlicePartition := slice[rightSlicePartitionBeginning:]
	rightSlicePartitionCopy := make([]T, len(rightSlicePartition))
	copy(rightSlicePartitionCopy, rightSlicePartition) // to prevent weird slice stuff
	leftAndMiddlePartition := append(leftSlicePartition, newMiddleSlicePartition...)
	leftAndMiddleAndRightPartition := append(leftAndMiddlePartition, rightSlicePartitionCopy...)

	return leftAndMiddleAndRightPartition, nil
}