package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	panic("not implemented")
}

func (s *APIServer) handleCreateExpense(w http.ResponseWriter, r *http.Request) error {
	createExpenseReq := &CreateExpenseRequest{}
	err := json.NewDecoder(r.Body).Decode(createExpenseReq)
	if err != nil {
		return err
	}

	err = s.UserStorer.CreateExpense(createExpenseReq.UserId, createExpenseReq.Amount, createExpenseReq.Description)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusCreated, nil)
}

func (s *APIServer) handleUpdateExpense(w http.ResponseWriter, r *http.Request) error {
	panic("not implemented")
}

func (s *APIServer) handleDeleteExpense(w http.ResponseWriter, r *http.Request) error {
	panic("not implemented")
}

func (s *APIServer) handleUploadExpenseReceipt(w http.ResponseWriter, r *http.Request) error {
	panic("not implemented")
}
