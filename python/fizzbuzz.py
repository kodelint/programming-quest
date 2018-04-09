def fizzbuzz(seq):
    for number in range(1, seq):
        if number % 15 == 0:
            print "number: ", number, "is fizzbuzz"
        elif number % 3 == 0:
            print "number: ", number, "is fizz"
        elif number % 5 == 0:
            print "number: ", number, "is buzz"

fizzbuzz(16)
