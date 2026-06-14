package handlers

import (
	model "bookapplication/Model"
	service "bookapplication/Service"
	"encoding/json"
	"net/http"
)

type UserHandler struct{
	AuthService service.AuthService
}

func (UHandler *UserHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request){
	var NewUser model.User
	if err:= json.NewDecoder(r.Body).Decode(&NewUser); err!= nil{
		http.Error(w,"The New User payload is invalid ... %v",http.StatusBadRequest)
		return
	}

	//call the repo and pass the user to the method
	if err:= UHandler.AuthService.RegisterNewUser(&NewUser);err!= nil{
		http.Error(w,"The CreateNewUser method failed...",http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"Message":"The New user has been successfully created"})
}

func (UHandler *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request){
	var LoginUser model.LoginDto
	if err:= json.NewDecoder(r.Body).Decode(&LoginUser);err!= nil{
		http.Error(w,"The login payload is invalid...",http.StatusBadRequest)
		return
	}

	//call the repo to Login
	AuthToken,err:= UHandler.AuthService.LoginUser(&LoginUser)
	if err!= nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AuthToken)
}

