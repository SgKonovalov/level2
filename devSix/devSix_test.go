package main

import (
	"reflect"
	"testing"

	"exeTwo.devThree/functions"
)

func TestSortByfKey(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{" Two   Feb   Tue "}

	for _, correctResult := range check {
		letTest := functions.SortByfKey(functions.ReadFromFile("text.txt"), 2)

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}

func TestSortBysKey(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"OneO Two Three Jan Feb Mar Monday Tue Wen"}

	for _, correctResult := range check {
		letTest := functions.SortBysKey(functions.ReadFromFile("text.txt"))

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}

func TestSortBydKey(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"OneO&Two&ThreeHereNoSeparatorJan&Feb&MarMonday&Tue&WenHereToo"}

	for _, correctResult := range check {
		letTest := functions.SortBydKey(functions.ReadFromFile("text.txt"), "&")

		if !reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}
