package ci

import "testing"

func TestCi(t *testing.T) {
	cis, err := GetCis()
	if err != nil {
		t.Error(err)
	}

	if len(cis) == 0 {
		t.Error("idiom.GetCis() failed")
	}
}
