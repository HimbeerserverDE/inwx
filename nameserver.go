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
	Name                string     `json:"name"`
	Type                RecordType `json:"type"`
	Content             string     `json:"content"`
	TTL                 int        `json:"ttl"`
	Priority            int        `json:"prio"`
	URLRedirectType     URLRdrType `json:"urlRedirectType"`
	URLRedirectTitle    string     `json:"urlRedirectTitle"`
	URLRedirectDesc     string     `json:"urlRedirectDescription"`
	URLRedirectKeywords string     `json:"urlRedirectKeywords"`
	URLRedirectFavIcon  string     `json:"urlRedirectFavIcon"`
	URLAppend           bool       `json:"urlAppend"`
}

// A NSRecord consists of an ID and RecordInfo.
type NSRecord struct {
	ID int `json:"id"`
	RecordInfo
}
