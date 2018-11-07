package vulgata

import "errors"

type Testament struct {
	Books []Book
}

func (t *Testament) GetBookNames() []string {
	var n []string

	for i := range t.Books {
		n = append(n, t.Books[i].Title)
	}

	return n
}

func (t *Testament) GetBook(bookNumber int) (*Book, error) {
	idx := bookNumber - 1
	if idx < 0 || idx >= len(t.Books) {
		return nil, errors.New("invalid book number")
	}

	return &t.Books[idx], nil
}

func (t *Testament) GetChapter(bookNumber int, chapterNumber int) (*Chapter, error) {
	b, err := t.GetBook(bookNumber)
	if err != nil {
		return nil, err
	}

	idx := chapterNumber - 1
	if idx < 0 || idx >= len(b.Chapters) {
		return nil, errors.New("invalid chapter number")
	}

	return &b.Chapters[idx], nil
}

func (t *Testament) GetVerse(bookNumber int, chapterNumber int, verseNumber int) (*Verse, error) {
	c, err := t.GetChapter(bookNumber, chapterNumber)
	if err != nil {
		return nil, err
	}

	idx := verseNumber - 1
	if idx < 0 || idx >= len(c.Verses) {
		return nil, errors.New("invalid verse number")
	}

	return &c.Verses[idx], nil
}
