# Code Clinic: Image Comparision

## Include somewhere
	* Google Deep Dream
		* google release
			* http://googleresearch.blogspot.co.uk/2015/06/inceptionism-going-deeper-into-neural.html
		* Popular Science article
			* http://www.popsci.com/turn-your-life-computers-dream-world
		* Websites - upload photos
			* http://psychic-vr-lab.com/deepdream/
			* deep dreaming lynda.jpg
				* http://psychic-vr-lab.com/deepdream/pics/1015029.html
	* http://rosettacode.org/wiki/Category:Image_processing
	* https://en.wikipedia.org/wiki/Sum_of_absolute_differences

## Introduction

`Lynda.com standard intro`

## Welcome
Comparing images is a field with much depth. Doctoral theses are written on image comparison. A quick Google search for "image comparison algorithms" reveals the complexity. `DO SEARCH - read a few algorithm names` This rabbit hole gets deep, quick.

In this code clinic, we are not going to focus on some of the many brilliant solutions for image analysis which already exist. Nor are we going to attempt to create some revolutionary new brilliant solution. Instead I'm going to create a solution which illustrates many of the powerful features of Go and best teaches you how to use Go.

Let me be clear, this is my ultimate goal: I want to teach you how to use Go, and I'm going to do it in a step-by-step fashion by walking you through an image analysis solution. To this end, I'm going to be very thorough and detailed. I would rather give you too much information, than not enough. If you're trying to learn something, and if you need more information, and if it's not there, well then, you're out of luck. I would rather provide you with all of the details so that you can use them if you need them, and if you don't, well that's what the fast-forward button is for.

Have you seen this?

I've found that not everyone knows about this.

When you're watching a video on Lynda, you can come down here and hit this button to make your video play either faster or slower. This might be helpful for you as sometimes my students complain that I talk a little fast, blam it on the caffeine, so if I'm going to fast for you, you can come down here and slow me down. Additionally, if you want to get to stuff that you don't already know, you can come down here and speed me up. This is how I watch a lot of videos. Admittedly, it takes a little getting used to, but once you get "speed watching," which is my own phrase which I just made up, kind of like "speed reading" but with video, so it's "speed watching," but once you get "speed watching," it becomes a reliable and useful tool for helping skim through material in online trainings.

So, in this image analysis code clinic for the Go programming language, we're going to be ripping apart images, comparing pixels, implementing concurrency, and rocking our cores with parallelism. I think you're going to be impressed with how easy it is to implement concurrent, parallel code in Go, and I'm excited to show you how to do it, so let's get started.  

## Problem Overview

Often, the very most important programming any of us ever do, comes before we do any programming.

I know that sounds like some weird zen koan, but this is what I mean. Before you sit down and start writing code for your program, the most important programming you can ever do, is to sit down and figure out just what it is you're trying to do.

Understanding your problem, having a plan of attack, is the first thing you need before you start writing code. And that's what we're going to do here at the outset: we're going to explore the problem and gain a thorough understanding of it, and then we're going to put together our plan of attack, the steps we're going to take to solve the problem.

Once we've done all of that, we'll start writing the code.

Obviously, if we have a stupid easy problem, we just start writing the code. What we're going to do is obvious, why go through any extra rigamorale. In this case, however, we have a fairly challenging problem. Image matching is not trivial, and our time will be well served by first sitting down and thinking through what we're going to do before we actually start trying to do anything.

So here's our problem:

We have a folder of images. Can we create an algorithm that can discern if one of the images came from the other image?

How would you solve this? Seriously, how would you solve it? Would you start looking for third-party packages? Would you dive into the standard library? Would you phone a friend? 

Would you get out a pencil and a piece of paper and start sketching things out?

Let's understand this problem first; fully; completely; utterly.

What are our images?

This is where I started.

Let me look at the images I need to identify. What are the images with which I am working? Well, I am not Google. I am not a team of computer scientists with their doctorates in image analysis. I am Todd McLeod. I am a lone dude working late into the night after his family has gone to bed. I didn't even study computer science. My bachelor's is in economics, and my master's is in business, so if that doesn't peak your interest, I don't know what will: you're either going to watch a train wreck, or you're going to see living proof that anybody who wants to learn programming, can learn it. That's the motto here at Lynda.com, and I like it. You can learn it. `verify` I learned to program just as you are learning more about programming right now: by watching videos on Lynda.com. 

Alright, so let's look at our images.

`show images`

It looks like the wedding picture has a match, and the parrot, and the statue of liberty. Though the statue of liberty has been flipped, and that is going to make our task a little bit more challenging.

One thing I noticed when looking at the way in which these files are named is that, just from their names, we can tell that certain files are sub-images of other files. You can see here we have an 'a' after this file, and this file, and this file. Well, this could be a really easy solution: read the names of the files, look for the letter 'a', and wahlah, we know which images have been taken from larger images. Though I like this solution, make things as simple as possible but not simpler, I don't think Lynda.com would be happy with it.
 
Something else which we can see here from the file names is how they're encoded. They're all jpeg. This is both a good thing, and possibly, a challenging thing. It's a good thing because we'll be comparing things which are of the same type, so to speak. They're all jpeg, and comparing a bunch of jpegs is going to be easier than comparing jpegs to gifs to pngs. So that's a good thing that they're all of the same type. What might be challenging about them being jpeg, however, is that jpeg is lossy.

So let's talk about image encoding for a second. Are you ready? Here comes my one minute lecture on how computers work. We're going to step back to learn about circuits and coding schemes, and then we'll learn about image coding schemes. Here we go:

Computers run on electricity. Electricity can be on or off. If we have one switch, we can encode two messages. Think about your porch light on halloween. If that switch is on, it means, "Come trick or treat." If it's off, it means, "Go away." One switch, two messages. What if you had two porch lights? How many messages could you convey based upon the on/off configurations of those two lights? Pause this video right now and try answering this question. Two lights, how many messages? Okay. You got it. Four. They both could be on, they both could be off, the left could be on and the right could be off, or the left could be off and the right could be on. Four different possible combinations, to which I could assign whatever coding scheme of meaning which I wanted to create. If they're both on, I could tell all of my friends that this means, "Hey I'm home, the party is on, bring the beer." If only the left one is on, I could tell my friends that this means, "We're going to the mountains tomorrow." Or whatever. Circuits. Coding schemes. That's it. That's how computers work. Circuits in various states of on-or-off mean something. We have arbitrarily created various coding schemes to define what various states of on-or-off mean. ASCII, UTF-8, these are text coding schemes. JPEG, GIF, PNG, RAW, these are image coding schemes.
   
Incidentally, instead saying on-or-off, we often just say ONE-or-ZERO. ONE is on, ZERO is off. This is why you always see ONEs and ZEROs associated with computers, and this is why the `power symbol` is what it is. Look: zero and one. Off and on. This is what you push if want your computer on, or off. This is what you push if you want your computer ONE or ZERO; on or off.

So computers run on electricity. Electricity can either be on or off. At the most fundamental level, this is how computers work: a bunch of circuits which are in some state of on-or-off, just like your porch light, and we've assigned coding schemes to various sequences of on-or-off patterns. This bunch of zeros and ones means CAPITAL W in the ASCII coding scheme, a slightly different arrangement of zeros and ones means this subtle shade of pink in the JPEG coding scheme.
 
 Now, we've learned about coding schemes, and we've learned about image coding schemes, but what was that bit about the jpeg coding scheme being lossy, and how is that going to make our image analysis job harder?
 
Image coding schemes can either be lossy or lossless. Lossy means you lose data. It's lossy. It loses stuff. When you save a file into a lossy coding scheme, you are going to lose data and you are never going to get it back. A lossless coding schemes means that, if you save a file into a lossless coding scheme, you will not lose data. It's lossless. You have lost nothing. If you've compressed a file with a lossless coding scheme, you can always get the data back.

There are trade-offs: lossless requires more data; lassy can be more compressed.

JPEG is cool because it gives us smaller file sizes, but since JPEG is a lossy coding scheme, it is throwing data away to get that smaller file size. What that means is that, when you take this image here from this larger image, and then save it as a JPEG, it is going to change the color information. It could change it a little, or it could change it a lot, depending upon the quality settings with which the excerpted file is sized.

So our task, for each one of these images, is to determine if the image came from one of the other images.

Visually we have seen that we have three matches to find, and that the statue-of-liberty photo has been flipped.

We have also seen that the images are JPEGs, which means that because JPEG is lossy, our color data will have changed to some extent between the two images - it won't just be a pure extraction of one smaller image from the larger image. There will be some distortion, and this is going to make our job harder.

I talked with many great programmers about this problem, and the phrase that kept coming up is that this is not a trivial problem. It was kind of surprising to me, actually, that this was a new idiom in the programming community: "This is not trivial." I literally heard that phrase several times, from several different people, when asking for their feedback on the approach I was taking.
 
And they were right, too. This problem is not trivial. 
 
At least not until you understand it. 
 
And that's what we're going to start doing in the next video.
 
## Problem Scale

The first thing I want to do is look at our images in photoshop.

I want to do this for two reasons:
1. first, I want to check the scale of the smaller image
2. and second, I want to visually get a sense of what we're going to try to do. 

So here's the wedding photograph, and the scale of the wedding photograph looks good. You can see here, this looks like this photo came from right here. I can try to manually align the images, there, and then I can see that smaller image on/off by clicking this "eye" icon right here. That's pretty cool. And I can switch my tools here, choose this "move tool", and I can bring up the picture "info" from the "window" drop down menu right here, and now I can see the coordinates of whatever pixel I'm pointing at. And now I can zoom in, and look at this pixel. And I can see the `pixel's coordinates` and also the `RGB` values making up that pixel.
 
This is pretty amazing.

You can get a sense of the enormity, and the impossibility, of the task which is before us.

This larger picture is made up of a bunch of pixels. We can calculate how many, as a matter of fact. We just take the resolution, which is 3867 x 2577, and we multiply the width by the height to get, 9,965,259. There are 9,965,259 pixels there. And then in the excerpted smaller image, there are, 1009 multiplied by 1831, 1,847,479 pixels. So, ultimately, we have to compare close to TWO MILLION pixels against about TEN MILLION pixels and see if there is a match.

Like I said, not trivial.

And we also have to take into consideration the fact that there will be some distortion; that we will not have a clean extraction of the smaller image from the larger image; that the color data of the smaller image will have changed when the smaller image was taken from the larger image.

So this was about the point, in the development of my solution, where I went and had a beer.

And then after I'd had another one, I came back to my desk, I got a pencil and a piece of paper, and this is what I came up with.

Take a large problem and make it smaller. Break it down.

To hell with comparing TWO MILLION against TEN MILLION. I'm just going to compare ONE against TEN MILLION.

Ok. I've greatly reduced the scope of my problem. I'm only going to compare one pixel against TEN MILLION other data points. Wait a minute. Can a computer do this? Wait, let me rephrase that question. How long will it take my computer to do this? Interestingly, and importantly, the speed at which my computer will be able to do this will be dependent upon the language in which I am working. Fortunately, for me, I know I'm working in a language which is blazingly fast. Just how fast, however, I'm not sure. TEN MILLION data points. That's a lot of data points. And then, ultimately, I'm going to have to compare every image in an entire folder of images against every other image. You can see here, in the folder of images I have been given, I have one, two, three, four, five, six, seven, eight, nine images. If each image is the same size as the wedding images, I will have 90 MILLION pixel data points against which I am going to compare just one pixel. And if I compare 9 images against each other, I will have compared one pixel against 810 MILLION pixel data points.

The scale, of this problem, is immense.

You see, this is what I'm thinking. What we have here is a needle in the haystack problem.

Let me give you an example. Let's say this is my haystack: "ABDCDEFGHHIJKLMNOPQRSTUVWXYZ". Now, let me ask you a question. Is this needle, "DEFG", in the haystack? Of course it is. That's easy. What about this needle, "GJK"?

Images are nothing more than data. At their most fundamental level, images are zeros and ones. If we come up a level in abstraction, the data of an image could be expressed in assembly, or coming up another level, in hexadecimal. We could even express images as RGB values, as we saw happening when we looked at each pixel in photoshop.

So, if my images are data, can I just compare the data in the needle against the data in the haystack? Can I use this approach to see if one collection of binary is a subset of another.

Just like I searched the alphabet haystack "ABDCDEFGHHIJKLMNOPQRSTUVWXYZ" for a matching value, in this case, "DEFG", now I am going to search the haystack of an image for the RGB value of the top left pixel in my needle.

This is where I got out a pencil and a piece of paper, and started sketching.

So, I'm going to take the top left pixel of my needle, and compare it against every pixel in the haystack. The haystack will be every other image. I'm going to look to see if the top left pixel of my needle matches a pixel in the haystack. I'm also going to introduce some sort of "fuzziness" factor. We've learned that jpeg is lossy and that, when an exceprt was taken from an image, the color data will have, to some extent, changed. So I'll have some threshold, and if the pixels are close enough in color, I'll say, "Good enough. We have a match."

That will be the starting point.

Once the top left pixel of the needle matches a pixel in the haystack, I will then begin comparing, in sequence, the underlying pixels. I will see if I have a matching row of pixels. If the top row of pixels, match, I will then compare all of the rows in the needle against the potentially corresponding pixels in the haystack. If, after comparing all of the pixels in the needle, against a region of pixels in the haystack, if the cumulative difference of all of these pixels is under some threshold, we'll say, "Alright. Here's where the needle image came from. It came from this haystack, in this exact location. And you know what else, here's the difference between all of their respective pixel color values. We can actually see how much lossiness was introduced when the one image, which became the needle, was extracted from the other image, the haystack."

Now this approach is not a new approach. It's called the "sum of average difference" in image analysis. You can find it here on wikipedia and read more about it. But it sure sounds like a whole hell-of a lot of processing.

Now, whenever the top left pixel in the needle gets a "close enough" match on one of the pixels in the haystack, I'm going to launch a whole new routine of processing, comparing those sequences of all of the pixels in the needle - TWO MILLION of them, remember? - against the underlying pixels in the haystack.

So, I'll be comparing one pixel against TEN MILLION, and then each time that one gets that "close enough" match, I'll be comparing TWO MILLION pixels against each other.

So, just ball-parking the worst-case scenario for the scale of the problem here:
* 9 images
* 10 million pixels per image
* compare one pixel to all of the pixels in all of the images: 90 million comparisons
* compare the top left pixel from each image to all of the pixels in all of the images: 810 million comparisons
* when there is a match, compare underlying sequences of pixels: 2 million comparisons * numberOfMatches

Now, this is a high ball-park, but even if we cut this ball-park figure in half, or even a third, or a quarter, that is still a whole hell-of a lot of processing.

In the next video, we'll see what we can do to scale it down.

## Solution Overview

We have a thorough understanding of the problem now. We have also developed an approach to solving the problem which looks promising if our computer can handle the processing. What is that approach, if we were to name it? It's the "finding the needle in the haystack" approach using the "sum of average differences."

If you're not clear on what "sum of average differences" actually means, in application, you will be clear on that very shortly.

First, however, here is the high-level overview of what we want to accomplish. This is our solution:
* open all of the images
* get the pixel data for each image
* store that pixel data in some data structure
* compare the top left pixel of each image against all of the other pixels
    * calculate the "sum of average differences"
* if the difference between two pixels is beneath some threshold, launch a sequence comparison

* stop comparing
	* if the sequences stop matching, then stop the comparison

* don't compare
	* a needle that is larger than the haystack

* don't compare
	* where the needle goes beyond the bounds of the haystack

This is what it looks like when I add in that last piece of logic. `photoshop wedding size compare` These are the only pixels in the "haystack" which could potentially match the top left pixel in the needle. You can see that here if I move the file around, if the top left pixel of the needled started here for instance, then the needled is going outside the bounds of the haystack. So that's not possible for the needle, for the top left corner of the needle, to have been taken from there.

There's one more thing I want to add into my solution: a visual representation of the processing that has occurred. For a "needle image" which has been found in a "haystack image", I want to create another image which shows exactly, to the pixel, where that needle matches in the haystack. I also want to show to what extent each pixel in the needle matched the corresponding pixel in the haystack.

Alright, are you ready to see the end results?

When explaining a solution to a problem, I like sharing the solution early on so that everyone can then see what it is we're working towards.

So here I have it, the solution to the image analysis problem.

How long do you think it will take to run?

I'm going to ask my computer to do all of that processing and then create some images.

Do you think it will take two seconds, ten seconds, ten minutes?

It's kind of fun to speculate about. We've talked about looking at hundreds of millions of data points. As for myself, I know that it would probably take several life-times for me to personally look at that many data points. But for a computer, for this computer, well, it's going to happen more quickly than that. I will tell you that much right now. And to me, this is one of the most amazing accomplishments of the human race, this incredible feat of engineering, to have created a machine which we can even think about asking to do such processing. Alright, enough blabbing, let's go.

* go build
	* get my binary executable
* ./file-name
* now it's running

Done.

`some number` seconds.

Incredible.

`this many data point comparisons`

And let's look at the images which were created.

Very cool.