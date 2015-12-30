1. overview
	* opening a file
	* closing a file
		* defer statement
	* error checking
	* verbiage:
		* I've download a file, I'm going to need to open it
		* if I open a file, I'm going to need to close it
		* opening a file might result in an error, so I'm going to have to check for it
			* error handling
				* no "try / catch"
					* though "try / catch" can be approximated with "panic / recover"
				* primary error checking mechanism
					* deal with errors immediately, where they arise
					* deal with errors programmatically
					* errors are just another type
						* https://godoc.org/builtin#error
					* we code responses for them
1. packages.Functions()
	* import
1. documentation
	* golang.org
		* standard library
	* godoc.org
		* standard library
		* third-party packages
1. os.Open
	* multiple returns
	* pointer to a file
		* pointers
			* memory addresses
			* variable, f, is storing a memory address
			* two operators
				* &, * 
			* part of the type
				*file or *int 	
		* pass by value
			* everything is pass by value in go
	* methods
		* f.Close()
			* func (f *File) Close() error
		* This is how you attach a method to a type in Go. Notice how I said that, "attaching a method to a type." The type here is a pointer to a file. That is the type. This declaration right here, this is how we declare functions as methods in go. This is called a receiver. This receiver in particular is a pointer receiver.
		* receiver
			* pointer receiver
			* value receiver
		* method sets
			* golang spec
			* diagram
```
Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T
```
		* constants can be typed or untyped
			* parallel type system
        	* https://golang.org/ref/spec#Constants
1. error checking
	* check errors close to where they are thrown
	* keep error checking specific
		* not idiomatic: try-catch for huge blocks of code
		