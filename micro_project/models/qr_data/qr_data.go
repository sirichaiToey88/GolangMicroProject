package qrdata

type TransferQrField struct {
	FieldNumber string
	FieldValue  string
	FieldLength int
}

type TransferQrData struct {
	Version           string
	Type              string
	TransferType      string
	Tag               string
	Merchant          string
	MerchantCity      string
	MerchantName      string
	County            string
	CurrencyCode      string
	Amount            float64
	CheckSum          string
	QRAdditionalField *TransferQrAdditionalField
	IsValid           bool
}

type TransferQrPromptPay struct {
	IsValid                bool
	Version                string
	Type                   string
	TransferType           string
	Tag                    string
	BankCode               string
	Amount                 float64
	MerchantCity           string
	MerchantName           string
	Merchant               string
	Aid                    string
	DestinationAccount     string
	Ota                    string
	TypeDestinationAccount string
	QRAdditionalField      *TransferQrAdditionalField
}

type TransferQrBillPayment struct {
	IsValid           bool
	Version           string
	Type              string
	TransferType      string
	Tag               string
	BillerId          string
	Amount            float64
	Reference1        string
	Reference2        string
	MerchantName      string
	Merchant          string
	Aid               string
	QRAdditionalField *TransferQrAdditionalField
}

type TransferQrVerifySlip struct {
	IsValid        bool
	Version        string
	Type           string
	TransferType   string
	Tag            string
	ApiId          string
	BankId         string
	TransactionRef string
	TransferQrData
}

type TransferQrAdditionalField struct {
	AdditionalBillNumber          string
	AdditionalMobileNumber        string
	AdditionalStoreID             string
	AdditionalLoyaltyNumber       string
	AdditionalReferenceID         string
	AdditionalConsumerID          string
	AdditionalTerminalID          string
	AdditionalPurposeTransaction  string
	AdditionalPursposeTransaction string
	AdditionalConsumerDataRequest string
}
