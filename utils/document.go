package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]Document, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	gzip, error := gzip.NewReader(file)

	if error != nil {
		return nil, error
	}

	defer gzip.Close()

	decoder := xml.NewDecoder(gzip)

	dump := struct {
		Documents []Document `xml:"doc"`
	}{}

	if err := decoder.Decode(&dump); err != nil {
		return nil, err
	}

	documents := dump.Documents

	for i := range documents {
		documents[i].ID = i
	}

	return documents, nil
}
