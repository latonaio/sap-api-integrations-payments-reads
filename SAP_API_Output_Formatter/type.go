package sap_api_output_formatter

type Payments struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	Payments      string `json:"payments_code"`
	Deleted       bool   `json:"deleted"`
}

type PaymentCollection struct {
	EntityLastChangedOn                string      `json:"EntityLastChangedOn"`     
	ETag                               string      `json:"ETag"`
	AccountPartyID                     string      `json:"AccountPartyID"`
	AccountPartyName                   string      `json:"AccountPartyName"`
	Amount                             string      `json:"Amount"`
	AmountCurrencyCode                 string      `json:"AmountCurrencyCode"`
	BankName                           string      `json:"BankName"`
	ChequeDate                         string      `json:"ChequeDate"`
	ChequeNumber                       string      `json:"ChequeNumber"`
	CreationDate                       string      `json:"CreationDate"`
	EmployeeResponsible                string      `json:"EmployeeResponsible"`
	EmployeeResposibleID               string      `json:"EmployeeResposibleID"`
	ID                                 string      `json:"ID"`
	InformationLifeCycleStatusCode     string      `json:"InformationLifeCycleStatusCode"`
	LastChangeDate                     string      `json:"LastChangeDate"`
	Name                               string      `json:"Name"`
	ObjectID                           string      `json:"ObjectID"`
	Payer                              string      `json:"Payer"`
	PaymentDate                        string      `json:"PaymentDate"`
	PaymentMode                        string      `json:"PaymentMode"`
	PaymentStatusCode                  string      `json:"PaymentStatusCode"`
	TransferStatusCode                 string      `json:"TransferStatusCode"`
}
