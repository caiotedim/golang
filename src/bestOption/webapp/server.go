package webapp

import (
	"net/http"
	"fmt"
	//"log"
	"..//util"
	"..//dao"
)

func form(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST"{
		w.WriteHeader(400)
	}
	if util.Upload(w, r) {
		fmt.Fprintf(w, "File uploaded successfully!")
	} else {
		fmt.Fprintf(w, "Fail to upload file!")
	}

	name := r.Form("name")
	img1 := r.Form("mod1")
	path1 := r.Form("file1")
	path2 := r.Form("file2")
	dao.Save(name, path1, path2, img1)



}


func Server(bind *string, port *int) {
	http.HandleFunc("/upload", form)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *bind, *port), nil)
}
