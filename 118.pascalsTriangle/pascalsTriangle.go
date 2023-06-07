package pascalstriangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	numRows--
	return recusion(numRows, []int{1})
}

func recusion(rows int, lestValues []int) [][]int {
	if rows == 0 {
		return [][]int{lestValues}
	}

	value := make([]int, 0, len(lestValues)+1)
	for i := 0; i < len(lestValues)+1; i++ {
		if i == 0 || len(lestValues) == i {
			value = append(value, 1)
			continue
		}

		sum := lestValues[i-1] + lestValues[i]
		value = append(value, sum)
	}

	rows--
	triangle := [][]int{lestValues}
	triangle = append(triangle, recusion(rows, value)...)
	return triangle
}
