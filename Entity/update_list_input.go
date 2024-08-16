package Entity

import "errors"

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (in *UpdateListInput) Validate() error {
	if in.Title == nil && in.Description == nil {
		return errors.New("update list title and description is nil")
	}

	return nil
}
