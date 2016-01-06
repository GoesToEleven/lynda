func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

return 4 * factorial(4-1)

return 4 * 3 * factorial(3-1)

return 4 * 3 * 2 * factorial(2-1)

return 4 * 3 * 2 * 1 * factorial(1-1)

return 4 * 3 * 2 * 1 * 1

return 24