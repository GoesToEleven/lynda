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