package maxDistance

import "math"

func K(colors []int) int {
	if len(colors) == 0 {
		return 0
	}

	low := 0
	hight := len(colors) - 1

	if colors[low] != colors[hight] {
		return hight - low
	}

	var i, j = 0, hight - 1
	for colors[i] == colors[hight] {
		i += 1
	}

	for colors[0] == colors[j] {
		j -= 1
	}

	return int(math.Max(float64(j), float64(hight - i)))

}
