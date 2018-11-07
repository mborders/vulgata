[![GoDoc](http://godoc.org/github.com/borderstech/vulgata?status.png)](http://godoc.org/github.com/borderstech/vulgata)
[![Build Status](https://travis-ci.org/borderstech/vulgata.svg?branch=master)](https://travis-ci.org/borderstech/vulgata)
[![Go Report Card](https://goreportcard.com/badge/github.com/borderstech/vulgata)](https://goreportcard.com/report/github.com/borderstech/vulgata)
[![codecov](https://codecov.io/gh/borderstech/vulgata/branch/master/graph/badge.svg)](https://codecov.io/gh/borderstech/vulgata)

# vulgata

Golang library containing the entire Holy Bible with both the Douay-Rheims English and Clementina Vulgata Latin texts.

Documentation here: https://godoc.org/github.com/borderstech/vulgata

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
```
