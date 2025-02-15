package request

import "github.com/andys920605/meme-coin/pkg/errors"

type CreateMemeCoin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateMemeCoin) Valid() error {
	if r.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}

type UpdateMemeCoin struct {
	Description string `json:"description"`
}

func (r *UpdateMemeCoin) Valid() error {
	if r.Description == "" {
		return errors.New("invalid description")
	}
	return nil
}
