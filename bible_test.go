package vulgata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestament_GetBookNames(t *testing.T) {
	b := NewBible()

	names := b.OldTestament.GetBookNames()
	assert.Equal(t, 46, len(names))
}

func TestTestament_GetBook(t *testing.T) {
	b := NewBible()

	book, err := b.OldTestament.GetBook(1)
	assert.Nil(t, err)
	assert.Equal(t, "Genesis", book.Title)
}

func TestTestament_GetBook_New(t *testing.T) {
	b := NewBible()

	book, err := b.NewTestament.GetBook(1)
	assert.Nil(t, err)
	assert.Equal(t, "Matthew", book.Title)
}

func TestTestament_GetBook_Invalid(t *testing.T) {
	b := NewBible()

	_, err := b.OldTestament.GetBook(0)
	assert.Equal(t, "invalid book number", err.Error())

	_, err = b.OldTestament.GetBook(100)
	assert.Equal(t, "invalid book number", err.Error())
}

func TestTestament_GetChapter(t *testing.T) {
	b := NewBible()

	chap, err := b.OldTestament.GetChapter(1, 1)
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), chap.ChapterNumber)
}

func TestTestament_GetChapter_InvalidBook(t *testing.T) {
	b := NewBible()

	_, err := b.OldTestament.GetChapter(0, 0)
	assert.Equal(t, "invalid book number", err.Error())

	_, err = b.OldTestament.GetChapter(0, 100)
	assert.Equal(t, "invalid book number", err.Error())
}

func TestTestament_GetChapter_InvalidChapter(t *testing.T) {
	b := NewBible()

	_, err := b.OldTestament.GetChapter(1, 0)
	assert.Equal(t, "invalid chapter number", err.Error())

	_, err = b.OldTestament.GetChapter(1, 100)
	assert.Equal(t, "invalid chapter number", err.Error())
}

func TestTestament_GetVerse(t *testing.T) {
	b := NewBible()

	vers, err := b.OldTestament.GetVerse(1, 1, 4)
	assert.Nil(t, err)
	assert.Equal(t, "And God saw the light that it was good; and he divided the light from the darkness.", vers.Text)
}

func TestTestament_GetVerse_New(t *testing.T) {
	b := NewBible()

	vers, err := b.NewTestament.GetVerse(4, 8, 32)
	assert.Nil(t, err)
	assert.Equal(t, "And you shall know the truth, and the truth shall make you free.", vers.Text)
}

func TestTestament_GetVerse_InvalidChapter(t *testing.T) {
	b := NewBible()

	_, err := b.OldTestament.GetVerse(1, 0, 0)
	assert.Equal(t, "invalid chapter number", err.Error())

	_, err = b.OldTestament.GetVerse(1, 100, 1000)
	assert.Equal(t, "invalid chapter number", err.Error())
}

func TestTestament_GetVerse_InvalidVerse(t *testing.T) {
	b := NewBible()

	_, err := b.OldTestament.GetVerse(1, 1, 0)
	assert.Equal(t, "invalid verse number", err.Error())

	_, err = b.OldTestament.GetVerse(1, 1, 1000)
	assert.Equal(t, "invalid verse number", err.Error())
}
