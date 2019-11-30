package person

import "github.com/go-msvc/errors"

//Gender ...
type Gender int

const (
	//GenderNotSpecified ...
	GenderNotSpecified Gender = iota
	//GenderMale ...
	GenderMale
	//GenderFemale ...
	GenderFemale
)

//Validate ...
func (g Gender) Validate() error {
	if g != GenderMale && g != GenderFemale {
		return errors.Errorf("expecting male|female")
	}
	return nil
}
