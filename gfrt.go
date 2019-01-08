package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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
	externalHostname       = "127.0.0.1"
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
		setExternalHostname()
		log.Printf("listening: %d", listenPortNumber)

		http.HandleFunc("/feed", feedHandler)
		http.HandleFunc("/example1", exampleArticle1Handler)
		if err := http.ListenAndServe(":"+strconv.Itoa(listenPortNumber), nil); err != nil {
			log.Fatal(err.Error())
		}
		return
	}
}

func setExternalHostname() {
	e := os.Getenv("GFRT_EXTERNAL_HOSTNAME")
	if e != "" {
		externalHostname = e
	}
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	responseCode := http.StatusOK
	switch r.Method {
	case http.MethodGet:
		if isRedirect {
			responseCode = http.StatusMovedPermanently
			http.Redirect(w, r, "http://www.example.com/", responseCode)
		} else {
			w.WriteHeader(responseCode)
			feedResponse(w)
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

func feedResponse(w http.ResponseWriter) {
	templateFilePath := "feed.goxml"
	t := template.Must(template.ParseFiles(templateFilePath))
	m := map[string]string{
		"EXTERNAL_URL": externalHostname,
	}
	if err := t.ExecuteTemplate(w, templateFilePath, m); err != nil {
		log.Fatal(err.Error())
	}
}

func exampleArticle1Handler(w http.ResponseWriter, r *http.Request) {
	responseCode := http.StatusOK
	switch r.Method {
	case http.MethodGet:
		templateFilePath := "example1.gohtml"
		t := template.Must(template.ParseFiles(templateFilePath))
		m := map[string]string{
			"EXTERNAL_URL": externalHostname,
		}
		if err := t.ExecuteTemplate(w, templateFilePath, m); err != nil {
			log.Fatal(err.Error())
		}
		writeDownLog(r.Method, r.RequestURI, http.StatusOK)
	default:
		responseCode = http.StatusMethodNotAllowed
		w.WriteHeader(responseCode)
		fmt.Fprintf(w, "Method not allowed")
		writeDownLog(r.Method, r.RequestURI, responseCode)
		return
	}
}
