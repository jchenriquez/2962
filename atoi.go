package atoi

import (
	"math"
	"strings"
	"unicode"
)

// MyAtoi mimics the way that strconv.atoi
func MyAtoi(str string) (ret int) {

	var start int
	var count int
	sign := 1
	trimmed := strings.Trim(str, " ")
	if len(trimmed) == 0 {
		return 0
	}

	if trimmed[0] == '-' {
		sign *= -1
		start++
		trimmed = trimmed[1:]
	} else if trimmed[0] == '+' {
		trimmed = trimmed[1:]
	}

	for _, r := range trimmed {
		if !unicode.IsDigit(r) {
			break
		}

		count++
	}

	for _, r := range trimmed {

		if !unicode.IsDigit(r) {
			break
		}

		digit := int(r - '0')

		if ret+digit*int(math.Pow(10, float64(count-1))) > math.MaxInt32 {
			if sign >= 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		ret += digit * int(math.Pow(10, float64(count-1)))
		count--
	}

	ret *= sign

	return
}
