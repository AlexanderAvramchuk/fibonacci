package fibonacci

func fib() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func Result(n int) int {
	var result int
	f := fib()
	for i := 0; i < n; i++ {
		result = f()
	}
	return result
}

func EvenSumResult(n int) int {
	var result int
	f := fib()
	for i := 0; i < n; i++ {
		if i:=f(); i%2 == 0 {
			result += i
		}
	}
	return result
}