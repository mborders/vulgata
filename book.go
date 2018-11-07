package vulgata

type Book struct {
	Title      string    `json:"title"`
	BookNumber uint8     `json:"bookNumber"`
	Chapters   []Chapter `json:"Chapters"`
}

type Chapter struct {
	ChapterNumber uint8   `json:"chapterNumber"`
	Verses        []Verse `json:"verses"`
}

type Verse struct {
	VerseNumber uint8  `json:"verseNumber"`
	Text        string `json:"text"`
	TextLatin   string `json:"textLatin"`
}
