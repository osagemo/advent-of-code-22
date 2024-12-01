package main

import (
	"fmt"
	"testing"
)

const input1 = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestDay3Part1(t *testing.T) {
	result := Part1(input1)
	expected := 157

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}

func TestDay3Part2(t *testing.T) {
	result := Part2(input1)
	expected := 70

	if result != expected {
		fmt.Printf("got %v, expected %v\n", result, expected)
		t.Fail()
	}
}
