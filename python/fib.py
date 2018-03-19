def fib(number):
  """ This functions takes a number and produces the Fibonacci sequence """
  a,b = 0, 1
  print "Fibonacci Sequence from 0 to", number ## For formatting
  print "==============================="
  for i in range(0,number):
    print a,
    a,b = b, a + b

fib(10)
