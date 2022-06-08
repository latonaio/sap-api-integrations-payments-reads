package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-payments-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToPaymentCollection(raw []byte, l *logger.Logger) ([]PaymentCollection, error) {
	pm := &responses.PaymentCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PaymentCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	paymentCollection := make([]PaymentCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		paymentCollection = append(paymentCollection, PaymentCollection{
			EntityLastChangedOn:                data.EntityLastChangedOn,
			ETag:                               data.ETag,
			AccountPartyID:                     data.AccountPartyID,
			AccountPartyName:                   data.AccountPartyName,
			Amount:                             data.Amount,
			AmountCurrencyCode:                 data.AmountCurrencyCode,
			BankName:                           data.BankName,
			ChequeDate:                         data.ChequeDate,
			ChequeNumber:                       data.ChequeNumber,
			CreationDate:                       data.CreationDate,
			EmployeeResponsible:                data.EmployeeResponsible,
			EmployeeResposibleID:               data.EmployeeResposibleID,
			ID:                                 data.ID,
			InformationLifeCycleStatusCode:     data.InformationLifeCycleStatusCode,
			LastChangeDate:                     data.LastChangeDate,
			Name:                               data.Name,
			ObjectID:                           data.ObjectID,
			Payer:                              data.Payer,
			PaymentDate:                        data.PaymentDate,
			PaymentMode:                        data.PaymentMode,
			PaymentStatusCode:                  data.PaymentStatusCode,
			TransferStatusCode:                 data.TransferStatusCode,
		})
	}

	return paymentCollection, nil
}
