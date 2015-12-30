1. calculating median
	* the median is the middle
		* eg, 1, 4, 5 --> median is 4
		* eg, 1, 2, 4, 5 --> media is (2 + 4) / 2 = 3
	* to find the median, we first have to sort all of our values
1. pkg sort
	* https://godoc.org/sort
		* easy way:
			* sort.Float64s
			* sort.Ints
			* sort.Strings
		* more involved way:
			* sort.Interface
				* https://godoc.org/sort#Interface
1. Looking at internals
	* interesting to look at how these are used together:
		* sort.Strings
		* sort.StringSlice
		* sort.Sort
	* when we look at the standard library's interal code for sort.Strings we see this being used:
		* sort.StringSlice
		* sort.Sort


```
OUTPUT:
Total Records:  114773
BEFORE SORTING
19.5
19.5
19.5
19.5
19.68
19.68
19.78
19.78
19.8
19.8
AFTER SORTING
19.22
19.22
19.24
19.24
19.26
19.26
19.26
19.26
19.28
19.28
Air Temp:	 53.70605996183675 0
BEFORE SORTING
30.62
30.62
30.61
30.61
30.61
30.61
30.61
30.61
30.61
30.61
AFTER SORTING
11.97
17.86
17.96
23.81
29.49
29.49
29.49
29.49
29.49
29.49
Barometric:	 30.044527197165593 0
BEFORE SORTING
9.2
9.2
8.6
8.6
9.4
9.4
10
10
10
10
AFTER SORTING
0
0
0
0
0
0
0
0
0
0
Wind Speed:	 6.916508238000142 0
```