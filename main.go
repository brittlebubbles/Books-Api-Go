package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "log"
  "math/rand"
  "net/http"
  "strconv"
)

//bOOK Struct
type Book struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Author *Author `json:"author"`
}

type Author struct {
  Firstname  string `json:"firstname"`
  Lastname string `json:"lastname"`
}

var books []Book

//Functions for routers
func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  _ = json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range books {
            if item.ID == params["id"] {
              _ = json.NewEncoder(w).Encode(item)
              return
            }
  }
  _ = json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa(rand.Intn(100000000))
  books = append(books, book)
  _ = json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.ID == params["id"] {
      books = append(books[:index], books[index+1:]...)
      var book Book
      _ = json.NewDecoder(r.Body).Decode(&book)
      book.ID = params["id"]
      books = append(books, book)
      _ = json.NewEncoder(w).Encode(book)
      return
    }
  }
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range books {
    if item.ID == params["id"] {
      books = append(books[:index], books[index+1:]...)
      break
    }
  }
  _ = json.NewEncoder(w).Encode(books)
}

func main()  {
  //  Initialize Router
  r := mux.NewRouter()

  //Mock Data
  books = append(books, Book{ID: "1", Isbn: "32213", Title: "The Gods Are not to be blamed", Author: &Author{Firstname: "Dave", Lastname: "mUZE"}})
  books = append(books, Book{ID: "2", Isbn: "11113", Title: "sUNcITY", Author: &Author{Firstname: "Joe", Lastname: "Wqn"}})

  // Route Handlers
  r.HandleFunc("/api/books", getBooks).Methods("GET")
  r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
  r.HandleFunc("/api/books", createBook).Methods("POST")
  r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
  r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

  //Listening to Ports
  log.Fatal(http.ListenAndServe(":8000", r))

}

