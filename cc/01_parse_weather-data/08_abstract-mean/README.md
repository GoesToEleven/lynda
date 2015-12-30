1. Abstraction
	* look this file, and previous iteration, side-by-side
	* pause the video
		* can you translate how the abstraction is working?
1. functions
	* func
		* keyword
		* https://golang.org/ref/spec#Keywords
	* language spec can be confusing
		* https://golang.org/ref/spec#Function_types
		* https://golang.org/ref/spec#Function_declarations
		* https://golang.org/ref/spec#Function_literals
		* EBNF
			* https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_Form
	* need to know
		* func (r receiver) identifier(p parameters) (r returns) { func body }
		* capitalization matters
			* unexported / Exported
			* not-visible / Visible
			* we don't use the wording
				* private / Public
				* never mentioned in spec or effective go
1. Questions
	* Have we improved our code?
	* Our code is less readable
		* a goal of Go is to have readable code
			* this is good for maintainability
			* choose longer and readable over clever and compact
	* Performance consideration
		* I'm looping through rows 3x more than in previous iteration
		* Is this impacting my performance negatively?
	* Performance consideration
		* I'm passing a [][]string to my func
		* Am I passing a large data structure?
		* Is this going to take up a lot of memory?