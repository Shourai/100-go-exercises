# 100+ golang programming exercises

## 1. Level description
### Level 1	Beginner
Beginner means someone who has just gone through an introductory go course.
Normally, the answers could directly be found in the textbooks.

### Level 2	Intermediate
Intermediate means someone who has just learned go, but already has a relatively strong programming background from before.

### Level 3	Advanced.
Using rich libraries functions and data structures and algorithms.

----

## 2. Problem template

Question
Hints

----

## 3. Questions

If input is required, assume it is provided by stdin.

### Question 1
Level 1

Question:
Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included).
The numbers obtained should be printed in a comma-separated sequence on a single line.


### Question 2
Level 1

Question:
Write a program which can compute the factorial of a given numbers.
The results should be printed in a comma-separated sequence on a single line.
Suppose the following input is supplied to the program:
```
8
```
Then, the output should be:
```
40320
```

### Question 3
Level 1

Question:
With a given integral number n, write a program to generate a dictionary that contains (i, i*i) such that is an integral number between 1 and n (both included). and then the program should print the dictionary.
Suppose the following input is supplied to the program:
```
8
```
Then, the output should be:
```
{1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64}
```

### Question 4
Level 1

Question:
Write a program which accepts a sequence of comma-separated numbers from console and generate a list and a tuple which contains every number.
Suppose the following input is supplied to the program:
```
34,67,55,33,12,98
```
Then, the output should be:
```
['34', '67', '55', '33', '12', '98']
('34', '67', '55', '33', '12', '98')
```

### Question 5
Level 1

Question:
Define a class which has at least two methods:
getString: to get a string from console input
printString: to print the string in upper case.
Also please include simple test function to test the class methods.


### Question 6
Level 2

Question:
Write a program that calculates and prints the value according to the given formula:

```
Q = Square root of [(2 * C * D)/H]
```
Following are the fixed values of C and H:
C is 50. H is 30.
D is the variable whose values should be input to your program in a comma-separated sequence.
Example
Let us assume the following comma separated input sequence is given to the program:
```
100,150,180
```
The output of the program should be:
```
18,22,24
```

Hints:
If the output received is in decimal form, it should be rounded off to its nearest value (for example, if the output received is 26.0, it should be printed as 26)


### Question 7
Level 2

Question:
Write a program which takes 2 digits, X,Y as input and generates a 2-dimensional array. The element value in the i-th row and j-th column of the array should be i*j.
Note: i=0,1.., X-1; j=0,1,¡­Y-1.
Example
Suppose the following inputs are given to the program:
```
3,5
```
Then, the output of the program should be:
```
[[0, 0, 0, 0, 0], [0, 1, 2, 3, 4], [0, 2, 4, 6, 8]]
```

### Question 8
Level 2

Question:
Write a program that accepts a comma separated sequence of words as input and prints the words in a comma-separated sequence after sorting them alphabetically.
Suppose the following input is supplied to the program:
```
without,hello,bag,world
```
Then, the output should be:
```
bag,hello,without,world
```


### Question 9
Level 2

Question£º
Write a program that accepts sequence of lines as input and prints the lines after making all characters in the sentence capitalized.
Suppose the following input is supplied to the program:
```
Hello world
Practice makes perfect
```
Then, the output should be:
```
HELLO WORLD
PRACTICE MAKES PERFECT
```


### Question 10
Level 2

Question:
Write a program that accepts a sequence of whitespace separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically.
Suppose the following input is supplied to the program:
```
hello world and practice makes perfect and hello world again
```
Then, the output should be:
```
again and hello makes perfect practice world
```

### Question 11
Level 2

Question:
Write a program which accepts a sequence of comma separated 4 digit binary numbers as its input and then check whether they are divisible by 5 or not. The numbers that are divisible by 5 are to be printed in a comma separated sequence.
Example:
```
0100,0011,1010,1001
```
Then the output should be:
```
1010
```
Notes: Assume the data is input by console.


### Question 12
Level 2

Question:
Write a program, which will find all such numbers between 1000 and 3000 (both included) such that each digit of the number is an even number.
The numbers obtained should be printed in a comma-separated sequence on a single line.


### Question 13
Level 2

Question:
Write a program that accepts a sentence and calculate the number of letters and digits.
Suppose the following input is supplied to the program:
```
hello world! 123
```
Then, the output should be:
```
LETTERS 10
DIGITS 3
```


### Question 14
Level 2

Question:
Write a program that accepts a sentence and calculate the number of upper case letters and lower case letters.
Suppose the following input is supplied to the program:
```
Hello world!
```
Then, the output should be:
```
UPPER CASE 1
LOWER CASE 9
```


### Question 15
Level 2

Question:
Write a program that computes the value of a+aa+aaa+aaaa with a given digit as the value of a.
Suppose the following input is supplied to the program:
```
9
```
Then, the output should be:
```
11106
```


### Question 16
Level 2

Question:
Use a list comprehension to square each odd number in a list. The list is input by a sequence of comma-separated numbers.
Suppose the following input is supplied to the program:
```
1,2,3,4,5,6,7,8,9
```
Then, the output should be:
```
1,9,25,49,81
```

### Question 17
Level 2

Question:
Write a program that computes the net amount of a bank account based a transaction log from console input. The transaction log format is shown as following:
```
D 100
W 200
```

D means deposit while W means withdrawal.
Suppose the following input is supplied to the program:
```
D 300
D 300
W 200
D 100
```
Then, the output should be:
```
500
```


### Question 18
Level 3

Question:
A website requires the users to input username and password to register. Write a program to check the validity of password input by users.
Following are the criteria for checking the password:
1. At least 1 letter between [a-z]
2. At least 1 number between [0-9]
1. At least 1 letter between [A-Z]
3. At least 1 character from [$#@]
4. Minimum length of transaction password: 6
5. Maximum length of transaction password: 12
Your program should accept a sequence of comma separated passwords and will check them according to the above criteria. Passwords that match the criteria are to be printed, each separated by a comma.
Example
If the following passwords are given as input to the program:
```
ABd1234@1,a F1#,2w3E*,2We3345
```
Then, the output of the program should be:
```
ABd1234@1
```

### Question 19
Level 3

Question:
You are required to write a program to sort the (name, age, height) tuples by ascending order where name is string, age and height are numbers. The tuples are input by console. The sort criteria is:
1: Sort based on name;
2: Then sort based on age;
3: Then sort by score.
The priority is that name > age > score.
If the following tuples are given as input to the program:
```
Tom,19,80
John,20,90
Jony,17,91
Jony,17,93
Json,21,85
```
Then, the output of the program should be:
```
[('John', '20', '90'), ('Jony', '17', '91'), ('Jony', '17', '93'), ('Json', '21', '85'), ('Tom', '19', '80')]
```

### Question 20
Level 3

Question:
Define a class with a generator which can iterate the numbers, which are divisible by 7, between a given range 0 and n.

### Question 21
Level 3

Question
A robot moves in a plane starting from the original point (0,0).
The robot can move toward UP, DOWN, LEFT and RIGHT with a given steps.
The trace of robot movement is shown as the following:
UP 5
DOWN 3
LEFT 3
RIGHT 2
¡­
The numbers after the direction are steps.
Please write a program to compute the distance from current position after a sequence of movement and original point.
If the distance is a float, then just print the nearest integer.
Example:
If the following tuples are given as input to the program:
```
UP 5
DOWN 3
LEFT 3
RIGHT 2
```
Then, the output of the program should be:
```
2
```


### Question 22
Level 3

Question:
Write a program to compute the frequency of the words from the input.
The output should output after sorting the key alphanumerically.
Suppose the following input is supplied to the program:
```
New to Python or choosing between Python 2 and Python 3? Read Python 2 or Python 3.
```
Then, the output should be:
```
2:2
3.:1
3?:1
New:1
Python:5
Read:1
and:1
between:1
choosing:1
or:2
to:1
```

### Question 23
level 1

Question:
Write a function which can calculate square value of a number

### Question 24
Level 1

Question:

Python has many built-in functions, and if you do not know how to use it, you can read document online or find some books.
But Python has a built-in document function for every built-in functions.

Please write a program to print some Python built-in functions documents, such as abs(), int(), raw_input()

And add document for your own function

### Question 34

Define a function which can print a dictionary where the keys are numbers between 1 and 20 (inclusive) and the values are square of keys.

### Question 52
Create a method for the area and circumference of a circle.
Take the radius as input.

### Question 53
Create a method for the area and perimiter of a rectangle.
Take the width and height as input.
