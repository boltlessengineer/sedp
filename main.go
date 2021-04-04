package main // import "github.com/boltlessengineer/sedp"

import (
	"encoding/json"
	"net/http"
)

func message(w http.ResponseWriter, req *http.Request) {
	data := make(map[string]string)
	data["type"] = "text"
	json.NewEncoder(w).Encode(data)
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/message", message)

	http.ListenAndServe(":3000", nil)
}
