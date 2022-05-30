package main

import (
	"fmt"

	"github.com/rwv/chinese-xinhua-go/ci"
	"github.com/rwv/chinese-xinhua-go/idiom"
	"github.com/rwv/chinese-xinhua-go/word"
	"github.com/rwv/chinese-xinhua-go/xiehouyu"
)

func main() {
	idioms, err := idiom.GetIdioms()
	if err == nil {
		fmt.Printf("成语: %d\n", len(idioms))
	}

	words, err := word.GetWords()
	if err == nil {
		fmt.Printf("汉字: %d\n", len(words))
	}

	cis, err := ci.GetCis()
	if err == nil {
		fmt.Printf("词语: %d\n", len(cis))
	}

	xiehouyus, err := xiehouyu.GetXieHouYus()
	if err == nil {
		fmt.Printf("歇后语: %d\n", len(xiehouyus))
	}
}
