package lcof

func myPow(x float64, n int) float64 {
	var pow func(a float64, sq int) float64

	pow = func(a float64, sq int) float64 {
		res := 1.0
		for ; n != 0; n >>= 1 {
			if n&1 == 1 {
				res *= x
			}
			x *= x
		}
		return res
	}
	if n >= 0 {
		return pow(x, n)
	}

	return 1/pow(x, n)

}
