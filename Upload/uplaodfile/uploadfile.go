package uplaodfile

import (
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ioutil.ReadAll()
	} else if r.Method == "POST" {

	}

}
