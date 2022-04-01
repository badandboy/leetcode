package smallestDifference

import (
	"math"
	"sort"
)

func SmallestDifference(a []int, b []int) int {
	lena := len(a)
	lenb := len(b)

	sort.Ints(a)
	sort.Ints(b)
	i := 0
	j := 0
	minSum := math.Abs(float64(a[0] - b[0]))

	for i < lena && j < lenb {
		if a[i] == b[j] {
			return 0
		}

		minSum = math.Min(math.Abs(float64(a[i]-b[j])), minSum)
		if a[i] > b[j] {
			j += 1
		} else {
			i += 1
		}
	}

	return int(minSum)

}
