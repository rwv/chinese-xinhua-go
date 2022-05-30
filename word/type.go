package word

type Word struct {
	Word        string `json:"word"`
	OldWord     string `json:"oldword"`
	Strokes     string `json:"strokes"`
	Pinyin      string `json:"pinyin"`
	Radicals    string `json:"Radicals"`
	Explanation string `json:"explanation"`
	More        string `json:"more"`
}
