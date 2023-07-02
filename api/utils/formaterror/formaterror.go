package formaterror

import (
	"errors"
	"strings"
)

var errMapping = map[string]string{
	"nickname":       "Nickname Already Taken",
	"email":          "Email Already Taken",
	"title":          "Title Already Taken",
	"hashedPassword": "Incorrect Password",
}

func FormatError(err string) error {
	for key, val := range errMapping {
		if strings.Contains(err, key) {
			return errors.New(val)
		}
	}

	return errors.New("Incorrect Details")
}
