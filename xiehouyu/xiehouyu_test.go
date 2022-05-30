package xiehouyu

import "testing"

func TestXieHouYu(t *testing.T) {
	xiehouyus, err := GetXieHouYus()
	if err != nil {
		t.Error(err)
	}

	if len(xiehouyus) == 0 {
		t.Error("xiehouyu.GetXiehouyus() failed")
	}
}
