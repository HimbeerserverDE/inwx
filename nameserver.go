package inwx

// A NSType is either Master or Slave.
type NSType string

const (
	Master NSType = "MASTER"
	Slave         = "SLAVE"
)

// SlaveInfo is information about a slave nameserver.
type SlaveInfo struct {
	Hostname string `json:"name"`
	Addr     string `json:"ip"`
}

// A RecordType specifies the type of a DNS record.
type RecordType string

const (
	A          = "A"
	AAAA       = "AAAA"
	AFSDB      = "AFSDB"
	ALIAS      = "ALIAS"
	CAA        = "CAA"
	Cert       = "CERT"
	CNAME      = "CNAME"
	HINFO      = "HINFO"
	KEY        = "KEY"
	LOC        = "LOC"
	MX         = "MX"
	NAPTR      = "NAPTR"
	NS         = "NS"
	OpenPGPKey = "OPENPGPKEY"
	PTR        = "PTR"
	RP         = "RP"
	SMIMEA     = "SMIMEA"
	SOA        = "SOA"
	SRV        = "SRV"
	SSHFP      = "SSHFP"
	TLSA       = "TLSA"
	TXT        = "TXT"
	URI        = "URI"
	URL        = "URL"
)

// A URLRdrType specifies which method of HTTP redirection to use.
type URLRdrType string

const (
	Permanent = "HEADER301"
	Temporary = "HEADER302"
	Frame     = "FRAME"
)

// A RecordInfo contains DNS as well as INWX specific record information.
type RecordInfo struct {
	Name                string     `json:"name,omitempty"`
	Type                RecordType `json:"type,omitempty"`
	Content             string     `json:"content,omitempty"`
	TTL                 int        `json:"ttl,omitempty"`
	Priority            int        `json:"prio,omitempty"`
	URLRedirectType     URLRdrType `json:"urlRedirectType,omitempty"`
	URLRedirectTitle    string     `json:"urlRedirectTitle,omitempty"`
	URLRedirectDesc     string     `json:"urlRedirectDescription,omitempty"`
	URLRedirectKeywords string     `json:"urlRedirectKeywords,omitempty"`
	URLRedirectFavIcon  string     `json:"urlRedirectFavIcon,omitempty"`
	URLAppend           bool       `json:"urlAppend,omitempty"`
}

// A NSRecord consists of an ID and RecordInfo.
type NSRecord struct {
	ID int `json:"id"`
	RecordInfo
}
