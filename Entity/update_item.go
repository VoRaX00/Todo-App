package Entity

import "errors"

type UpdateItem struct {
	Title       *string `json:"title" db:"title"`
	Description *string `json:"description" db:"description"`
	Done        *bool   `json:"done" db:"done"`
}

func (item *UpdateItem) Validate() error {
	if item.Title == nil && item.Description == nil && item.Done == nil {
		return errors.New("UpdateItem needs a title, description, done")
	}
	return nil
}
