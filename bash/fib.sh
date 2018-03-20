#!/bin/bash

# Found the example from following link https://stackoverflow.com/questions/42987607/shell-script-to-generate-fibonacci-series

fib() { 
    local nr rez rez1 rez2
    if [ $1 -eq 1 -o $1 -eq 2 ]; then
        echo 1
    else
        let nr=$1-1
        rez1=$(fib $nr)
        let nr=$1-2
        rez2=$(fib $nr)
        let rez=$rez1+$rez2
        echo $rez
    fi
}

for n in $(seq 1 $1)
do
 fib $n
done
