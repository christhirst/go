package db

import (
	"encoding/json"
	"io"
)

func (a *Account) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func (a *Account) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)

}
