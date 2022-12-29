// Package faas is used for function definitions
package faas

import (
	"github.com/pprodev/hello-api/handlers/rest"
	"net/http"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
