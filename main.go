package main // import "github.com/boltlessengineer/sedp"

import (
	webpush "github.com/SherClockHolmes/webpush-go"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func app(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.URL.Path)
	var localPath string
	if len(r.URL.Path) <= 1 {
		localPath = "app/index.html"
	} else {
		localPath = "app/" + r.URL.Path
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

func server(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	log.Println(r.Host, r.URL.Path, params)

	// Decode Subscription
	s := &webpush.Subscription{}
	json.Unmarshal([]byte("<YOUR_SUBSCRIPTION>"), s)
	resp, err := webpush.SendNotification([]byte("Test"), s, &webpush.Options{
		Subscriber:      "example@example.com",
		VAPIDPublicKey:  "<YOUR_VAPID_PUBLIC_KEY>",
		VAPIDPrivateKey: "<YOUR_VAPID_PRIVATE_KEY>",
		TTL:             30,
	});
	if err != nil {
		log.Println("=========[Error]=========");
	}
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/", app)
	http.HandleFunc("/server", server)

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
