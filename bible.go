package vulgata

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

// Bible contains the old testament and new testament
type Bible struct {
	OldTestament Testament
	NewTestament Testament
}

const bibleTar = "./bible.tar.gz"
const oldTestamentFilename = "old_testament.json"
const newTestamentFilename = "new_testament.json"

// NewBible creates a new bible instance
func NewBible() *Bible {
	bible := &Bible{}

	f, _ := os.Open(bibleTar)
	defer f.Close()

	gzf, _ := gzip.NewReader(f)
	defer gzf.Close()

	tr := tar.NewReader(gzf)

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

	return bible
}

func decode(r io.Reader) []Book {
	var books []Book
	json.NewDecoder(r).Decode(&books)
	return books
}
