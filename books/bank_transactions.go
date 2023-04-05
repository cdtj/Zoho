package books

type Transaction struct {
	TransactionID   string `json:"transaction_id,omitempty"`
	FromAccountID   string `json:"from_account_id,omitempty"`
	FromAccountName string `json:"from_account_name,omitempty"`
	ToAccountID     string `json:"to_account_id,omitempty"`
	ToAccountName   string `json:"to_account_name,omitempty"`
	TransactionType string `json:"transaction_type,omitempty"`
	CurrencyID      string `json:"currency_id,omitempty"`
	CurrencyCode    string `json:"currency_code,omitempty"`
	PaymentMode     string `json:"payment_mode,omitempty"`
	ExchangeRate    string `json:"exchange_rate,omitempty"`
	Date            string `json:"date,omitempty"`
	CustomerID      string `json:"customer_id,omitempty"`
	CustomerName    string `json:"customer_name,omitempty"`
	VendorID        string `json:"vendor_id,omitempty"`
	VendorName      string `json:"vendor_name,omitempty"`
	ReferenceNumber string `json:"reference_number,omitempty"`
	Description     string `json:"description,omitempty"`
	BankCharges     string `json:"bank_charges,omitempty"`
	TaxID           string `json:"tax_id,omitempty"`
	Documents       []struct {
		FileName   interface{} `json:"file_name,omitempty"`
		DocumentID interface{} `json:"document_id,omitempty"`
	} `json:"documents,omitempty"`
	IsInclusiveTax             bool   `json:"is_inclusive_tax,omitempty"`
	TaxName                    string `json:"tax_name,omitempty"`
	TaxPercentage              string `json:"tax_percentage,omitempty"`
	TaxAmount                  string `json:"tax_amount,omitempty"`
	SubTotal                   string `json:"sub_total,omitempty"`
	TaxAuthorityID             string `json:"tax_authority_id,omitempty"`
	TaxAuthorityName           string `json:"tax_authority_name,omitempty"`
	TaxExemptionID             string `json:"tax_exemption_id,omitempty"`
	TaxExemptionCode           string `json:"tax_exemption_code,omitempty"`
	Total                      string `json:"total,omitempty"`
	BcyTotal                   string `json:"bcy_total,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	VATTreatment               string `json:"vat_treatment,omitempty"`
	ProductType                string `json:"product_type,omitempty"`
	AcquisitionVatID           string `json:"acquisition_vat_id,omitempty"`
	AcquisitionVatName         string `json:"acquisition_vat_name,omitempty"`
	AcquisitionVatPercentage   string `json:"acquisition_vat_percentage,omitempty"`
	AcquisitionVatAmount       string `json:"acquisition_vat_amount,omitempty"`
	ReverseChargeVatID         string `json:"reverse_charge_vat_id,omitempty"`
	ReverseChargeVatName       string `json:"reverse_charge_vat_name,omitempty"`
	ReverseChargeVatPercentage string `json:"reverse_charge_vat_percentage,omitempty"`
	ReverseChargeVatAmount     string `json:"reverse_charge_vat_amount,omitempty"`
	FiledInVatReturnID         string `json:"filed_in_vat_return_id,omitempty"`
	FiledInVatReturnName       string `json:"filed_in_vat_return_name,omitempty"`
	FiledInVatReturnType       string `json:"filed_in_vat_return_type,omitempty"`
	ImportedTransactions       []ImportedTransaction
	Tags                       []Tag                     `json:"tags,omitempty"`
	LineItems                  []BankTransactionLineItem `json:"line_items,omitempty"`
}

type ImportedTransaction struct {
	ImportedTransactionID string `json:"imported_transaction_id,omitempty"`
	Date                  string `json:"date,omitempty"`
	Amount                string `json:"amount,omitempty"`
	Payee                 string `json:"payee,omitempty"`
	Description           string `json:"description,omitempty"`
	ReferenceNumber       string `json:"reference_number,omitempty"`
	Status                string `json:"status,omitempty"`
	AccountID             string `json:"account_id,omitempty"`
}

type BankTransactionLineItem struct {
	FromAccountID   string  `json:"from_account_id"`
	FromAccountName string  `json:"from_account_name"`
	PaymentMode     string  `json:"payment_mode"`
	CustomerID      string  `json:"customer_id"`
	CustomerName    string  `json:"customer_name"`
	VendorID        string  `json:"vendor_id"`
	VendorName      string  `json:"vendor_name"`
	SubTotal        float64 `json:"sub_total"`
	Total           float64 `json:"total"`
	BCYTotal        float64 `json:"bcy_total"`
	Tags            []Tag   `json:"tags"`
}
