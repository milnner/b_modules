package datacheck

import (
	"regexp"
	"strings"

	apptypes "github.com/milnner/b_modules/apptypes"
	errapp "github.com/milnner/b_modules/errors"
)

func CheckName(x string) error {
	if len(x) > 50 {
		return errapp.NewNameLengthOutOfBoundError()
	} else if len(x) < 1 {
		return errapp.NewNameBelowLengthLimit()
	}
	return nil

}

func CheckSurname(x string) error {
	if len(x) <= 50 {
		return nil
	}
	return errapp.NewSurnameLengthOutOfBoundError()
}

func CheckEmail(x string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)

	if !(re.MatchString(x)) {
		return errapp.NewEmailFormatError()
	}
	return nil
}

func CheckSex(x apptypes.Sex) error {
	exist_type := x == apptypes.Sex(apptypes.Sexs{}.Male()) || x == apptypes.Sex(apptypes.Sexs{}.Female()) || x == apptypes.Sex(apptypes.Sexs{}.Other())

	if exist_type {
		return nil
	}
	return errapp.NewUnknownTypeOfSexError()
}

func CheckPassword(x string) error {
	if len(x) >= 15 {
		return nil
	} else {
		return errapp.NewLengthPasswordUnderTheLimit()
	}
}

func CheckMismatchPassword(x, y string) error {
	if strings.EqualFold(x, y) {
		return nil
	}
	return errapp.NewPasswordMismatchError()
}
