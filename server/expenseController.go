package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleExpense(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		return s.handleCreateExpense(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleExpenseById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetExpense(w, r)
	}

	if r.Method == http.MethodDelete {
		return s.handleDeleteExpense(w, r)
	}

	if r.Method == http.MethodPut {
		return s.handleUpdateExpense(w, r)
	}

	if r.Method == http.MethodPost {
		return s.handleUploadExpenseReceipt(w, r)
	}

	return fmt.Errorf("method %s not allowed", r.Method)
}

func (s *APIServer) handleGetExpense(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	idAsStr, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}

	expense, err := s.Storer.GetExpense(int64(idAsStr))
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, expense)
}

func (s *APIServer) handleCreateExpense(w http.ResponseWriter, r *http.Request) error {
	createExpenseReq := &CreateExpenseRequest{}

	err := json.NewDecoder(r.Body).Decode(createExpenseReq)
	if err != nil {
		return err
	}

	err = s.Storer.CreateExpense(createExpenseReq.UserId, createExpenseReq.Amount, createExpenseReq.Description)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusCreated, nil)
}

func (s *APIServer) handleUpdateExpense(w http.ResponseWriter, r *http.Request) error {
	updateExpenseRequest := &UpdateExpenseRequest{}

	err := json.NewDecoder(r.Body).Decode(updateExpenseRequest)
	if err != nil {
		return err
	}

	panic("not implemented")
}

func (s *APIServer) handleDeleteExpense(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	idAsStr, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}

	err = s.Storer.DeleteExpense(int64(idAsStr))
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, nil)
}

func (s *APIServer) handleUploadExpenseReceipt(w http.ResponseWriter, r *http.Request) error {
	r.ParseMultipartForm(1024 * 1024 * 10)

	file, _, err := r.FormFile("recipt")
	if err != nil {
		return err 	
	}
	
	defer file.Close()

	s.Storer.StoreRecipt(1, file)

	return WriteJson(w, http.StatusOK, "recipt uploaded")
}

