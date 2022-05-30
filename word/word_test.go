package word

import "testing"

func TestWord(t *testing.T) {
	words, err := GetWords()
	if err != nil {
		t.Error(err)
	}

	if len(words) == 0 {
		t.Error("word.GetWords() failed")
	}
}
