fizz () {
	if [[ $1 -lt 1 ]]; then
		exit 0
	fi
	if [[ $num -lt 1 ]]; then
		exit 0
	fi
	if [ $((num % 15)) == 0 ]; then
	      echo "Number: $num is fizzbuzz"
        elif [ $((num % 5)) == 0 ]; then
  	      echo "Number: $num is buzz"
        elif [ $((num % 3)) == 0 ]; then
	      echo "Number: $num is fizz"
	fi
let num=$num-1
fizz $num
}

num=$1
fizz $num
