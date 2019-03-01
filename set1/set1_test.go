package main

import (
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := hexToBase64(input)
	if output != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal(output)
	}
}

func TestXor(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"

	aBytes := hexToBytes(a)
	bBytes := hexToBytes(b)
	cBytes := xor(aBytes, bBytes)

	c := bytesToHex(cBytes)
	if c != "746865206b696420646f6e277420706c6179" {
		t.Fatal(c)
	}
}

func TestFreq(t *testing.T) {
	text := loadFile("./book.txt")
	f := freq(text)
	for char, count := range f {
		fmt.Printf("%s: %d\n", string(char), count)
	}
}

func TestScoreMap(t *testing.T) {
	text := loadFile("./book.txt")
	f := freq(text)
	for char, count := range f {
		fmt.Printf("%s: %d\n", string(char), count)
	}
	getScoreMap(f)
}
