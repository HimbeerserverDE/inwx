package inwx

import (
	"context"
	"fmt"

	"github.com/ybbus/jsonrpc/v3"
)

type Call interface {
	method() string
	expectedStatus() []Status
}

type Response struct {
	StatusCode Status
	Data       any
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
	resp, err := c.rpcClient.Call(context.Background(), call.method(), call)
	if err != nil {
		return nil, err
	}

	meta, ok := resp.Result.(map[string]any)
	if !ok {
		return nil, &ErrInvalidResponse{resp}
	}

	code := meta["code"].(Status)
	data := meta["resData"]

	var expectedStatus bool
	for _, expected := range call.expectedStatus() {
		if expected == code {
			expectedStatus = true
			break
		}
	}

	if !expectedStatus {
		return nil, &ErrUnexpectedStatus{call.expectedStatus(), code}
	}

	response := &Response{
		StatusCode: code,
		Data:       data,
	}

	return response, nil
}
