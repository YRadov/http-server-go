package validation

import (
	"regexp"
	"net/http"
	"errors"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error)  {

	m := validPath.FindStringSubmatch(r.URL.Path)

	if m == nil {

		http.NotFound(w, r)

		return "", errors.New("Invalid Page Title")
	}

	return m[2], nil // The title is the second subexpression.

}//GetTitle