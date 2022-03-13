package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func garbages2json(garbages []Artifact) []byte {
	encjson, _ := json.Marshal(garbages)
	return encjson
}

func artifact2json(a Artifact) []byte {
	encjson, _ := json.Marshal(a)
	return encjson
}

func gainArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Json PLZ~")
		return
	}
	// body, _ := ioutil.ReadAll(r.Body)
	// var params map[string]interface{}
	// fmt.Println(body)
	data := struct {
		Resin  int32
		Domain string
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Invalid json format: %s", err)
		return
	}
	// fmt.Println(data.Resin)
	// garbages := gainArtifacts("Peak of Vindagnyr", 20)
	garbages := gainArtifacts(data.Domain, data.Resin)
	w.Write(garbages2json(garbages))
}

func artifactEnhancingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Json PLZ~")
		return
	}

	body := struct {
		Embryo   Artifact
		DogFoods []DogFood
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Fprintf(w, "Invalid json format: %s", err)
		return
	}

	a := body.Embryo
	a.levelUp(body.DogFoods)
	w.Write(artifact2json(a))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong")
}

func main() {

	logger := log.New(os.Stdout, "[http] ", log.LstdFlags)
	logger.Printf("hello~")

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/gain", gainArtifactsHandler)
	http.HandleFunc("/enhance", artifactEnhancingHandler)
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
