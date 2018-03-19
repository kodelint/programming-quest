number = 30
def fib(number):
    if number == 0:
        return 1
    elif number == 1:
        return 1
    elif number == 2:
        return 2
    elif number > 2:
        return fib(number - 1) + fib(number - 2)


print("Fibonacci Sequence from 0 to", number) ## For formatting
print("======================================")
for n in range(0,number):
    print(fib(n), end=', ')
