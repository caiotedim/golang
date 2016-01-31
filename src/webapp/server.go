package webapp

import (
	"net/http"
	"net/url"
	"log"
	"bo"
	"fmt"
)

/*type QueryString struct {
	path string
	file string
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: [%s]", r.Method)
	w.Header().Set("Server", "CryptoFiles")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
        log.Fatal(err)
        w.WriteHeader(500)
    }	
	
	if r.URL.Path != "/" {
		w.WriteHeader(400)
	}
	
	if r.Method == "OPTIONS" {
		w.Header().Set("Allow", "GET, DELETE, POST, OPTIONS")
		w.WriteHeader(200)
	} else if r.Method == "GET" {
		
		if _, ok := m["file"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400)
		} else if _, ok := m["path"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400) 
			
		} else { 
		
			queryString := map[string]string{
				"file": m["file"][0],
				"path": m["path"][0],
			}
			
			body, respCode := bo.GetFile(queryString)
			//w.Header().Set("X-Content", string(body))
			
			w.WriteHeader(respCode)
			w.Write(body)
			
		}
		
		
	} else if r.Method == "POST" {
		
		if _, ok := m["file"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400)
		} else if _, ok := m["path"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400) 
			
		} else { 
		
			queryString := map[string]string{
				"file": m["file"][0],
				"path": m["path"][0],
			}
		
			//log.Printf("%s ", queryString)
			//file := "teste"
			// call to manipulatin data
			file := make([]byte, r.ContentLength, r.ContentLength)
			r.Body.Read(file)
			if ok, respCode := bo.SetFile(queryString, file); ok {
				w.WriteHeader(respCode)
			} else {
				w.WriteHeader(500)
			}
		}
	} else if r.Method == "DELETE" {
		
		if _, ok := m["file"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400)
		} else if _, ok := m["path"]; !ok {
			log.Printf("Invalid Parameters!")
			w.WriteHeader(400) 	
		} else { 
		
			queryString := map[string]string{
				"file": m["file"][0],
				"path": m["path"][0],
			}
			
			body, respCode := bo.RemoveFile(queryString)
			//w.Header().Set("X-Content", string(body))
			
			w.WriteHeader(respCode)
			w.Write(body)
		}
		
	} else {
		w.WriteHeader(400)
	}
	
}

func Server(bind *string, port *int) {
	http.HandleFunc("/cert", handlerCert)
	http.HandleFunc("/", handler)
    http.ListenAndServe(fmt.Sprintf("%s:%d", *bind, *port), nil)
}
