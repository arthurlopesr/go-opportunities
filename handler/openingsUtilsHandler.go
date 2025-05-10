package handler

import "fmt"

func validateHasOpeningId(openingId string) error {
	if openingId == "" {
		return fmt.Errorf("error: openingId is required")
	}
	return nil
}
