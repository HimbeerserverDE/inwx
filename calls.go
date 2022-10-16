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

// A NSRecordInfoCall returns detailed information about a DNS record.
type NSRecordInfoCall struct {
	DomainName string `json:"domain,omitempty"`
	DomainID   int    `json:"roId,omitempty"`
	RecordID   int    `json:"recordId,omitempty"`
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Content    string `json:"content,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	Priority   int    `json:"prio,omitempty"`
}

func (c *NSRecordInfoCall) method() string           { return "nameserver.info" }
func (c *NSRecordInfoCall) expectedStatus() []Status { return []Status{Success} }

// A NSUpdateRecordsCall updates one or more DNS records.
type NSUpdateRecordsCall struct {
	IDs                 []int                 `json:"id"`
	Name                string                `json:"name,omitempty"`
	Type                RecordType            `json:"type,omitempty"`
	Content             string                `json:"content,omitempty"`
	Priority            int                   `json:"prio,omitempty"`
	TTL                 int                   `json:"ttl,omitempty"`
	URLRedirectType     RecordURLRedirectType `json:"urlRedirectType,omitempty"`
	URLRedirectTitle    string                `json:"urlRedirectTitle,omitempty"`
	URLRedirectDesc     string                `json:"urlRedirectDescription,omitempty"`
	URLRedirectFavIcon  string                `json:"urlRedirectFavIcon,omitempty"`
	URLRedirectKeywords string                `json:"urlRedirectKeywords,omitempty"`
	URLAppend           bool                  `json:"urlAppend,omitempty"`
	TestingMode         bool                  `json:"testing,omitempty"`
}

func (c *NSUpdateRecordsCall) method() string           { return "nameserver.updateRecord" }
func (c *NSUpdateRecordsCall) expectedStatus() []Status { return []Status{Success} }
