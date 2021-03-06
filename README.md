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

* **Question 3**: Write a program to check if given string is **Palindrom**   
**Answer**: 

   - Using [rune](https://godoc.org/golang.org/x/text/runes) for `golang` 
    
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `python`      | [palindrome.py](python/palindrome.py) |
     | `golang`      | [palindrome.go](golang/palindrome.go) |  

            >> python palindrom.py madam
            madam is palindrom
            >> python palindrom.py foobaar
            foobaar is not palindrom

            >> go run palindrom.go madman
            Given String: madman is not palindrom
            >> go run palindrom.go madam
            Given String: madam is palindrom


* **Question 4**: Write a program to print **fizzbuzz** if **number** is divisible by `15`, prints **fizz** divisible by `3` and **buzz** if divisible by `15`  
**Answer**: 

   - Using [strconv](https://golang.org/pkg/strconv/) for `golang` for `string` to `int` conversion
 
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [fizzbuzz.sh](bash/fizzbuzz.sh) |
     | `python`      | [fizzbuzz.py](python/fizzbuzz.py) | 
     | `golang`      | [fizzbuzz.go](golang/fizzbuzz.go) |



            >> sh bash/fizzbuzz.sh 15
            Number: 3 is fizz
            Number: 5 is buzz
            Number: 6 is fizz
            Number: 9 is fizz
            Number: 10 is buzz
            Number: 12 is fizz
            Number: 15 is fizzbuzz
            
            >> python python/fizzbuzz.py
            number:  3 is fizz
            number:  5 is buzz
            number:  6 is fizz
            number:  9 is fizz
            number:  10 is buzz
            number:  12 is fizz
            number:  15 is fizzbuzz  

            >> go run fizzbuzz.go 15
            Number: 3 is fizz
            Number: 5 is buzz
            Number: 6 is fizz
            Number: 9 is fizz
            Number: 10 is buzz
            Number: 12 is fizz
            Number: 15 is fizzbuzz
            
   - Using [Recursion](https://bash.cyberciti.biz/guide/Recursive_function)
   
     | Language       | Script Name  | 
     | ------------- |:-------------:|
     | `bash`      | [fizzbuzz-recursive.sh](bash/fizzbuzz-recursive.sh)|
     | `python`      | [fizzbuzz-recursive.py](python/fizzbuzz-recursive.py)| 
     | `golang`      | [fizzbuzz-recursive.go](golang/fizzbuzz-recursive.go)|  

            >> sh fizzbuzz-recursive.sh 15
            Number: 15 is fizzbuzz
            Number: 12 is fizz
            Number: 10 is buzz
            Number: 9 is fizz
            Number: 6 is fizz
            Number: 5 is buzz
            Number: 3 is fizz

            >> python fizzbuzz-recursive.py 15
            number:  15 is fizzbuzz
            number:  12 is fizz
            number:  10 is buzz
            number:  9 is fizz
            number:  6 is fizz
            number:  5 is buzz
            number:  3 is fizz

            >> go run fizzbuzz-recursive.go 15
            Number: 15 is fizzbuzz
            Number: 12 is fizz
            Number: 10 is buzz
            Number: 9 is fizz
            Number: 6 is fizz
            Number: 5 is buzz
            Number: 3 is fizz

* **Question 5**: Write a program to compare `Phone Numbers` from `2` different `csv` files and generate records, sample files and expected records example are given below:  

   - Hint [encoding/csv](https://golang.org/pkg/encoding/csv/) for `golang`, [csv](https://docs.python.org/2/library/csv.html) for `python`

            > cat first.csv
            Sam,111222333,77878 Gkaall Way
            Ray,2223334444,77878 toorojd Way
            Will,111222333,77878 Gkaall Gundtom

            > cat last.csv
            Roy,10000,111222333
            Bhal,5000,2223334444
            Nahal,7000,111222333
            Pari,5000,5567778888

            > output
            {Sam Roy 111222333 77878 Gkaall Way 10000}
            {Will Nahal 111222333 77878 Gkaall Gundtom 7000}

**Answer**: 
    
  - Using [encoding/csv](https://golang.org/pkg/encoding/csv/)

      | Language       | Script Name  | 
      | ------------- |:-------------:|
      | `golang`      | [csv-io.go](golang/csv-io.go) |  

            >> go run csv-io.go file1.csv file2.csv
            {Sam Roy 111222333 77878 Gkaall Way 10000}
            {Will Nahal 111222333 77878 Gkaall Gundtom 7000}
