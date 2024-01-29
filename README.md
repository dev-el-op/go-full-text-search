# Go Full Text Search

## Introduction

This is a simple full text search engine written in Go.

## Sample dataset

You can download the dataset i used for testing from: [https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz)

## Usage

```go
go run main.go -p <path to dataset file> -q <query>
```

Where `<path to dataset file>` is the path to the dataset file and `<query>` is the query you want to search for.

In default the dataset file is `enwiki-latest-abstract1.xml.gz` and the query is `Barack Obama`.
