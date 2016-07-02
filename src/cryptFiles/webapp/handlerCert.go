package webapp

import (
	"net/http"
	//"net/url"
	"log"
)

func handlerCert(w http.ResponseWriter, r *http.Request) {
	log.Printf("Uploading certificate...")
	w.Header().Set("Server", "CryptoFiles")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	if r.Method == "OPTIONS" {
		w.Header().Set("Allow", "GET, DELETE, POST, OPTIONS")
		w.WriteHeader(200)
	} else if r.Method == "POST" {
		key = make([]byte, r.ContentLength, r.ContentLength)
		r.Body.Read(key)
		
		
	}
}
