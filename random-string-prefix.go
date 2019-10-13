package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	length           = 6
	haveLeadingDigit = false
)

func getRandomInt(max int) (num int64) {
	// this is from 0 to max-1
	b, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return b.Int64()
}

func sample(set []string) (item string) {
	return set[getRandomInt(len(set))]
}

func main() {
	letters := []string{
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m",
		"n", "p", "q", "r", "s", "t", "v", "w", "x", "z",
	}

	digits := []string{
		"0", "1", "2", "3", "5", "6", "7", "9",
	}

	all := append(letters, digits...)

	var buf bytes.Buffer

	if haveLeadingDigit {
		for range make([]int, length) {
			buf.WriteString(sample(all))
		}
	} else {
		buf.WriteString(sample(letters))
		for range make([]int, length-1) {
			buf.WriteString(sample(all))
		}
	}

	fmt.Println(buf.String())
}
