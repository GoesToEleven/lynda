```
If we look a the documentation for csv.NewReader we see that
a NewReader returns a pointer to a reader.

A pointer is simply the address where this variable is stored.
Instead of passing around the value stored in a variable
we can pass around the address where the variable is stored.
That's what a pointer allows you to do.

So here we have a pointer to a reader
If we click on the *Reader in the docs, we will be taken to this:

type Reader struct {
    Comma            rune // field delimiter (set to ',' by NewReader)
    Comment          rune // comment character for start of line
    FieldsPerRecord  int  // number of expected fields per record
    LazyQuotes       bool // allow lazy quotes
    TrailingComma    bool // ignored; here for backwards compatibility
    TrimLeadingSpace bool // trim leading space
    // contains filtered or unexported fields
}

This tells us about the Reader variable. It is of type struct.
A struct is a data "structure." It allows us to gather fields of data together.
We could have a struct for the classic programming vehicle example.
And that struct would then have fields for paint color, number of doors,
leather, rearview camera. These would all be fields of data that we could
gather together and place in a data structure.

Back to our Reader struct, we see that there are several fields of data
gathered together which describe a Reader. Notice one of the fields: Comma.
The description for this field tells us that comma is a "field delimiter (set to ',' by NewReader)"

Hmm? I wonder if I can reset this field, this variable, to something else.

This is the essence of programming.

Asking questions, then trying something to see if it works.

So let's give this a try, and of course, I already know how this is going to work out.

I'm setting rdr.comma equal to the escape notation for a tab character.
I'm choosing the tab character because it looks like my data in the text file
is possibly separated by tabs.

Now let's run this file and, waaaah-lahhh. Beautiful results.

[date       time     Air_Temp Barometric_Press Dew_Point Relative_Humidity Wind_Dir Wind_Gust Wind_Speed]
[2015_01_01 00:02:43 19.50 30.62 14.78 81.60 159.78 14.00  9.20]
2015_01_01 00:02:43

I have a single field. The data in my records has been broken up into fields.

We can verify this by asking for a different field, say, air temp which is
at position 1. Again, remember, slices use a zero based index, like so many
things in programming, so here we have date and time at position zero,
notice how they came out together when we said "Give me what's at index position
zero," and so here is air temp at position one.

Let's change this chunk of code to this ...

		if i == 1 {
			fmt.Println(row[1])
			break
		}

And now let's run this and ...

[date       time     Air_Temp Barometric_Press Dew_Point Relative_Humidity Wind_Dir Wind_Gust Wind_Speed]
[2015_01_01 00:02:43 19.50 30.62 14.78 81.60 159.78 14.00  9.20]
19.50

... there's our data for air temp.
```