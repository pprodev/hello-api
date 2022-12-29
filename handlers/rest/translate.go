// Package rest houses all rest handlers
package rest

import (
	"encoding/json"
	"github.com/pprodev/hello-api/translation"
	"log"
	"net/http"
	"strings"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := strings.ReplaceAll(r.URL.Path, "/", "")

	log.Printf("language set to %s, and word is set to %s", language, word)

	wordTranslation := translation.Translate(word, language)

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
