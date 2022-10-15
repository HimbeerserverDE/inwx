package inwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ybbus/jsonrpc/v3"
)

type rpcCall struct {
	method string `json:"method"`
	params any    `json:"params,omitempty"`
}

type Call interface {
	method() string
	expectedStatus() []Status
}

type Response struct {
	StatusCode Status `json:"code"`
	Data       any    `json:"resData"`
}

type ErrInvalidResponse struct {
	resp *jsonrpc.RPCResponse
}

func (e *ErrInvalidResponse) Error() string {
	return fmt.Sprintf("invalid API response: %v", e.resp)
}

type ErrUnexpectedStatus struct {
	expected []Status
	got      Status
}

func (e *ErrUnexpectedStatus) Error() string {
	const format = "unexpected status code: expected %v, got %v"
	return fmt.Sprintf(format, e.expected, e.got)
}

func (c *Client) Call(call Call) (*Response, error) {
	wrapped := &rpcCall{
		method: call.method(),
		params: call,
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
