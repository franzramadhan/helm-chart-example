package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getTrivia(w http.ResponseWriter, r *http.Request) {
	client := http.Client{}
	currentTime := time.Now()
	month := currentTime.Format("1")
	date := currentTime.Format("2")

	req, err := http.NewRequest(http.MethodGet, "http://numbersapi.com/"+month+"/"+date+"/date?json", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	defer res.Body.Close()

	var response map[string]interface{}

	fail := json.NewDecoder(res.Body).Decode(&response)
	if fail != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}

	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getTrivia)
	log.Println("Listening on *:8888")
	logging := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8888", logging))
}
