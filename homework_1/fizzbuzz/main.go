package main

import (
	fizzbuzz "code-cadests-2021/homework_1/fizzbuzz/fizzbuzz_test"
	"flag"
	"log"
)


func main() {
	start := flag.Int("start", 1, "Start value for Fizz Buzz game")
	end := flag.Int("end", 10, "End value for Fizz Buzz game")
	flag.Parse()
	fizzbuzzslice, err := fizzbuzz.Fizzbuzzgame(*start , *end)
	if err != nil {
		log.Fatal(err)
	}
	fizzbuzz.Printgameresults(fizzbuzzslice)

}
