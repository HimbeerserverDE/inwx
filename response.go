package inwx

import "encoding/json"

// A Response contains the Status of a response to a Call
// as well as the response parameters.
// Data is nil if there are no response parameters.
type Response struct {
	StatusCode Status           `json:"code"`
	Data       *json.RawMessage `json:"resData"`
}

// Into unmarshals the Response into a Call specific response struct.
func (r *Response) Into(resp any) error {
	return json.Unmarshal(*r.Data, resp)
}
