package telegram

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Response holds the information returned from a telegram message call
type Response struct {
	OK          bool   `json:"ok"`          // True if call succeeded
	ErrorCode   int    `json:"error_code"`  // The error code
	Description string `json:"description"` // Error description
}

// ReadFrom will read the request body and deserialize it into the request values
func (s *Response) ReadFrom(r io.ReadCloser) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if b != nil && len(b) != 0 {
		err = json.Unmarshal(b, &s)
		if err != nil {
			return err
		}
	}
	return nil
}
