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
				Words: []*Word{
					{Term: "foo"},
					{Term: "bar"},
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

func TestGetChain(t *testing.T) {
	tests := []struct {
		Description   string
		WordsReader   io.Reader
		StartWord     string
		EndWord       string
		ExpectedError error
		ExpectedChain []string
	}{
		{
			Description:   "when GetChain succesfully finds a path for the words",
			WordsReader:   strings.NewReader("foo\nbar\nfoe\nfee"),
			StartWord:     "foo",
			EndWord:       "fee",
			ExpectedChain: []string{"foo", "foe", "fee"},
		},
		{
			Description:   "when GetChain fails due start word not being found",
			WordsReader:   strings.NewReader("foo\nbar\nfoe\nfee"),
			StartWord:     "foobar",
			EndWord:       "fee",
			ExpectedError: errors.New("error, start word 'foobar' not found"),
		},
		{
			Description:   "when GetChain fails due end word not being found",
			WordsReader:   strings.NewReader("foo\nbar\nfoe\nfee"),
			StartWord:     "foo",
			EndWord:       "feo",
			ExpectedError: errors.New("error, end word 'feo' not found"),
		},
		{
			Description:   "when GetChain fails to find a chain between the words",
			WordsReader:   strings.NewReader("foo\nbar\nfoobarr\nfee"),
			StartWord:     "foo",
			EndWord:       "fee",
			ExpectedError: errors.New("error, no chain found between 'foo' and 'fee'"),
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			client, err := NewClient(test.WordsReader)
			if err != nil {
				t.Fatalf("non expected error calling NewClient, err: %v", err)
			}

			chain, err := client.GetChain(test.StartWord, test.EndWord)

			if !reflect.DeepEqual(test.ExpectedError, err) {
				t.Fatalf("expected err to be %v, got %v", test.ExpectedError, err)
			}

			if !reflect.DeepEqual(test.ExpectedChain, chain) {
				t.Fatalf("expected chain to be %v, got %v", test.ExpectedChain, chain)
			}
		})
	}
}
