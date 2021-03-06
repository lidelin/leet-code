package _0007_reverse_integer

const intMax = (1 << 31) - 1
const intMin = -(1 << 31)

func Reverse(x int) int {
	rev := 0

	for x != 0 {
		pop := x % 10
		x /= 10

		if (rev > intMax/10) || (rev == intMax/10 && pop > 7) {
			return 0
		}
		if (rev < intMin/10) || (rev == intMin/10 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}
