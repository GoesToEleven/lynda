```
// transpose needle
nIdx := (n.width - 1) - (i % n.width) + (n.width * (i / n.width))

AN EXAMPLE
eg, an example with an image that has width 10
so pixels in slice index position: 0 - 9

i	nIdx	(n.width - 1)	(i % n.width)	(n.width * (i / n.width))
0	9		9				0				0
1	8		9				1				0
2	7		9				2				0
3	6		9				3				0
4	5		9				4				0
5	4		9				5				0
6	3		9				6				0
7	2		9				7				0
8	1		9				8				0
9	0		9				9				0
10	19		9				0				10
11	18		9				1				10
12	17		9				2				10
13	16		9				3				10
14	15		9				4				10
15	14		9				5				10
16	13		9				6				10
17	12		9				7				10
18	11		9				8				10
19	10		9				9				10
20	29		9				0				20
21	28		9				1				20
22	27		9				2				20
23	26		9				3				20
24	25		9				4				20
25	24		9				5				20
26	23		9				6				20
27	22		9				7				20
28	21		9				8				20
29	20		9				9				20
```