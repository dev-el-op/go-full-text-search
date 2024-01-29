package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/dev-el-op/go-full-text-search/utils"
)

func main() {
	var dumpPath, query string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abctract dump path")
	flag.StringVar(&query, "q", "Barack Obama", "search query")
	flag.Parse()
	log.Println("Full text search is ready to use")

	startTime := time.Now()
	docs, err := utils.LoadDocuments(dumpPath)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded %d documents in %s", len(docs), time.Since(startTime))

	startTime = time.Now()

	index := make(utils.Index)
	index.Add(docs)
	log.Printf("Indexed %d documents in %s", len(index), time.Since(startTime))

	startTime = time.Now()

	matchedIds := index.Search(query)
	log.Printf("Found %d documents in %s", len(matchedIds), time.Since(startTime))

	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
