from collections import Counter
import string

def word_count(fname):
        with open(fname) as f:
            tokens = f.read().split()

        count = Counter([ token.translate(None, string.punctuation).lower() for token in tokens ])
        for word,occurance in sorted(count.iteritems(), key=lambda (k,v): (v,k), reverse=True):
            print "%s: %s" % (word,occurance)

print(word_count("data.txt"))
