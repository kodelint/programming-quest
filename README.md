### Coding Interview hmm..

**Interviewing** is a skill that you can get better at by studying, preparing, and practicing for it. It’s stressful to code in an interview, while someone scrutinizes every liine of code you write. Based on my experiences and Interviews I faced, I am sharing the insights and tips I gained  over the time. I will try to be as discriptive as possible with **Questions**, **Answers** and if possible the `code` as well

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
   - Using [Counter](https://docs.python.org/2/library/collections.html) and [String](https://docs.python.org/2/library/string.html) for `python`
   - Using [io/ioutil](https://golang.org/pkg/io/ioutil/) to **ReadFile** and [Strings](https://golang.org/pkg/strings/) to convert `[]byte` to `[]string` at [Line 29](https://github.com/kodelint/programming-quest/blob/66ad9d5117c3fd3244ded9563635d8b807ea8e8c/golang/count-uniq.go#L29) and Replace punctuations for `golang`
    
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [count-uniq.sh](bash/count-uniq.sh) |
     | `python`      | [count-uniq.py](python/count-uniq.py) |
     | `golang`      | [count-uniq.go](golang/count-uniq.go) |

            >> sh count-uniq.sh test-data.txt 3
            17 the
            14 data
            10 we
            9 of
            7 to

            >> python count-uniq.py
            the: 17
            data: 14
            we: 10
            of: 9
            to: 7

            >> go run count-uniq.go data.txt
            getting => 1
            production => 1
            me => 3
            online_inventory => 1
            store_bops_eligibility => 1
            We => 3
            loads => 1
            assumed => 1
            job => 2

* **Question 3**: Write a program to print **fizzbuzz** if **number** is divisible by `15`, prints **fizz** divisible by `3` and **buzz** if divisible by `15`  
**Answer**: 

    
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [fizzbuzz.sh](bash/fizzbuzz.sh) |
     | `python`      | [fizzbuzz.py](python/fizzbuzz.py) |

            >> sh bash/fizzbuzz.sh 15
            Number: 3 is fizz
            Number: 5 is buzz
            Number: 6 is fizz
            Number: 9 is fizz
            Number: 10 is buzz
            Number: 12 is fizz
            Number: 15 is fizzbuzz
            
            >> python python/fizzbuzz.py
            number:  1 is SKIPPED
            number:  2 is SKIPPED
            number:  3 is fizz
            number:  4 is SKIPPED
            number:  5 is buzz
            number:  6 is fizz
            number:  7 is SKIPPED
            number:  8 is SKIPPED
            number:  9 is fizz
            number:  10 is buzz
            number:  11 is SKIPPED
            number:  12 is fizz
            number:  13 is SKIPPED
            number:  14 is SKIPPED
            number:  15 is fizzbuzz
   - Using [Recursion](https://bash.cyberciti.biz/guide/Recursive_function)
   
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [fizzbuzz-recursive.sh](bash/fizzbuzz-recursive.sh)| 

            >> sh fizzbuzz-recursive.sh 15
            Number: 15 is fizzbuzz
            Number: 12 is fizz
            Number: 10 is buzz
            Number: 9 is fizz
            Number: 6 is fizz
            Number: 5 is buzz
            Number: 3 is fizz
