import sys

def fizzbuzz(number):
    if number < 3:
        sys.exit(0)

    if number % 15 == 0:
        print "number: ", number, "is fizzbuzz"
    elif number % 3 == 0:
        print "number: ", number, "is fizz"
    elif number % 5 == 0:
        print "number: ", number, "is buzz"
    fizzbuzz(number - 1)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print "Please provide a number"
        sys.exit(1)

    fizzbuzz(int(sys.argv[1]))
