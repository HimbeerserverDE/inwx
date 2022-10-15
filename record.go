package inwx

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

// A RecordURLRedirectType specifies which method of HTTP redirection to use.
type RecordURLRedirectType string

const (
	Permanent = "HEADER301"
	Temporary = "HEADER302"
	Frame     = "FRAME"
)
