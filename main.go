package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NvsHandler(w http.ResponseWriter, r *http.Request) { //this handles the login process, seemingly ignores the strings below? not sure why, this game is so strange
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		http.Error(w, "{\n  \"Strings\": {\n    \"MinClientVersion\": \"1.0\"\n  },\n  \"Throttles\": {\n    \"LanguageIndonesian\": 0.0,\n    \"LanguageItalian\": 0.0,\n    \"LanguageBrazilianPortuguese\": 0.0,\n    \"LanguageSpainSpanish\": 0.0,\n    \"LanguageMexicanSpanish\": 0.0,\n    \"LanguageKorean\": 0.0,\n    \"LanguageGerman\": 0.0,\n    \"LanguageFrench\": 0.0,\n    \"LanguageJapanese\": 0.0\n  },\n  \"VersionHash\": \"024cc1f9c050b71c0d1b1684e06574bb2beae7009e0ad56eab2174b9389ee6bd\"\n}", http.StatusOK)
	}
}

func v2Handler(w http.ResponseWriter, r *http.Request) { //new pvz3 related, ignore
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusOK)
	}
}

func v1Handler(w http.ResponseWriter, r *http.Request) { //new pvz3 related, ignore
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusOK)
	}
}

func TeleHandler(w http.ResponseWriter, r *http.Request) { //new pvz3 related, ignore
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusOK)
	}
}

func PSHandler(w http.ResponseWriter, r *http.Request) { //This is what gets the asset download request to appear, even if all you send is a status 200 response which is what we're doing here.
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	var dir string
	log.Printf("Old PvZ3 testing")

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/nvs/v1/global", NvsHandler) //these are used in old pvz3
	r.HandleFunc("/session/v1/init", v1Handler) //this is used in 2021 pvz3
	r.HandleFunc("/session/v2/init", v2Handler) //these are used in 2024 pvz3
	r.HandleFunc("/pub/telemetry/v1/publish-pre-session-client-telemetry-events", TeleHandler) //these are used in 2024 pvz3
	r.HandleFunc("/player-session/v1/initialize", PSHandler) //these are used in old pvz3
	r.PathPrefix("").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir)))) //static files, aka assets (anything not in here will be in the folders)

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:80",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 0 * time.Second,
		ReadTimeout:  0 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
