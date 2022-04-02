package largestOddNumber

func LargestOddNumber(str string) string {
	i := len(str) - 1
	for i >= 0 {
		if int(str[i]%2) == 1 {
			return str[0 : i+1]
		}

		i -= 1
	}

	return ""
}
