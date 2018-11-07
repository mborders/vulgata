package vulgata

// Book represents a single book in a testament of the bible
type Book struct {
	Title      string    `json:"title"`
	BookNumber uint8     `json:"bookNumber"`
	Chapters   []Chapter `json:"Chapters"`
}

// Chapter represents a single chapter in a book of the bible
type Chapter struct {
	ChapterNumber uint8   `json:"chapterNumber"`
	Verses        []Verse `json:"verses"`
}

// Verse represents a single verse in a chapter of the bible,
// with both English and Latin texts
type Verse struct {
	VerseNumber uint8  `json:"verseNumber"`
	Text        string `json:"text"`
	TextLatin   string `json:"textLatin"`
}
