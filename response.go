package inwx

// A Response contains the Status of a response to a Call
// as well as the response parameters.
// Data is nil if there are no response parameters.
type Response struct {
	StatusCode Status `json:"code"`
	Data       any    `json:"resData"`
}
