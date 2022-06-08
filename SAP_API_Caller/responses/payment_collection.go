package responses

type PaymentCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			} `json:"results"`
			} `json:"d"`
		}
