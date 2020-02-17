package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debug("server start ...")

	r := mux.NewRouter()
	r.HandleFunc("/", auth)
	// http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func auth(w http.ResponseWriter, r *http.Request) {
	logrus.Debug("xxxxxxxxxxxxx")
	w.Write([]byte("xxhellow world"))
}
