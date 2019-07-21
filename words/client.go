package words

import (
	"bufio"
	"errors"
	"io"
)

// Client is words package core struct, define & containsall the needed
// settings to perform Word operations.
//
// This "Client" idea came to my mind from the standard library http.Client
// struct. http.Client is a struct that provides a single entrypoint for the
// main operations on the http package. I'm used to that pattern for my packages
// and words.Client will provide all word's features needed for the Kata.
type Client struct {
	Words []Word
}

// NewClient loads a feed of words, an initialize them.
//
// words: An io.Reader with a words dictionary, word per line.
//
// If there is no words on the provided words io.Reader it will return an error
//
// Returns a reference to the created Client or an error if any.
func NewClient(words io.Reader) (*Client, error) {
	client := &Client{
		Words: []Word{},
	}

	scanner := bufio.NewScanner(words)
	for scanner.Scan() {
		client.Words = append(client.Words, Word{scanner.Text()})
	}

	if len(client.Words) == 0 {
		return nil, errors.New("error, no words found on the given input")
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return client, nil
}
