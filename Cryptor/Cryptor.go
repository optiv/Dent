package Cryptor

import (
	"math/rand"
	crand "math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const HEX = "ABCDEF1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[crand.Intn(len(letters))]

	}
	return string(b)
}

func RandCLSID() string {
	nn := 12
	e := make([]byte, nn)
	for i := range e {
		e[i] = HEX[rand.Intn(len(HEX))]

	}

	b := "{FFFDC614-B694-4AE6-AB38-" + string(e) + "}"
	return string(b)
}

func VarNumberLength(min, max int) string {
	var r string
	crand.Seed(time.Now().UnixNano())
	num := crand.Intn(max-min) + min
	n := num
	r = RandStringBytes(n)
	return r
}

func GenerateNumer(min, max int) int {
	crand.Seed(time.Now().UnixNano())
	num := crand.Intn(max-min) + min
	n := num
	return n

}
