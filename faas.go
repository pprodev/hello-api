// Package faas is used for function definitions
package faas

import (
	"github.com/pprodev/hello-api/handlers/rest"
	"github.com/pprodev/hello-api/translation"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	translateHandler.TranslateHandler(w, r)
}
