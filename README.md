### Coding Interview hmm..

We all have faced problems dealing with uncertainty with *Coding Interviews*. Based on my experiences and Interviews I faced, I am trying to put together a repo with **Questions**, **Answer** and if possible the `code` as well

* **Question 1**: Write a program to print **fibonacci sequence**  
**Answer**: 

    Fibonacci sequences is `1 1 2 3 5 8 13 21 ...`. It is always addition of previous `2` values  
    **Formula:**  `Xn = Xn-1 + Xn-2`  
   - Using [for loop](https://wiki.python.org/moin/ForLoop)
    
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [fib.sh](bash/fib.sh) |
     | `python`      | [fib.py](python/fib.py)|
     | `golang`      | [fib.go](golang/fib.go)|  

            sh bash/fib.sh
            Fibonacci Sequence from 0 to 10
            ===============================
            0 1 1 2 3 5 8 13 21 34


            python python/fib.py
            Fibonacci Sequence from 0 to 10
            ===============================
            0 1 1 2 3 5 8 13 21 34

            go run fib.go
            0 1 1 2 3 5 8 13 21 34

     
            
   - Using [Recursion](https://www.python-course.eu/recursive_functions.php)
   
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `python`      | [fib-recursive.py](python/fib-recursive.py)| 
     | `golang`      | [fib-recursive.go](golang/fib-recursive.go)|   

            python python/fib-recursive.py
            Fibonacci Sequence from 0 to 30
            ======================================
            1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040


            go run golang/fib-recursive.go
            Fibonacci Sequence from 0 to 10
            ================================
            0
            1
            1
            2
            3
            5
            8
            13
            21

* **Question 2**: Write a program to print **Unique Words and Count** from given file without the punctuations  
**Answer**: 
     We will use `tr -d '[:punct:]'` to remove the punctuations
 
   - Using [tr](https://en.wikipedia.org/wiki/Tr_(Unix)) for `bash`
   - Uning [Counter](https://docs.python.org/2/library/collections.html) and [String](https://docs.python.org/2/library/string.html) for `python`
    
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [count-uniq.sh](bash/count-uniq.sh) |
     | `python`      | [count-uniq.py](python/count-uniq.py) |

            sh count-uniq.sh test-data.txt 3
            17 the
            14 data
            10 we
            9 of
            7 to

            python count-uniq.py
            the: 17
            data: 14
            we: 10
            of: 9
            to: 7
