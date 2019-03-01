package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {

}

func hexToBase64(h string) string {
	bytes, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func hexToBytes(h string) []byte {
	bytes, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	return bytes
}

func bytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("the length of the buffers is not the same")
	}
	c := make([]byte, len(a))
	for index := range a {
		c[index] = a[index] ^ b[index]
	}
	return c
}

func loadFile(fn string) string {
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// returns the frequency of characters
// in an string
func freq(txt string) map[rune]int {
	m := map[rune]int{}
	for _, c := range txt {
		if unicode.IsLetter(c) {
			if _, ok := m[c]; ok {
				m[c]++
			} else {
				m[c] = 1
			}
		}
	}
	return m
}

// given a frequency map returns a weigth map
// of characters. Normalizes the values.
func getScoreMap(freq map[rune]int) map[rune]float64 {
	// get total
	var total int
	for _, count := range freq {
		total += count
	}
	freqFloat64 := make(map[rune]float64, len(freq))
	for char := range freq {
		freqFloat64[char] = float64(freq[char]) / float64(total)
		fmt.Printf("%s: %f.2%%\n", string(char), freqFloat64[char])
	}
	return freqFloat64
}
