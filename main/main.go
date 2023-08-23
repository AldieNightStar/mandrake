package main

import (
	"fmt"

	"github.com/AldieNightStar/mandrake"
)

func main() {
	dat := mandrake.Encode(
		[]byte("Hello"),
		[]byte("pw888"),
	)

	fmt.Println("WAS:\n======================")
	fmt.Println(dat)

	dat = mandrake.Decode(dat, []byte("pw879"))

	fmt.Println("BECAME:\n======================")
	fmt.Println(string(dat))

}
