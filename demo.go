package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Author string `json:"Author"`
	Desc   string `json:"Desc"`
}

var Books []Book

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, bck := range Books {
		if bck.Id == key {
			json.NewEncoder(w).Encode(bck)
		}
	}
}

func createNewBook(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Book
	json.Unmarshal(reqBody, &article)
	Books = append(Books, article)
	json.NewEncoder(w).Encode(article)
	fmt.Fprintf(w, "new book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update bbok")
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Book
	json.Unmarshal(reqBody, &article)
	for index, Book := range Books {
		if Book.Id == id {
			Books = append(Books[:index], article)
		}
		json.NewEncoder(w).Encode(article)
    fmt.Println()
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, Book := range Books {
		if Book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Prints the home page")
}

func allBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Displays  all books detail")
	json.NewEncoder(w).Encode(Books)
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/allbook", allBook)
	r.HandleFunc("/newbook", createNewBook).Methods("POST")
	r.HandleFunc("/singlebook/{id}", returnSingleBook)
	r.HandleFunc("/updateBook/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/deletebook/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("this is main program")
	Books = []Book{
		Book{Id: "1", Name: "C", Author: "Balaguruswamy", Desc: "About the C programming"},
		Book{Id: "2", Name: "Java", Author: "Balaguruswamy", Desc: "About the C programming"},
	}
	handleRequests()
}
