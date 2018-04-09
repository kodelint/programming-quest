import sys

def fizzbuzz(seq):
    for number in range(1, seq):
        if number % 15 == 0:
            print "number: ", number, "is fizzbuzz"
        elif number % 3 == 0:
            print "number: ", number, "is fizz"
        elif number % 5 == 0:
            print "number: ", number, "is buzz"

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print "Please provide a number"
        sys.exit(1)

    fizzbuzz(int(sys.argv[1]))
