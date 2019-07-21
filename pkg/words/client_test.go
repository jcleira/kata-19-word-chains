package words

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		Description    string
		WordsReader    io.Reader
		ExpectedError  error
		ExpectedClient *Client
	}{
		{
			Description: "when NewClient succesfully creates a client",
			WordsReader: strings.NewReader("foo\nbar\nfoobar\nbarfoo"),
			ExpectedClient: &Client{
				Words: []Word{
					{Term: "foo"},
					{Term: "bar"},
					{Term: "foobar"},
					{Term: "barfoo"},
				},
			},
		},
		{
			Description:   "when NewClient fails due empty words input",
			WordsReader:   strings.NewReader(""),
			ExpectedError: errors.New("error, no words found on the given input"),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			client, err := NewClient(test.WordsReader)

			if !reflect.DeepEqual(err, test.ExpectedError) {
				t.Fatalf("expected err to be %v, got %v", test.ExpectedError, err)
			}

			if !reflect.DeepEqual(client, test.ExpectedClient) {
				t.Fatalf("expected Client to be %v, got %v", test.ExpectedClient, client)
			}
		})
	}
}
