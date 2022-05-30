package ci

import (
	"embed"
	"encoding/json"
)

//go:embed ci.json
var f embed.FS

func GetCis() ([]Ci, error) {
	var cis = []Ci{}

	file, err := f.Open("ci.json")
	if err != nil {
		return cis, err
	}

	err = json.NewDecoder(file).Decode(&cis)
	if err != nil {
		return cis, err
	}

	return cis, nil
}
