package inwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type rpcCall struct {
	Method string `json:"method"`
	Params any    `json:"params,omitempty"`
}

// A Call is anything that provides a method name
// and a list of status codes that indicate basic success.
// It can be executed through a Client.
type Call interface {
	method() string
	expectedStatus() []Status
}

// A Response contains the Status of a response to a Call
// as well as the response parameters.
// Data is nil if there are no response parameters.
type Response struct {
	StatusCode Status `json:"code"`
	Data       any    `json:"resData"`
}

// ErrUnexpectedStatus indicates that a Call was responded to
// but the Status doesn't match any of the expected Statuses of the Call.
type ErrUnexpectedStatus struct {
	expected []Status
	got      Status
}

func (e *ErrUnexpectedStatus) Error() string {
	const format = "unexpected status code: expected %v, got %v"
	return fmt.Sprintf(format, e.expected, e.got)
}

// Call sends a Call to the API endpoint and waits for a response or an error.
// It returns net.ErrClosed if the Client is closed.
func (c *Client) Call(call Call) (*Response, error) {
	wrapped := &rpcCall{
		Method: call.method(),
		Params: call,
	}

	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(wrapped); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", string(c.Endpoint()), body)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	response := &Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	var expectedStatus bool
	for _, expected := range call.expectedStatus() {
		if expected == response.StatusCode {
			expectedStatus = true
			break
		}
	}

	if !expectedStatus {
		return nil, &ErrUnexpectedStatus{call.expectedStatus(), response.StatusCode}
	}

	return response, nil
}
