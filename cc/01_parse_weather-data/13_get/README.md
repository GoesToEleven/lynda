1. pkg net/http
	* https://godoc.org/net/http
	* fun
		* status teapot
			* https://godoc.org/net/http#pkg-constants
		* referer (should be: referrer)
			* https://godoc.org/net/http#Request.Referer
1. Challenge
	* How is http.Get() being called?
	* there is no Get() func at the top level.
1. Answer
```
down below
```





























	* var DefaultClient = &Client{}
		* DefaultClient is the default Client and is used by Get, Head, and Post.
	* look at the internals of the standard library
		* https://golang.org/src/net/http/client.go#L276
1. read through the code
	* get returns (resp *Response)
		* https://godoc.org/net/http#Response
		* type Response struct
			* Body io.ReadCloser
				* https://godoc.org/io#ReadCloser
					* Reader
					* Closer
	* ioutil.ReadAll
		* takes a Reader

