package main

import (
	"api-test/datastore"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	stu datastore.StudentStore
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	stu = &datastore.Students{}
	stu.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	//api.HandleFunc("/books/authors/{author}", searchByAuthor).Methods(http.MethodGet)
	//api.HandleFunc("/books/book-name/{bookName}", searchByBookName).Methods(http.MethodGet)
	api.HandleFunc("/student/exam_roll/{exam_roll}", searchByExamRoll).Methods(http.MethodGet)
	//api.HandleFunc("/book/isbn/{isbn}", deleteByISBN).Methods(http.MethodDelete)
	//api.HandleFunc("/book", createBook).Methods(http.MethodPost)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
