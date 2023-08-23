package mandrake

import (
	"bytes"
	"hash/fnv"
	"math/rand"
)

func hash(dat []byte) int64 {
	h := fnv.New64()
	h.Write(dat)
	return int64(h.Sum64())
}

func Encode(dat []byte, password []byte) []byte {
	result := bytes.Clone(dat)
	r := rand.NewSource(hash(password))

	// Add random number of bytes
	toAdd := int(r.Int63() % 128)
	for i := 0; i < toAdd; i++ {
		result = append(result, byte(255-i))
	}

	// Add some special number to each byte
	// And of course add each password number
	specAdder := byte(r.Int63() % 250)

	var passwordAdd byte
	var randomByte byte
	for id, val := range result {
		randomByte = byte(r.Int63() % 250)
		passwordAdd = password[id%len(password)]
		result[id] = val + specAdder + passwordAdd + randomByte
	}

	return result
}

func Decode(dat []byte, password []byte) []byte {
	result := bytes.Clone(dat)

	r := rand.NewSource(hash(password))

	// Get how many bytes were added
	wasAdded := int(r.Int63() % 128)

	// Get special adder
	specAdder := byte(r.Int63() % 250)

	// Decryption
	var passwordSub byte
	var randomByte byte
	for id, val := range result {
		randomByte = byte(r.Int63() % 250)
		passwordSub = password[id%len(password)]
		result[id] = ((val - passwordSub) - specAdder) - randomByte
	}

	// Remove last bytes
	resultLen := len(result)
	if resultLen-wasAdded > 0 {
		result = result[0 : resultLen-wasAdded]
	}

	return result
}
