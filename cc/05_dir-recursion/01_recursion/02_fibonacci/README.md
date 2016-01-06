 fibonacci
 The next number is found by adding up the two numbers before it. 
 
 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ... 
 
 The 2 is found by adding the two numbers before it (1+1)
 
 ////////
 put this func into the code instead
 then run golint


func fibonacci(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
	return
}


you'll see this:
05_dir-recursion $ golint ./...
01_recursion/02_fibonacci/main.go:13:9: if block ends with a return statement, so drop this else and outdent its block

/////////////////////////////////////

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

fibonacci(4)
-------------
fibonacci(4-1) + fibonacci(4-2)
fibonacci(3) + fibonacci(2)
-------------
fibonacci(3-1) + fibonacci(3-2) + fibonacci(2-1) + fibonacci(2-2)
fibonacci(2) + fibonacci(1) + fibonacci(1) + fibonacci(0)
-------------
fibonacci(n-1) + fibonacci(n-2) + fibonacci(n-1) + fibonacci(n-2) + fibonacci(n-1) + fibonacci(n-2) + fibonacci(n-1) + fibonacci(n-2)

fibonacci(2-1) + fibonacci(2-2) + 1 + 1 + 1

fibonacci(1) + fibonacci(0) + 1 + 1 + 1
-------------
1 + 1 + 1 + 1 + 1