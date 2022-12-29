package main

import (
	"fmt"
	"github.com/pprodev/hello-api/handlers"
	"github.com/pprodev/hello-api/handlers/rest"
	"github.com/pprodev/hello-api/translation"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)
	//log.Fatal(http.ListenAndServe(addr, mux))

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

//type Resp struct {
//	Language    string `json:"language"`
//	Translation string `json:"translation"`
//}
