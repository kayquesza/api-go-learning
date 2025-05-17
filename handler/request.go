package handler

import "fmt"

func errParanIsRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required and must be of type %s", name, typ)
}

// CreateOpening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Role == "" {
		return errParanIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParanIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParanIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParanIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParanIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParanIsRequired("salary", "int")
	}
	return nil
}
