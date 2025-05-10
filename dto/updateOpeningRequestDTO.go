package dto

import "fmt"

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (req *UpdateOpeningRequest) Validate() error {
	if req.Role != "" || req.Company != "" || req.Location != "" || req.Remote != nil || req.Link != "" || req.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided for update")
}
