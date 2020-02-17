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
	r.HandleFunc("/{key}/{u}", auth)
	// r.HandleFunc("/LoginService", auth)
	// http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func auth(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("url %s", r.URL.String())
	name, passwd, ok := r.BasicAuth()
	logrus.Debugf("%s request. %v %v %v", time.Now().Format(time.RFC1123), name, passwd, ok)
	// w.WriteHeader(http.StatusForbidden)
	w.WriteHeader(http.StatusOK)
}
