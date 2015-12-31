1. reverse
	* figuring out the logic while transposing a picture is like cutting your own hair in a mirror
		* you think you're going to move your hand one way, it goes the other way
1. a lot of close matches on statue of liberty
	* run to see list of haystack pixel points compared
1. to run the file:

		```
		go run *.go
		
		instead of
		
		go run main.go
		```
```
// transpose needle
nIdx := (n.width - 1) - (i % n.width) + (n.width * (i / n.width))

eg, an example with an image that has width 10
so pixels in slice index position: 0 - 9

i	nIdx	(n.width - 1)	(i % n.width)	(n.width * (i / n.width))
0	9		9				0
1	8		9				1
2	7		9				2
3	6		9				3
4	5		9				4
5	4		9				5
6	3		9				6
7	2		9				7
8	1		9				8
9	0		9				9
10	9		9				0
11	8		9				1
12	7		9				2
13	6		9				3
14	5		9				4
15	4		9				5
16	3		9				6
17	2		9				7
18	1		9				8
19	0		9				9
20	9		9				0
21	8		9				1
22	7		9				2
23	6		9				3
24	5		9				4
25	4		9				5
26	3		9				6
27	2		9				7
28	1		9				8
29	0		9				9
```
	

 