package inwx

type loginCall struct {
	User            string `json:"user"`
	Passwd          string `json:"pass"`
	CaseInsensitive bool   `json:"case-insensitive,omitempty"`
}

func (c *loginCall) method() string           { return "account.login" }
func (c *loginCall) expectedStatus() []Status { return []Status{Success} }

type logoutCall struct{}

func (c *logoutCall) method() string           { return "account.logout" }
func (c *logoutCall) expectedStatus() []Status { return []Status{SuccessClosing} }
