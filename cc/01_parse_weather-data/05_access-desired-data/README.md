```
false
true
[date       time     Air_Temp Barometric_Press Dew_Point Relative_Humidity Wind_Dir Wind_Gust Wind_Speed]
[2015_01_01 00:02:43 19.50 30.62 14.78 81.60 159.78 14.00 9.20]
string string string
19.50 30.62 9.20

Trimming space

Looking at type
We'll need to convert these strings to some numeric value to work with them in go.
Go is statically typed. That's good news for you. No more sloppy code. No more
late night bugs making you pull your hair out because of dynamic type issues.
Go's strong typing means that you have to exactly know the type of your value.
Only certain types will work in certain situations. For math, you cannot use a string.
If you want to convert one type to another type, you have to specifically tell
the compiler to do this.

Printf vs Println, formatting verbs, and the fmt package
in addition to %T, which tells us the type of a value,
three of my favorites: %d, %b, %x -  printing in base 10, base 2, base 16

```