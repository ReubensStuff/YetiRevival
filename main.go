package main

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func removePrependingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for len(r.URL.Path) != 0 && string(r.URL.Path[0]) == "/" {
			r.URL.Path = r.URL.Path[1:]
		}
		r.URL.Path = "/" + r.URL.Path
		next.ServeHTTP(w, r)
	})
}

func HandleUnknownRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("[OUT] !!!!!!!!!!!! Unknown endpoint:", r.URL.Path)
	output, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println("[ERR] Couldn't dump request:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	log.Println("Request data:", string(output))
}

func NvsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[OUT] Entering login...")
	switch r.RequestURI {
	//nvs/v1/global from New PvZ3 servers
	case "GET":
		http.Error(w, "{\n  \"Strings\": {\n    \"MinClientVersion\": \"1.0\"\n  },\n  \"Throttles\": {\n    \"LanguageIndonesian\": 0.0,\n    \"LanguageItalian\": 0.0,\n    \"LanguageBrazilianPortuguese\": 0.0,\n    \"LanguageSpainSpanish\": 0.0,\n    \"LanguageMexicanSpanish\": 0.0,\n    \"LanguageKorean\": 0.0,\n    \"LanguageGerman\": 0.0,\n    \"LanguageFrench\": 0.0,\n    \"LanguageJapanese\": 0.0\n  },\n  \"VersionHash\": \"024cc1f9c050b71c0d1b1684e06574bb2beae7009e0ad56eab2174b9389ee6bd\"\n}", http.StatusOK)
	default:
		log.Printf("[ERR] Login request is a %s when it's supposed to be a POST!\n", r.Method)
		http.Error(w, "Cannot use "+r.Method+" with /login", http.StatusMethodNotAllowed)
	}

}
func VolatileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[OUT] Entering login...")
	switch r.Method {
	//Volatile/VolatileOverrides.json from New PvZ3 servers
	case "GET":
		http.Error(w, "{\n    \"ErrorDialogOverrides\" : {\n        \"ERROR_SERVER_DOWN_MESSAGE_ANCHOR\" : \"PvZ 3 Q&A\"\n    }\n}", http.StatusOK)
	default:
		log.Printf("[ERR] Login request is a %s when it's supposed to be a POST!\n", r.Method)
		http.Error(w, "Cannot use "+r.Method+" with /login", http.StatusMethodNotAllowed)
	}
}

func main() {
	log.Printf("Old PvZ3 testing..")
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/nvs/v1/global", NvsHandler) // minimum game version / available languages
	r.Use(loggingMiddleware)
	r.PathPrefix("/").HandlerFunc(HandleUnknownRequest)

	port := "80"
	log.Printf("Server started on port %s\n", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: removePrependingSlashes(r),
	}

	server.SetKeepAlivesEnabled(false)

	panic(server.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)

	})
}
