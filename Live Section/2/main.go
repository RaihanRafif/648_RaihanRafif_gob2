package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"live_2/models"
)

type Punyaku struct {
	Name string ` json:"name" form:"name"`
	Hobi string ` json:"hobi" form:"hobi"`
}

type CustomRespons struct {
	MessageCode int
	Status      bool
	Punyaku     Punyaku
}

func main() {
	mux := http.NewServeMux()

	endpoint := http.HandlerFunc(greet)

	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening to port 8080")

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func middleware1(next http.Handler) http.Handler {
	var p = models.Punyaku{
		"Rehan",
		"ngoding",
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal(p)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})
}

func middleware2(next http.Handler) http.Handler {
	var p = Punyaku{
		"Rehan",
		"ngoding",
	}

	var t = CustomRespons{
		201,
		true,
		p,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, _ := json.Marshal(t)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})
}
