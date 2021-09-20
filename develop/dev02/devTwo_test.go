package main

import "testing"

func TestTranslateString(t *testing.T) {
	check := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"qwe\\5":   "qwe\\\\\\\\\\",
	}

	for testString, correctResult := range check {
		letTest, err := translateString(testString)
		if err != nil {
			t.Fatalf("Error %v occurated by parsing %s string. Correct is %v.", err, letTest, correctResult)
		}
		if letTest != correctResult {
			t.Fatalf("Incorrect parsing string %s. Result: %v, expected: %v", testString, letTest, correctResult)
		}
	}
}

func TestWrongStringsTranslateString(t *testing.T) {
	incorrect := "45"

	letTest, err := translateString(incorrect)
	if letTest != "" {
		t.Fatalf("Wrong parsing string %s, NOT EMPTY STRING. Returning %v", incorrect, letTest)
	}
	if err == nil {
		t.Fatalf("Wrong parsing string %s, expected ERROR", incorrect)
	}

}
