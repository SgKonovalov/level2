package main

import (
	"reflect"
	"testing"

	"exeTwo.devThree/functions"
)

func TestSortMain(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"6. Iota June 3. Gamma March 5. Epsilon May 4. Delta April 2. Beta  February 1. Alfa January"}

	for _, correctResult := range check {
		letTest := functions.SortMain(functions.ReadFromFile("text.txt"), false, true, false, false, false, false, 2)

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}

func TestSortByM(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"6. Iota June 5. Epsilon May 4. Delta April 3. Gamma March 2. Beta  February 1. Alfa January"}

	for _, correctResult := range check {
		letTest := functions.SortByM(functions.ReadFromFile("text.txt"), false, true, false, false)

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}

func TestSortByu(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"6. Iota June 5. Epsilon May 3. Gamma March 1. Alfa January 4. Delta April 2. Beta  February"}

	for _, correctResult := range check {
		letTest := functions.SortMain(functions.ReadFromFile("text.txt"), false, false, false, true, false, false, 2)

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}
