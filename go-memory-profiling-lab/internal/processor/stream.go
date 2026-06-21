package processor

import (
	"encoding/json"
	"fmt"
	"io"
)

// InputUser and OutputUser are used by the streaming examples.
type InputUser struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
	Name   string `json:"name"`
}

type OutputUser struct {
	ID string `json:"id"`
}

// FilterActiveUsersInMemory decodes the full input array and builds the full
// output slice before writing the response.
func FilterActiveUsersInMemory(r io.Reader, w io.Writer) error {
	var users []InputUser
	if err := json.NewDecoder(r).Decode(&users); err != nil {
		return err
	}

	active := make([]OutputUser, 0, len(users))
	for _, user := range users {
		if user.Active {
			active = append(active, OutputUser{ID: user.ID})
		}
	}

	return json.NewEncoder(w).Encode(active)
}

// FilterActiveUsersStreaming reads one JSON array element at a time and writes
// matching output records progressively.
func FilterActiveUsersStreaming(r io.Reader, w io.Writer) error {
	dec := json.NewDecoder(r)

	token, err := dec.Token()
	if err != nil {
		return err
	}
	if delimiter, ok := token.(json.Delim); !ok || delimiter != '[' {
		return fmt.Errorf("expected JSON array")
	}

	if _, err := io.WriteString(w, "["); err != nil {
		return err
	}

	first := true
	enc := json.NewEncoder(w)

	for dec.More() {
		var user InputUser
		if err := dec.Decode(&user); err != nil {
			return err
		}

		if !user.Active {
			continue
		}

		if !first {
			if _, err := io.WriteString(w, ","); err != nil {
				return err
			}
		}
		first = false

		if err := enc.Encode(OutputUser{ID: user.ID}); err != nil {
			return err
		}
	}

	token, err = dec.Token()
	if err != nil {
		return err
	}
	if delimiter, ok := token.(json.Delim); !ok || delimiter != ']' {
		return fmt.Errorf("expected end of JSON array")
	}

	_, err = io.WriteString(w, "]")
	return err
}
