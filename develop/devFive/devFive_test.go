package main

import (
	"reflect"
	"testing"

	"exeTwo.devThree/functions"
)

func TestSortByAKey(t *testing.T) {
	var check = make(map[int][]string)
	check[1] = []string{"Come hither, come hither, come hither Here shall he see"}

	for _, correctResult := range check {
		letTest := functions.SortByAKey(functions.ReadFromFile("text.txt"), "Unto the sweet bird's throat", 2)

		if reflect.DeepEqual(letTest, correctResult) {
			t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, correctResult)
		}
	}
}

func TestSortBycKey(t *testing.T) {
	var check = 8

	letTest := functions.SortBycKey(functions.ReadFromFile("text.txt"))

	if letTest != check {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}

func TestSortByFKey(t *testing.T) {
	var check = "And tune his merry note"

	letTest := functions.SortByFKey(functions.ReadFromFile("text.txt"), "And tune his merry note")

	if letTest != check {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}

func TestSortByiKey(t *testing.T) {
	var check = "And tune his merry note"

	letTest := functions.SortByiKey(functions.ReadFromFile("text.txt"), "And TUNE his merry NOTE")

	if letTest != check {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}

func TestSortByvKey(t *testing.T) {
	var check []string
	check = append(check, "Under the greenwood tree And tune his merry note Unto the sweet bird's throat Come hither, come hither, come hither Here shall he see No enemy But winter and rough weather.")

	letTest := functions.SortByvKey(functions.ReadFromFile("text.txt"), "And TUNE his merry NOTE")

	if reflect.DeepEqual(letTest, check) {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}

func TestSortBynKey(t *testing.T) {

	check := 3
	letTest := functions.SortBynKey(functions.ReadFromFile("text.txt"), "And tune his merry note")

	if check != letTest {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}

func TestSortByCKey(t *testing.T) {

	var check []string

	check = append(check, "Who loves to lie with me, And tune his merry note Come hither, come hither, come hither Here shall he see")

	letTest := functions.SortByCKey(functions.ReadFromFile("text.txt"), "Unto the sweet bird's throat", 2)

	if reflect.DeepEqual(check, letTest) {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}
