package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// need
// go:embed frontend/dist
//var frontend embed.FS

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()

	http.Handle(
		"/api/v1/law/",
		http.HandlerFunc(getRandomLaw),
	)

	//stripped, err := fs.Sub(frontend, "frontend/dist")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//frontendFS := http.FileServer(http.FS(stripped))
	//http.Handle("/", frontendFS)

	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

type Law struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

var HackerNews = []Law{
	{
		Name:       "Amdahl's Law",
		Definition: "Amdahl's Law is a formula which shows the potential speedup of a computational task which can be achieved by increasing the resources of a system.",
	},
	{
		Name:       "Conway's Law",
		Definition: "This law suggests that the technical boundaries of a system will reflect the structure of the organisation.",
	},
	{
		Name:       "Gall's Law",
		Definition: "A complex system that works is invariably found to have evolved from a simple system that worked.",
	},
}

func getRandomLaw(w http.ResponseWriter, r *http.Request) {
	// 1. generate a random number
	// 2. generate a random law
	// 3. transform law to json
	// 4. response json

	randomLaw := HackerNews[rand.Intn(len(HackerNews))]
	j, err := json.Marshal(randomLaw)
	if err != nil {
		http.Error(w, "couldn't retrieve random hacker law", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))

}
