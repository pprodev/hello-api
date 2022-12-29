// Package rest houses all rest handlers.
package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Translator interface {
	Translate(word string, language string) string
}

// TranslateHandler will translate calls for caller.
type TranslateHandler struct {
	service Translator
}

// NewTranslateHandler will create a new instance of the handler using a translation service.
func NewTranslateHandler(service Translator) *TranslateHandler {
	return &TranslateHandler{
		service: service,
	}
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler will take a given request with a path value of the word to be translated and a query parameter of the language to translate to.
func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")

	log.Printf("language set to %s, and word is set to %s", language, word)

	wordTranslation := t.service.Translate(word, language)

	log.Printf("translation is %s", wordTranslation)

	if wordTranslation == "" {
		w.WriteHeader(404)
		return
	}

	resp := Resp{
		Language:    language,
		Translation: wordTranslation,
	}

	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
