package inwx

// A Status contains basic information
// about whether and how a Call was processed.
type Status int

const (
	Success              Status = 1000
	SuccessPending              = 1001
	SuccessNoMsgs               = 1300
	SuccessDequeue              = 1301
	SuccessClosing              = 1500
	WrongPasswd                 = 1200
	UnknownCall                 = 2000
	SyntaxErr                   = 2001
	UseErr                      = 2002
	MissingParam                = 2003
	WrongParamRange             = 2004
	ParamSyntaxErr              = 2005
	UnimplementedCall           = 2101
	UnimplementedOpt            = 2102
	UnimplementedExt            = 2103
	BillingErr                  = 2104
	DenyRenewal                 = 2105
	DenyTransfer                = 2106
	AuthError                   = 2200
	AuthzError                  = 2201
	InvalidAuthzInfo            = 2202
	PendingTransfer             = 2300
	NotPendingTransfer          = 2301
	Exists                      = 2302
	NotExist                    = 2303
	StatusProhibitsOp           = 2304
	AssocProhibitsOp            = 2305
	ParamPolicyErr              = 2306
	UnimplementedService        = 2307
	DataMgmtPolicyErr           = 2308
	CmdErr                      = 2400
	CmdErrClosing               = 2500
	AuthErrClosing              = 2501
	SessionOverClosing          = 2502
)
