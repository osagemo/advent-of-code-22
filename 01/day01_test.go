package main

import (
	"fmt"
	"testing"
)

const input1 = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestDay1Part1(t *testing.T) {
	result := Part1(input1)
	expected := 24000

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	result := Part2(input1)
	expected := 45000

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
