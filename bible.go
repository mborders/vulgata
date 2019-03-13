package vulgata

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/derekparker/trie"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

// Bible contains the old testament and new testament
type Bible struct {
	OldTestament Testament
	NewTestament Testament
	searchTree   *trie.Trie
}

// SearchNode represents a search result for a given verse inquiry
type SearchNode struct {
	Book    *Book
	Chapter *Chapter
	Verse   *Verse
}

// String creates a string representation for the SearchNode, ex. Genesis 1:1
func (s *SearchNode) String() string {
	return fmt.Sprintf("%s %d:%d",
		s.Book.Title,
		s.Chapter.ChapterNumber,
		s.Verse.VerseNumber)
}

const bibleTar = "bible.tar.gz"
const oldTestamentFilename = "old_testament.json"
const newTestamentFilename = "new_testament.json"

// NewBible creates a new bible instance
func NewBible() *Bible {
	_, currFile, _, _ := runtime.Caller(0)
	filename := fmt.Sprintf("%s/%s", path.Dir(currFile), bibleTar)

	bible := &Bible{
		searchTree: trie.New(),
	}

	f, _ := os.Open(filename)
	defer f.Close()

	gzf, _ := gzip.NewReader(f)
	defer gzf.Close()

	tr := tar.NewReader(gzf)

	// Build testaments, books, chapters, verses
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}

		switch h.Name {
		case oldTestamentFilename:
			bible.OldTestament = Testament{Books: decode(tr)}
		case newTestamentFilename:
			bible.NewTestament = Testament{Books: decode(tr)}
		}
	}

	// Build search tree
	var books []Book
	books = append(books, bible.OldTestament.Books...)
	books = append(books, bible.NewTestament.Books...)

	for i := range books {
		bk := books[i]

		for j := range bk.Chapters {
			ch := bk.Chapters[j]

			for k := range ch.Verses {
				vr := strings.ToLower(ch.Verses[k].Text)
				bible.searchTree.Add(vr, SearchNode{
					Book:    &bk,
					Chapter: &ch,
					Verse:   &ch.Verses[k],
				})
			}
		}
	}

	return bible
}

func decode(r io.Reader) []Book {
	var books []Book
	json.NewDecoder(r).Decode(&books)
	return books
}

// Search finds top matching verses based on the given query.
// The number of search results are restricted by maxResults
func (b *Bible) Search(query string, maxResults int) []SearchNode {
	t := b.searchTree
	keys := t.FuzzySearch(strings.ToLower(query))
	var verses []SearchNode

	for k := range keys {
		res, _ := t.Find(keys[k])
		verses = append(verses, res.Meta().(SearchNode))

		if len(verses) >= maxResults {
			break
		}
	}

	return verses
}
