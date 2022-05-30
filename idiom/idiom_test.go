package idiom

import "testing"

func TestIdiom(t *testing.T) {
	idioms, err := GetIdioms()
	if err != nil {
		t.Error(err)
	}

	if len(idioms) == 0 {
		t.Error("idiom.GetIdioms() failed")
	}
}
