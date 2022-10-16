package inwx

import "time"

// A NSRecordInfoResponse contains the information that was requested
// about a specific DNS record.
type NSRecordInfoResponse struct {
	DomainID      int       `json:"roId"`
	DomainName    string    `json:"domain"`
	DomainType    NSType    `json:"type"`
	MasterAddr    string    `json:"masterIp"`
	LastZoneCheck time.Time `json:"lastZoneCheck"`
	SlaveDNS      SlaveInfo `json:"slaveDns"`
	SOASerial     string    `json:"SOAserial"`
	RecordCount   int       `json:"count"`
	Record        NSRecord  `json:"record"`
}
