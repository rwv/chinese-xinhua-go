package idiom

import (
	"embed"
	"encoding/json"
)

//go:embed idiom.json
var f embed.FS

func GetIdioms() ([]Idiom, error) {
	var idioms = []Idiom{}

	file, err := f.Open("idiom.json")
	if err != nil {
		return idioms, err
	}

	err = json.NewDecoder(file).Decode(&idioms)
	if err != nil {
		return idioms, err
	}

	return idioms, nil
}
