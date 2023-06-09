package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"basic_crud_go/src/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/get-notes", controller.GetAllNotes).Methods("GET")
	router.HandleFunc("/get-note/{noteId}", controller.GetNoteByID).Methods("GET")
	router.HandleFunc("/create-note", controller.CreateNote).Methods("POST")
	router.HandleFunc("/update-note/{noteId}", controller.UpdateNote).Methods("PUT")
	router.HandleFunc("/delete-note/{noteId}", controller.DeleteNote).Methods("DELETE")
	router.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}