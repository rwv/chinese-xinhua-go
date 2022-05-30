package idiom

type Idiom struct {
	Derivation   string `json:"derivation"`
	Example      string `json:"example"`
	Explanation  string `json:"explanation"`
	Pinyin       string `json:"pinyin"`
	Word         string `json:"word"`
	Abbreviation string `json:"abbreviation"`
}
