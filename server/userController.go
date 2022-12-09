package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		return s.handleCreateUser(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleUserById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetUser(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteUser(w, r)
	}

	if r.Method == http.MethodPut {
		return s.handleUpdateUser(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	idAsStr, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}

	user, err := s.UserStorer.GetUser(int64(idAsStr))
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, user)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	creatUserReq := &CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(creatUserReq)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, apiError{Err: err.Error()})
	}

	err = s.UserStorer.CreateUser(creatUserReq.Name, creatUserReq.Email, creatUserReq.Country)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, apiError{Err: err.Error()})
	}

	return WriteJson(w, http.StatusCreated, creatUserReq)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	updateUserReq := &UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(updateUserReq)
	if err != nil {
		return WriteJson(w, http.StatusBadRequest, apiError{Err: err.Error()})
	}

	vars := mux.Vars(r)
	id := vars["id"]
	idAsStr, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}

	user, err := s.UserStorer.GetUser(int64(idAsStr))
	if err != nil {
		return err
	}

	user.Name = updateUserReq.Name
	user.Email = updateUserReq.Email
	user.Country = updateUserReq.Country

	err = s.UserStorer.UpdateUser(user)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, user)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]
	idAsStr, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}

	err = s.UserStorer.DeleteUser(int64(idAsStr))
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, "User deleted")
}
