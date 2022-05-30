package word

import (
	"embed"
	"encoding/json"
)

//go:embed word.json
var f embed.FS

func GetWords() ([]Word, error) {
	var words = []Word{}

	file, err := f.Open("word.json")
	if err != nil {
		return words, err
	}

	err = json.NewDecoder(file).Decode(&words)
	if err != nil {
		return words, err
	}

	return words, nil
}
