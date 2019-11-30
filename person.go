package person

import (
	"time"

	"github.com/go-msvc/errors"
)

//Person ...
type Person struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Dob     time.Time `json:"dob"`
	Gender  Gender    `json:"gender"`
}

//Validate ...
func (p Person) Validate() error {
	if len(p.Name) == 0 {
		return errors.Errorf("missing name")
	}
	if len(p.Surname) == 0 {
		return errors.Errorf("missing surname")
	}
	if p.Dob.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)) {
		return errors.Errorf("dob before 1900-01-01")
	}
	if err := p.Gender.Validate(); err != nil {
		return errors.Errorf("invalid gender")
	}
	return nil
}
