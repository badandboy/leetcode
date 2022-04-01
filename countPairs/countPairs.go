package countPairs

func CountPairs(deliciousness []int) int {
	len1 := len(deliciousness)
	sum := 0
	for i := 0; i < len1-1; i++ {
		for j := i + 1; j < len1; j++ {
			temp := deliciousness[i] + deliciousness[j]
			if temp&(temp-1) == 0 {
				sum += 1
			}
		}
	}

	return sum

}
