# Full adder

## Question

Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.

### Example

```go
Input: a = 1, b = 2
Output: 3
```

```go
Input: a = -2, b = 3
Output: 1
```

## Commentary

this [website](https://www.kaoriya.net/blog/2013/02/04/) refers impletentation in C.

variable "c" represents "carry".
carry means the number added on 1 digit above when a calculation result on a certain digit beyond the number it can't express.
I am indicating carry's boolean value table below.  
`c = a & b`

|  a  |  b  |  c  |
| :---: | :---: | :---: |
| 0 | 0 | 0 |
| 1 | 0 | 0 |
| 0 | 1 | 0 |
| 1 | 1 | 1 |

then variable "a" is a bit that calcuration result `a & b` that is got rid of carry. it's a xor caluculation.
|  a  |  b  |  new a  |
| :---: | :---: | :---: |
| 0 | 0 | 0 |
| 1 | 0 | 1 |
| 0 | 1 | 1 |
| 1 | 1 | 0 |

simulation of for loop.
In integer, a is 1, b is 3.

|  i  |  a  |  b  |  a & b  |  new a  |  new b  |
| :---: | :---: | :---: | :---: | :---: | :---: |
| 1 | 01 | 10 | 1 | 10 | 10 |
| 2 | 10 | 10 | 10 | 0 | 100 |
| 3 | 0 | 100 | 0 | 100 | 0 |

calculation result is a, therefore 100(bit) = 4(int)

### memo

xor  
11  
01  
10
