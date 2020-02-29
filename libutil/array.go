package libutil

func InArray(array []string, value string) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			index = i
			return index
		}
	}
	return index
}

func UintInArray(array []uint, value uint) int {
	index := -1
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			index = i
			return index
		}
	}
	return index
}
