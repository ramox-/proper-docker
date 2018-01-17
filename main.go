package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID      string   `json:"id,omitempty"`
	Name    string   `json:"name,omitempty"`
	Contact *Contact `json:"address,omitempty"`
}

type Contact struct {
	Tel     string `json:"tel,omitempty"`
	Address string `json:"address,omitempty"`
}

var people []Person

func GetPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, x := range people {
		if x.ID == params["id"] {
			json.NewEncoder(w).Encode(x)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = uuid.New().String()
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, x := range people {
		if x.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: uuid.New().String(), Name: "FOO BARSON", Contact: &Contact{Tel: "001122334455", Address: "Someplace 111, 12345 Foo"}})
	people = append(people, Person{ID: uuid.New().String(), Name: "BAR FOOSON", Contact: &Contact{Tel: "009988776655", Address: "Somewhere 222, 98765 Bar"}})
	router.HandleFunc("/person", GetPersons).Methods("GET")
	router.HandleFunc("/person", CreatePerson).Methods("POST")
	router.HandleFunc("/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
