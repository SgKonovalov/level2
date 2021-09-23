package base

import (
	"reflect"
	"testing"

	base "exeTwo.devFour/base"
)

func TestGetTheResult(t *testing.T) {

	check := map[string][]string{
		"тяпка":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	letTest, err := base.GetTheResult(true, true)
	if err != nil {
		t.Fatalf("Error %v occurated by parsing %s string. Correct is %v.", err, letTest, check)
	}
	if !reflect.DeepEqual(letTest, check) {
		t.Fatalf("Incorrect parsing: %v and %v", check, letTest)
	}
}

func TestGetTheResultIfOnlyOneResult(t *testing.T) {

	check := map[string][]string{
		"тяпка": {"пятак", "пятка", "тяпка"},
	}

	letTest, err := base.GetTheResult(true, false)
	if err != nil {
		t.Fatalf("Error %v occurated by parsing %s string. Correct is %v.", err, letTest, check)
	}
	if !reflect.DeepEqual(letTest, check) {
		t.Fatalf("Incorrect parsing: %v and %v", check, letTest)
	}
}

func TestGetTheResultIfNoResults(t *testing.T) {

	check := map[string][]string{}

	letTest, err := base.GetTheResult(false, false)
	if err != nil {
		t.Fatalf("Error %v occurated by parsing %s string. Correct is %v.", err, letTest, check)
	}
	if !reflect.DeepEqual(letTest, check) {
		t.Fatalf("Incorrect parsing: %v and %v", check, letTest)
	}
}
