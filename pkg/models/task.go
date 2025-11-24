package models

import (
	"errors"
	"net/http"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	EndDate     string `json:"end_date"`
	Categorie   string `json:"categorie"`
}

func (t *Task) Bind(r *http.Request) error {

	if t.Title == "" {
		return errors.New("Title cannot be empty")
	} else if t.EndDate == "" {
		return errors.New("EndDate cannot be empty")
	} else if t.Categorie == "" {
		return errors.New("Categorie cannot be empty")
	}

	return nil
}
