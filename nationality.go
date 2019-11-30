package person

import (
	"github.com/go-msvc/item"
)

//Nationality of a person
type Nationality struct {
	PersonID  item.ID `json:"id-person"`
	CountryID item.ID `json:"id-country"`
	Nr        string  `json:"nr"` //national id nr
}
