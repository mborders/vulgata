[![GoDoc](http://godoc.org/github.com/mborders/vulgata?status.png)](http://godoc.org/github.com/mborders/vulgata)
[![Build Status](https://travis-ci.org/mborders/vulgata.svg?branch=master)](https://travis-ci.org/mborders/vulgata)
[![Go Report Card](https://goreportcard.com/badge/github.com/mborders/vulgata)](https://goreportcard.com/report/github.com/mborders/vulgata)
[![codecov](https://codecov.io/gh/mborders/vulgata/branch/master/graph/badge.svg)](https://codecov.io/gh/mborders/vulgata)

# vulgata

Golang library containing the entire Holy Bible with both the Douay-Rheims English and Clementina Vulgata Latin texts.

Documentation here: https://godoc.org/github.com/mborders/vulgata

## Example Usage

```go
// Create a new bible instance
b := vulgata.NewBible()

// Get all book names for the old testament
names := b.OldTestament.GetBookNames()

// Get book 1 from the old testament
book, err := b.OldTestament.GetBook(1)
fmt.Print(book.Title) // title of book

// Get book 1, chapter 1 from the old testament
chap, err := b.OldTestament.GetChapter(1, 1)

// Get book 4, chapter 8, verse 32 from the new testament
vers, err := b.NewTestament.GetVerse(4, 8, 32)
fmt.Print(vers.Text) // English
fmt.Print(vers.TextLatin) // Latin

// Search for verses
v := b.Search("truth make you free", 10) // max of 10 results
fmt.Print(v[0].Book.Title)
fmt.Print(v[0].Chapter.ChapterNumber)
fmt.Print(v[0].Verse.Text)
fmt.Print(v[0].String()) // John 8:32
```
