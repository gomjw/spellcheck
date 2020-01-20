package spellcheck

import (
	"fmt"
	"testing"

)

var (
	shouldPass = []string{"Hello", "What", "Are", "yoU", "dOinG"}
	shouldFail = []string{"aosijf", "jdaaiUPHAiusf", "asdh72"}
)

func TestCheck(t *testing.T) {

	fmt.Println("\n--- Testing correct words ---")
	for _, word := range shouldPass {
		if CheckWord(word) {
			fmt.Println("\"", word, "\"", "passed")
		} else {
			t.Errorf("%s failed\n", word)
		}
	}

	fmt.Println("\n--- Testing incorrect words ---")
	for _, word := range shouldFail {
		if !CheckWord(word) {
			fmt.Println("\"", word, "\"", "passed")
		} else {
			t.Errorf("%s failed\n", word)
		}
	}

}
