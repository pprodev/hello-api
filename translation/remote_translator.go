package translation

import (
	"fmt"
	"github.com/pprodev/hello-api/handlers/rest"
	"log"
	"strings"
)

// verify that the struct we are building satisfies the interface.this will cause a compile time error if not satisfied.
var _ rest.Translator = &RemoteService{}

// RemoteService will allow for external calls to existent service for translations.
type RemoteService struct {
	client HelloClient
	cache  map[string]string
}

// HelloClient will call external service.
type HelloClient interface {
	Translate(word, language string) (string, error)
}

// NewRemoteService creates a new implementation of RemoteService.
func NewRemoteService(client HelloClient) *RemoteService {
	return &RemoteService{
		client: client,
		cache:  make(map[string]string),
	}
}

// Translate will take a given word and try to find the result using the client.
func (s *RemoteService) Translate(word string, language string) string {
	word = strings.ToLower(word)
	language = strings.ToLower(language)

	key := fmt.Sprintf("%s:%s", word, language)

	tr, ok := s.cache[key]

	if ok {
		return tr
	}

	resp, err := s.client.Translate(word, language)
	if err != nil {
		log.Println(err)
		return ""
	}
	s.cache[key] = resp
	return resp
}
