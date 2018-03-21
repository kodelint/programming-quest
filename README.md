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
