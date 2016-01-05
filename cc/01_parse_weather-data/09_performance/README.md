1. time
	* https://godoc.org/time
	* some of the things that I like, and find useful, about pkg time
		* type Time
			* time.Now() returns type Time
				* note: pkg.Func notation
				* in the above,
					* 'time' is the package
					* 'Now()' is the func
				* look at the methods attached to type Time
					* where type Time is the "receiver"
		* type Duration
			* https://github.com/GoesToEleven/GolangTraining/search?utf8=%E2%9C%93&q=sleep
2. benchmarking
	* another more accurate way to measure the performance of your code
	* check out this:
		* https://godoc.org/testing