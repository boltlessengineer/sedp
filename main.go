package main // import "github.com/boltlessengineer/sedp"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func message(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	data["type"] = "text"
	json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
}

func app(w http.ResponseWriter, req *http.Request) {
	var localPath string
	if len(req.URL.Path) <= 1 {
		localPath = "app/index.html"
	} else {
		localPath = "app/" + req.URL.Path
	}
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func main() {
	http.HandleFunc("/", app)
	http.HandleFunc("/message", message)

	http.ListenAndServe(":3000", nil)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
