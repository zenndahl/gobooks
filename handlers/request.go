package handlers

import (
	"fmt"
	"time"
)

func errParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}

func (r *CreateBookRequest) Validate() error {
	if r.Name == "" {
		return errParamRequired("name", "string")
	}
	if r.Author == "" {
		return errParamRequired("author", "string")
	}
	if r.Genre == "" {
		return errParamRequired("genre", "string")
	}
	if r.Year == "" {
		return errParamRequired("year", "bool")
	}
	return nil
}

type UpdateBookRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}

func (r *UpdateBookRequest) Validate() error {
	if r.Name == "" || r.Author == "" || r.Genre == "" || r.Year == "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type FinishBookRequest struct {
	Name       string    `json:"name"`
	Author     string    `json:"author"`
	Genre      string    `json:"genre"`
	Year       string    `json:"year"`
	Finished   string    `json:"finished"`
	FinishedAt time.Time `json:"finishedAt"`
	Rating     int       `json:"rating"`
}

func (r *FinishBookRequest) Validate() error {
	if r.Name == "" {
		return errParamRequired("name", "string")
	}
	if r.Author == "" {
		return errParamRequired("author", "string")
	}
	if r.Genre == "" {
		return errParamRequired("genre", "string")
	}
	if r.Year == "" {
		return errParamRequired("year", "bool")
	}
	if r.Finished == "" {
		return errParamRequired("year", "bool")
	}
	if r.Rating < 0 || r.Rating > 10 {
		return fmt.Errorf("param: rating (type: int) must be between 0 and 10")
	}

	return nil
}
