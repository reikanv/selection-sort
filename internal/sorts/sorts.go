package sorts

func SelectionSort(data *[]int) (compareCount int, shiftCount int) {
	length := len(*data)

	for i := 0; i < length-1; i++ {
		minIndex := i

		for j := i + 1; j < length; j++ {
			if (*data)[j] < (*data)[minIndex] {
				minIndex = j
				compareCount++
			}
		}

		if minIndex != i {
			(*data)[i], (*data)[minIndex] = (*data)[minIndex], (*data)[i]
			shiftCount++
		}
	}

	return
}

func BubbleSort(data *[]int) (compareCount int, shiftCount int) {
	length := len((*data))

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if (*data)[j] > (*data)[j+1] {
				(*data)[j], (*data)[j+1] = (*data)[j+1], (*data)[j]
				compareCount++
				shiftCount++
			}
		}
	}

	return
}
