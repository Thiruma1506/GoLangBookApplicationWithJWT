package handlers

import (
	model "bookapplication/Model"
	repository "bookapplication/Repository"
	"encoding/json"
	"log"
	"net/http"
)

type BookHandler struct {
	Repo repository.Repo
}

func (repo *BookHandler) CreateNewBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	defer r.Body.Close()

	//Decode the request payload
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//call the repo to pass the payload
	if err := repo.Repo.CreateNewBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//write the response message
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"Message": "The Book has been persisted successfully"}); err != nil {
		log.Fatalf("The response has nor been sent... %v", err.Error())
	}
}

func (repo *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdatedBook model.Book
	var IncomingId string
	defer r.Body.Close()

	//Decode the book from the request and save it in the local var
	if err := json.NewDecoder(r.Body).Decode(&UpdatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	IncomingId = UpdatedBook.BookId
	//pass the address reference of the UpdatedBook Var
	if err := repo.Repo.UpdateBook(&UpdatedBook, IncomingId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//write the response message and the response status to the response writer
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"Message": "The Book with id " + IncomingId + " has been successfully updated"})
}

func (repo *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	// CALL THE REPO AND THE METHOD TO FETCH THE BOOK BY ID
	var id string
	id = r.URL.Query().Get("id")

	book, err := repo.Repo.GetBookById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contenet-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (repo *BookHandler) DeleteBookById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		log.Fatalf("The ID in request is invalid %v", id)
		http.Error(w, "requesing ID is invalid", http.StatusBadRequest)
		return
	}

	if err := repo.Repo.DeleteBookById(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Message": "The Book with ID - " + id + " has been deleted successfully"})
}

func (repo *BookHandler) CreateListOfBooks(w http.ResponseWriter, r *http.Request) {
	var Books []*model.Book

	if err := json.NewDecoder(r.Body).Decode(&Books); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repo.Repo.CreateListOfBooks(Books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message: ": "The List of Books are successfully persisted"})
}
