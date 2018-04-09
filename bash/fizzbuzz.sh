fizz () {
	if [ $((num % 15)) == 0 ]; then
	      echo "Number: $num is fizzbuzz"
      elif [ $((num % 5)) == 0 ]; then
	      echo "Number: $num is buzz"
      elif [ $((num % 3)) == 0 ]; then
	      echo "Number: $num is fizz"
	fi
}

for num in $(seq 1 $1)
do
  fizz $num
done
