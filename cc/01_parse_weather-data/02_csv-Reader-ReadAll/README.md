#Video 1
1. next steps
	* open the file, now what?
	* I need to read the file
	* how am I going to read the file?
		* io.ReadFull
		* ioutil.ReadAll
		* ioutil.ReadFile
		* bufio.NewReader
		* csv.NewReader
			* exploring:
				* encoding/ascii85
				* looks like I could decode PDF's
	* interfaces
		* many of the functions above take a type "reader"
			* this is an interface
		* polymorphism
		* we'll learn more about this in a minute
	* preview the data
		* what needs to be read?
			* this data looks like it's tab delimited
				* delimiter: something that separates fields in a record
		* I'm going with csv.NewReader
			* we'll be able to set the delimiter to something other than a comma, as we'll see
1. reading the file
	* trace it through the docs
		* NewReader gives me a *reader
		* a *reader gives me these methods
			* Read
				* Read reads one record 
				* The record is a slice of strings with each string representing one field.
					* []string
			* ReadAll
				* reads all of the records
					* [][]string
						* A multidimensional slice, or more accurately, a two dimensional slice
						* like a database table, or an excel spreadsheet

#Video 2
1. for range
	* `lang spec` vs `effective go`
		* in this case, `effective go` is better
			* `If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.`
1. Interfaces & Polymorphism
	* interfaces
		* Interfaces are types that declare behavior 
			* This behavior is never implemented by the interface type directly, but instead by user-defined types via methods
				* the "concrete" type implements the behavior
		* Interfaces are implicitly implemented
		* reader interface
			* one of the most common interfaces
			* NewReader takes a reader
				* Read(p []byte) (n int, err error)
			* *file implements reader interface, implicitly, by having this method
				* Read(p []byte) (n int, err error)