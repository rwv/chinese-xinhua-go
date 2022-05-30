package xiehouyu

import (
	"embed"
	"encoding/json"
)

//go:embed xiehouyu.json
var f embed.FS

func GetXieHouYus() ([]XieHouYu, error) {
	var xiehouyus = []XieHouYu{}

	file, err := f.Open("xiehouyu.json")
	if err != nil {
		return xiehouyus, err
	}

	err = json.NewDecoder(file).Decode(&xiehouyus)
	if err != nil {
		return xiehouyus, err
	}

	return xiehouyus, nil
}
