package inwx

import "time"

// A NSRecordInfoResponse contains all NSRecords that match the search.
type NSRecordInfoResponse struct {
	DomainID      int        `json:"roId"`
	DomainName    string     `json:"domain"`
	DomainType    NSType     `json:"type"`
	MasterAddr    string     `json:"masterIp"`
	LastZoneCheck time.Time  `json:"lastZoneCheck"`
	SlaveDNS      SlaveInfo  `json:"slaveDns"`
	SOASerial     string     `json:"SOAserial"`
	RecordCount   int        `json:"count"`
	Records       []NSRecord `json:"record"`
}
