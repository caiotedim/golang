package util

import (
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"fmt"
	"io"
)

func createDirectory(basePath string) bool {
	if err := os.MkdirAll(basePath, 0755); err != nil {
		log.Printf("Error to create path:[%s] \n %s", basePath, err)
		return false
	}
	return true
}

func createFile(name string, content []byte) bool {
	basePath := "/tmp/bestOption/"
	basePath += "/"
	basePath += name

	if ! createDirectory(basePath) {
		return false
	}

	if err := ioutil.WriteFile(basePath, content, 0644); err != nil {
		log.Printf("Error to create file:[%s] \n %s", name, err)
		return false
	}
	return true
}

func Upload(w http.ResponseWriter, r *http.Request) bool {
	basePath := "/tmp/bestOption/"
	if ! createDirectory(basePath) {
		return false
	}
	for i:= 1; i <=2; i++ {
		upload := fmt.Sprintf("%s%d", "file", i)
		file, header, err := r.FormFile(upload)

		if err != nil {
			//fmt.Fprintln(w, err)
			log.Printf("%s \n %s", w, err)
			return false
		}

		defer file.Close()

		out, err := os.Create(basePath + header.Filename)
		if err != nil {
			//fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			log.Printf("%s \n %s", w, err)
			return false
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			//fmt.Fprintln(w, err)
			log.Printf("%s \n %s", w, err)
			return false
		}

	}
	return true
}