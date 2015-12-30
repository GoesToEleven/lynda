1. zero-based index
```
	if len(sorted)%2 == 0 {
		// even number of items
		// for example: 3, 5, 8, 9
		// median is (5 + 8) / 2 = 6.5
		middle := len(sorted)/2
		higher := sorted[middle]
		lower := sorted[middle-1]
		return (higher + lower) / 2
	}
	// odd number of items
	// for example: 3, 5, 8
	// median is 5
	middle := len(sorted)/2
	return sorted[middle]
}
```
	* idiomatic go: to not have an "else" clause
		* the return determines path of execution

```
OUTPUT:
Total Records:  114773
Air Temp:	 53.70605996183675 52.84
Barometric:	 30.044527197165593 30
Wind Speed:	 6.916508238000142 5.6
```