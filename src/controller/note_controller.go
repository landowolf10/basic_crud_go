package controller

import (
	"basic_crud_go/src/config"
	"basic_crud_go/src/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var noteList []model.Note
var result model.CreatedNote
var singleResult model.Note
var note model.Note
var noteId string

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	noteList = nil
	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM notes")

	if err != nil {
		log.Print(err)
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&note.NoteId, &note.UserId, &note.Owner, &note.Title, &note.Content)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			noteList = append(noteList, note)
		}
	}

	setJson(w, r, "Success!", "GET")
}

func checkIfNoteExists(noteId string) bool {
	db := config.Connect()
	defer db.Close()
	//var product entities.Product
	_, err := db.Query("SELECT * FROM notes WHERE noteid=?", noteId)

	if err != nil {
		log.Print(err)
		//respondWithError(w, http.StatusInternalServerError, err.Error())
		return false
	}

	return true
}

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	noteId = mux.Vars(r)["noteId"]

	if !checkIfNoteExists(noteId) {
		json.NewEncoder(w).Encode("Note not found!")
		return
	}

	//rows, err := db.Query("SELECT * FROM notes WHERE noteid=?", noteId)

	singleResult = model.Note{NoteId: note.NoteId, UserId: note.UserId, Owner: note.Owner, Title: note.Title, Content: note.Content}
	setJson(w, r, "Success!", "GET")

	//setJson(w, r, "Success!", "GET")
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&note); err != nil {
		log.Print(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("INSERT INTO notes (userid, owner, title, content) VALUES (?, ?, ?, ?)", note.UserId, note.Owner, note.Title, note.Content)

	if err != nil {
		log.Print(err)
		return
	}

	result = model.CreatedNote{UserId: note.UserId, Owner: note.Owner, Title: note.Title, Content: note.Content}

	setJson(w, r, "Note created successfully!", "POST")
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	noteId = mux.Vars(r)["noteId"]

	if err := decoder.Decode(&note); err != nil {
		log.Print(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("UPDATE notes SET owner=?, title=?, content=? WHERE noteid=?", note.Owner, note.Title, note.Content, noteId)

	if err != nil {
		log.Print(err)
		return
	}

	result = model.CreatedNote{UserId: note.UserId, Owner: note.Owner, Title: note.Title, Content: note.Content}

	setJson(w, r, "Note updated successfully!", "PUT")
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	db := config.Connect()
	defer db.Close()

	noteId := mux.Vars(r)["noteId"]

	defer r.Body.Close()

	_, err := db.Exec("DELETE FROM notes WHERE noteid=?", noteId)

	if err != nil {
		log.Print(err)
		return
	}

	setJson(w, r, "Note deleted successfully!", "DELETE")
}

func setJson(w http.ResponseWriter, r *http.Request, message string, httpMethod string) {
	var response model.Response
	var singleResponse model.SingleResponse
	var createdResponse model.CreatedResponse
	var deletedResponse model.DeletedResponse

	singleResponse.Status = 200
	singleResponse.Message = message
	singleResponse.Data = singleResult

	response.Status = 200
	response.Message = message
	response.Data = noteList

	createdResponse.Status = 200
	createdResponse.Message = message
	createdResponse.Data = result

	deletedResponse.Status = 200
	deletedResponse.Message = message

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if httpMethod == "POST" || httpMethod == "PUT" {
		json.NewEncoder(w).Encode(createdResponse)
	} else if httpMethod == "DELETE" {
		json.NewEncoder(w).Encode(deletedResponse)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}
