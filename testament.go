package vulgata

import "errors"

// Testament consists of a list of books, old vs. new
type Testament struct {
	Books []Book
}

// GetBookNames creates a list of all book titles for the testament
func (t *Testament) GetBookNames() []string {
	var n []string

	for i := range t.Books {
		n = append(n, t.Books[i].Title)
	}

	return n
}

// GetBook obtains a book within the testament by its number
func (t *Testament) GetBook(bookNumber int) (*Book, error) {
	idx := bookNumber - 1
	if idx < 0 || idx >= len(t.Books) {
		return nil, errors.New("invalid book number")
	}

	return &t.Books[idx], nil
}

// GetChapter obtains a chapter within the provided book by its number
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

// GetVerse obtains a verse within the provided book/chapter by its number
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
