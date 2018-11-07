package vulgata

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

type Bible struct {
	OldTestament Testament
	NewTestament Testament
}

func NewBible() *Bible {
	f, _ := os.Open("bible.tar.gz")
	defer f.Close()

	gzf, _ := gzip.NewReader(f)
	tarReader := tar.NewReader(gzf)

	bible := &Bible{}

	for {
		h, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		buf := bytes.Buffer{}
		io.Copy(&buf, tarReader)

		var books []Book
		json.NewDecoder(&buf).Decode(&books)

		switch h.Name {
		case "old_testament.json":
			bible.OldTestament = Testament{Books: books}
		case "new_testament.json":
			bible.NewTestament = Testament{Books: books}
		}
	}

	return bible
}
