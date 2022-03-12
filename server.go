package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func garbages2json(garbages []Artifact) []byte {
	encjson, _ := json.Marshal(garbages)
	return encjson
}

func gainArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	garbages := gainArtifacts("Peak of Vindagnyr", 20)
	w.Write(garbages2json(garbages))
}

func main() {

	logger := log.New(os.Stdout, "[http] ", log.LstdFlags)
	logger.Printf("hello~")

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/gainArtifacts", gainArtifactsHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
