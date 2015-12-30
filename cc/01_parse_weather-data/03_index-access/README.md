```
The next thing I'm looking at are my results.

One of the main things to notice are the brackets on each side of each record.

These brackets are the notation for a slice. All of our data is in a slice.
Remember that ReadAll returns a [][]string.
We are looping over our [][]string.
Through each iteration of the loop, we are seeing a []string.
That's what we are being shown here:

[]string
[]string
[]string
[]string
...
[]string

As in many programming languages, we can access the elements within a slice by index position.
That's what's going on with this code:

row[0]

I'm saying, give me the element that's in position zero.
In Go, again as in many programming languages, slices use a zero based index.
The element in the first position has an index of zero.
This can sometimes be funny for people just learning programming.
To get the first element, it's at position zero.
To get the second element, it's at position one.
To get the eighth element, it's at position what?
Can you answer that question?
Pause the video and ask yourself that question, then try to answer it.
With a zero based index, what is the index of the eighth element?
Alright, so here's the answer:
To get the eighth element, it's at position seven.

You'll also notice I've used added a single keyword: break.
Break will break me out of the loop so that, after 2 records, the loop will exit.
Why is it 2 records?
Because I'm getting the record with index position 0 from my [][]string,
	which will give me a []string, which I'm assigning to a variable called "row"
and then I'm getting the records with index position 1 from my [][]string,
	which will give me a []string, which I'm assigning to a variable called "row"
and then I'm accessing the element in position zero from "row"

What is going to happen when we run the code?

This is a great way to learn programming.
You ask yourself, what is going to happen when we run the code?
And then you run the code to see if what you imagined was going to happen
is what actually happened.

So again, pause the video, and ask yourself:
What is going to happen when we run the code.

Ok, have you done it?
Alright, so now let's run the code and see what happens.

RUN CODE

Well, that didn't quite work. Look at what happened there:

[date       time    	Air_Temp	Barometric_Press	Dew_Point	Relative_Humidity	Wind_Dir	Wind_Gust	Wind_Speed]

I got my header row.

Alright, so now we need to get two rows.

How would you go about doing this? Do you have a solution?

If you feel up for it, pause this video right now and think about it,
pencil it out, how would you get just the first two rows to print using
what you have learned so far about go.

Ok. Maybe you're back from pausing the video, or maybe you just want me to
give you the answer. Either way, here it is.

We remove this underscore character, put in an i, then check i and break
when i is equal to 1. Why one? Because slices have a zero based index.
In zero position, we have our header. In position one, we have our first row
of data. After position one, after we've printed our first row of data,
when i is == to 1, break.

Now let's run it.

Perfect. Look at that output:

[date       time    	Air_Temp	Barometric_Press	Dew_Point	Relative_Humidity	Wind_Dir	Wind_Gust	Wind_Speed]
[2015_01_01 00:02:43	19.50	30.62	14.78	81.60	159.78	14.00	 9.20]

We have our row.

Now we need to see if we can access one of the fields in the row.
To do that, I'm going to add in this line of code ...

fmt.Println(row[0])

... right before the break.

Now, what will happen is that on that first record of data, we're asking
for the first field in that data.

If my record, here, is already broken up into fields, somehow, magically,
luckily, then I should get back just the date.

If it's not already broken up into fields, I'm going to get back the
entire record without the brackets. In other words, the data in the first
position will just be one big string of data which isn't broken up yet into
individual fields.

Let's see what happens.

[date       time    	Air_Temp	Barometric_Press	Dew_Point	Relative_Humidity	Wind_Dir	Wind_Gust	Wind_Speed]
[2015_01_01 00:02:43	19.50	30.62	14.78	81.60	159.78	14.00	 9.20]
2015_01_01 00:02:43	19.50	30.62	14.78	81.60	159.78	14.00	 9.20

Look at that. My data in each record isn't broken up into fields, yet.

Well, that's alright. It means a little bit more work for us, but it's
going to be good work and we've got the tools to do it.

We'll see next how to break each record up into individual fields.
```