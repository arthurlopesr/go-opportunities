package dto

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (req *CreateOpeningRequest) Validate() error {
	fields := []struct {
		value interface{}
		name  string
		typ   string
	}{
		{req.Role, "role", "string"},
		{req.Company, "company", "string"},
		{req.Location, "location", "string"},
		{req.Remote, "remote", "bool"},
		{req.Link, "link", "string"},
		{req.Salary > 0, "salary", "int"},
	}

	for _, field := range fields {
		switch value := field.value.(type) {
		case string:
			if value == "" {
				return errParamIsRequired(field.name, field.typ)
			}
		case *bool:
			if value == nil {
				return errParamIsRequired(field.name, field.typ)
			}
		case bool:
			if !value {
				return errParamIsRequired(field.name, field.typ)
			}
		case int64:
			if value <= 0 {
				return errParamIsRequired(field.name, field.typ)
			}
		}
	}

	return nil
}
