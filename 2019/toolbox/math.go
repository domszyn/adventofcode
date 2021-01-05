package toolbox

func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

func LcmN(nums ...int64) int64 {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}
	lcm := Lcm(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		lcm = Lcm(nums[i], lcm)
	}

	return lcm
}
