package handlers

import "fmt"

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

type UpdateOpeningRequest struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   string `json:"year"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Name == "" || r.Author == "" || r.Genre == "" || r.Year == ""{
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
