1. calculating median
	* we need to determine if we take the middle number, or the middle two numbers and divide them by two
		* eg, 1, 4, 5 --> median is 4
		* eg, 1, 2, 4, 5 --> media is (2 + 4) / 2 = 3
1. constants of a kind
	* parallel type system
	* we saw this before with method sets
	* in the output
		* the literal value 2 becomes an int in this expression len(sorted)/2
		* the literal value 2 becomes an float64 in this expression float64(len(sorted))/2
	* constants can be typed or untyped
		* https://golang.org/ref/spec#Constants
```
OUTPUT:
len(sorted):  114773
len(sorted)/2:  57386
float64(len(sorted))/2:  57386.5
len(sorted):  114773
len(sorted)/2:  57386
float64(len(sorted))/2:  57386.5
len(sorted):  114773
len(sorted)/2:  57386
float64(len(sorted))/2:  57386.5
```