package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	version = "0.0.1"
)

var (
	isVersion              bool
	isRedirect             bool
	listenPortNumber       int
	redirectDestinationURL string
)

func main() {
	flag.BoolVar(&isVersion, "v", false, "print version")
	flag.IntVar(&listenPortNumber, "p", 80, "designate listen port number")
	flag.StringVar(&redirectDestinationURL, "r", "http://www.example.com/", "redirect destination URL")
	flag.Parse()

	switch {
	case isVersion:
		fmt.Println("gfrt " + version)
		return
	default:
		log.Printf("listening: %d", listenPortNumber)
		http.HandleFunc("/", rootHandler)
		if err := http.ListenAndServe(":"+strconv.Itoa(listenPortNumber), notFoundHandler()); err != nil {
			log.Fatal(err.Error())
		}
		return
	}
}

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
		writeDownLog(r.Method, r.RequestURI, http.StatusNotFound)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	responseCode := http.StatusOK
	switch r.Method {
	case http.MethodGet:
		if isRedirect {
			responseCode = http.StatusMovedPermanently
			http.Redirect(w, r, "http://www.example.com/", responseCode)
		} else {
			w.WriteHeader(responseCode)
			fmt.Fprintf(w, "TODO") // TODO: implement RSS feed response
		}
		writeDownLog(r.Method, r.RequestURI, responseCode)
		return
	case http.MethodPut: // switch to redirect mode
		responseCode = http.StatusNoContent
		w.WriteHeader(responseCode)
		isRedirect = true
		writeDownLog(r.Method, r.RequestURI, responseCode)
		return
	case http.MethodDelete: // switch to feed mode
		responseCode = http.StatusNoContent
		w.WriteHeader(responseCode)
		isRedirect = false
		writeDownLog(r.Method, r.RequestURI, responseCode)
		return
	default:
		responseCode = http.StatusMethodNotAllowed
		w.WriteHeader(responseCode)
		fmt.Fprintf(w, "Method not allowed")
		writeDownLog(r.Method, r.RequestURI, responseCode)
		return
	}
}

func writeDownLog(requestMethod, requestURI string, responseCode int) {
	log.Printf("%s %s %d", requestMethod, requestURI, responseCode)
}
