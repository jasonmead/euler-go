package main

func fib() func() int {
	a, b := 1, 1
	f := func() int {
		a, b = b, a+b
		return a
	}

	return f
}

func main() {
	sum := 0
	f := fib()
	for i := f(); i < 4000000; i = f() {
		if (i % 2 == 0) {
			sum += i
		}
	}

	println(sum)
}
