import sys

def is_palindrom(text):
    rev_text = text[::-1]
    if text == rev_text:
        print text + " is palindrom"
    else:
        print text + " is not palindrom"


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print "Please provide a string"
        sys.exit(1)

    is_palindrom(sys.argv[1])
