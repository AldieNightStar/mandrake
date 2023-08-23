package main

import (
	"fmt"

	"github.com/AldieNightStar/mandrake"
)

func main() {
	PASSWORD := []byte("PASSWORD")
	DAT := []byte("This is the best solution")

	str := mandrake.EncodeBase64(DAT, PASSWORD)

	fmt.Println("B64 Encoded:\n================")
	fmt.Println(str)

	dat, err := mandrake.DecodeBase64(str, PASSWORD)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("B64 Decoded:\n================")
	fmt.Println(string(dat))
}
