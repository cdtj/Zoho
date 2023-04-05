package books

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

const moduleInvoices = "invoices"

func (c *API) CreateInvoice(inv Invoice) (data *InvoiceResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         moduleInvoices,
		URL:          fmt.Sprintf("https://books.zoho.%s/api/v3/%s", c.ZohoTLD, moduleInvoices),
		Method:       zoho.HTTPPost,
		RequestBody:  inv,
		ResponseData: &InvoiceResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to create the invoice: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*InvoiceResponse); ok {
		return v, nil
	}

	return nil, fmt.Errorf("data retrieved was not 'InvoiceResponse'")
}

type Invoice struct {
	CustomerID            string               `json:"customer_id,omitempty"`
	CurrencyID            string               `json:"currency_id,omitempty"`
	ContactPersons        []string             `json:"contact_persons,omitempty"`
	InvoiceNumber         string               `json:"invoice_number,omitempty"`
	PlaceOfSupply         string               `json:"place_of_supply,omitempty"`
	VATTreatment          string               `json:"vat_treatment,omitempty"`
	TaxTreatment          string               `json:"tax_treatment,omitempty"`
	GSTTreatment          string               `json:"gst_treatment,omitempty"`
	GSTNo                 string               `json:"gst_no,omitempty"`
	CFDIUsage             string               `json:"cfdi_usage,omitempty"`
	ReferenceNumber       string               `json:"reference_number,omitempty"`
	TemplateID            string               `json:"template_id,omitempty"`
	Date                  string               `json:"date,omitempty"`
	PaymentTerms          int                  `json:"payment_terms,omitempty"`
	PaymentTermsLabel     string               `json:"payment_terms_label,omitempty"`
	DueDate               string               `json:"due_date,omitempty"`
	Discount              float64              `json:"discount,omitempty"`
	IsDiscountBeforeTax   bool                 `json:"is_discount_before_tax,omitempty"`
	DiscountType          string               `json:"discount_type,omitempty"`
	IsInclusiveTax        bool                 `json:"is_inclusive_tax,omitempty"`
	ExchangeRate          float64              `json:"exchange_rate,omitempty"`
	RecurringInvoiceID    string               `json:"recurring_invoice_id,omitempty"`
	InvoicedEstimateID    string               `json:"invoiced_estimate_id,omitempty"`
	SalespersonName       string               `json:"salesperson_name,omitempty"`
	CustomFields          []InvoiceCustomField `json:"custom_fields,omitempty"`
	LineItems             []InvoiceLineItem    `json:"line_items,omitempty"`
	PaymentOptions        PaymentOptions       `json:"payment_options,omitempty"`
	AllowPartialPayments  bool                 `json:"allow_partial_payments,omitempty"`
	CustomBody            string               `json:"custom_body,omitempty"`
	CustomSubject         string               `json:"custom_subject,omitempty"`
	Notes                 string               `json:"notes,omitempty"`
	Terms                 string               `json:"terms,omitempty"`
	ShippingCharge        float64              `json:"shipping_charge,omitempty"`
	Adjustment            float64              `json:"adjustment,omitempty"`
	AdjustmentDescription string               `json:"adjustment_description,omitempty"`
	Reason                string               `json:"reason,omitempty"`
	TaxAuthorityID        string               `json:"tax_authority_id,omitempty"`
	TaxExemptionID        string               `json:"tax_exemption_id,omitempty"`
	AvataxUseCode         string               `json:"avatax_use_code,omitempty"`
	AvataxExemptNo        string               `json:"avatax_exempt_no,omitempty"`
	TaxID                 string               `json:"tax_id,omitempty"`
	ExpenseID             string               `json:"expense_id,omitempty"`
	SalesorderItemID      string               `json:"salesorder_item_id,omitempty"`
	AvataxTaxCode         string               `json:"avatax_tax_code,omitempty"`
	TimeEntryIDs          []string             `json:"time_entry_ids,omitempty"`
}

type InvoiceCustomField struct {
	CustomfieldID string `json:"customfield_id,omitempty"`
	Value         string `json:"value,omitempty"`
}

type PaymentOptions struct {
	PaymentGateways []struct {
		Configured       bool   `json:"configured,omitempty"`
		AdditionalField1 string `json:"additional_field1,omitempty"`
		GatewayName      string `json:"gateway_name,omitempty"`
	} `json:"payment_gateways,omitempty"`
}

type InvoiceLineItem struct {
	ItemID             string   `json:"item_id,omitempty"`
	ProjectID          string   `json:"project_id,omitempty"`
	TimeEntryIDs       []string `json:"time_entry_ids,omitempty"`
	ProductType        string   `json:"product_type,omitempty"`
	HSNOrSAC           int      `json:"hsn_or_sac,omitempty"`
	SATItemKeyCode     int      `json:"sat_item_key_code,omitempty"`
	UnitKeyCode        string   `json:"unitkey_code,omitempty"`
	WarehouseID        string   `json:"warehouse_id,omitempty"`
	ExpenseID          string   `json:"expense_id,omitempty"`
	ExpenseReceiptName string   `json:"expense_receipt_name,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
	ItemOrder          int      `json:"item_order,omitempty"`
	BCYRate            float64  `json:"bcy_rate,omitempty"`
	Rate               float64  `json:"rate,omitempty"`
	Quantity           float64  `json:"quantity,omitempty"`
	Unit               string   `json:"unit,omitempty"`
	DiscountAmount     float64  `json:"discount_amount,omitempty"`
	Discount           float64  `json:"discount,omitempty"`
	Tags               []Tag    `json:"tags,omitempty"`
	TaxID              string   `json:"tax_id,omitempty"`
	TDSTaxID           string   `json:"tds_tax_id,omitempty"`
	TaxName            string   `json:"tax_name,omitempty"`
	TaxType            string   `json:"tax_type,omitempty"`
	TaxPercentage      float64  `json:"tax_percentage,omitempty"`
	TaxTreatmentCode   string   `json:"tax_treatment_code,omitempty"`
	HeaderName         string   `json:"header_name,omitempty"`
}

type Tag struct {
	IsTagMandatory bool   `json:"is_tag_mandatory,omitempty"`
	TagID          string `json:"tag_id,omitempty"`
	TagName        string `json:"tag_name,omitempty"`
	TagOptionID    string `json:"tag_option_id,omitempty"`
	TagOptionName  string `json:"tag_option_name,omitempty"`
}

type InvoiceResponse struct {
	Invoice InvoiceResponseDetails `json:"invoice,omitempty"`
	ApiResponse
}

type InvoiceResponseDetails struct {
	InvoiceID              string                    `json:"invoice_id,omitempty"`
	ACHPaymentInitiated    bool                      `json:"ach_payment_initiated,omitempty"`
	InvoiceNumber          string                    `json:"invoice_number,omitempty"`
	IsPreGST               bool                      `json:"is_pre_gst,omitempty"`
	PlaceOfSupply          string                    `json:"place_of_supply,omitempty"`
	GSTNo                  string                    `json:"gst_no,omitempty"`
	GSTTreatment           string                    `json:"gst_treatment,omitempty"`
	CFDIUsage              string                    `json:"cfdi_usage,omitempty"`
	CFDIReferenceType      string                    `json:"cfdi_reference_type,omitempty"`
	ReferenceInvoiceID     string                    `json:"reference_invoice_id,omitempty"`
	VATTreatment           string                    `json:"vat_treatment,omitempty"`
	TaxTreatment           string                    `json:"tax_treatment,omitempty"`
	VATRegNo               string                    `json:"vat_reg_no,omitempty"`
	Date                   string                    `json:"date,omitempty"`
	Status                 string                    `json:"status,omitempty"`
	PaymentTerms           int                       `json:"payment_terms,omitempty"`
	PaymentTermsLabel      string                    `json:"payment_terms_label,omitempty"`
	DueDate                string                    `json:"due_date,omitempty"`
	PaymentExpectedDate    string                    `json:"payment_expected_date,omitempty"`
	LastPaymentDate        string                    `json:"last_payment_date,omitempty"`
	ReferenceNumber        string                    `json:"reference_number,omitempty"`
	CustomerID             string                    `json:"customer_id,omitempty"`
	CustomerName           string                    `json:"customer_name,omitempty"`
	ContactPersons         []string                  `json:"contact_persons,omitempty"`
	CurrencyID             string                    `json:"currency_id,omitempty"`
	CurrencyCode           string                    `json:"currency_code,omitempty"`
	ExchangeRate           float64                   `json:"exchange_rate,omitempty"`
	Discount               float64                   `json:"discount,omitempty"`
	IsDiscountBeforeTax    bool                      `json:"is_discount_before_tax,omitempty"`
	DiscountType           string                    `json:"discount_type,omitempty"`
	IsInclusiveTax         bool                      `json:"is_inclusive_tax,omitempty"`
	RecurringInvoiceID     string                    `json:"recurring_invoice_id,omitempty"`
	IsViewedByClient       bool                      `json:"is_viewed_by_client,omitempty"`
	HasAttachment          bool                      `json:"has_attachment,omitempty"`
	ClientViewedTime       string                    `json:"client_viewed_time,omitempty"`
	LineItems              []InvoiceResponseLineItem `json:"line_items,omitempty"`
	PaymentReminderEnabled bool                      `json:"payment_reminder_enabled,omitempty"`
	PaymentMade            float64                   `json:"payment_made,omitempty"`
	CreditsApplied         float64                   `json:"credits_applied,omitempty"`
	TaxAmountWithheld      float64                   `json:"tax_amount_withheld,omitempty"`
	Balance                float64                   `json:"balance,omitempty"`
	WriteOffAmount         float64                   `json:"write_off_amount,omitempty"`
	AllowPartialPayments   bool                      `json:"allow_partial_payments,omitempty"`
	PricePrecision         int                       `json:"price_precision,omitempty"`
	PaymentOptions         struct {
		PaymentGateways []struct {
			Configured       bool   `json:"configured,omitempty"`
			AdditionalField1 string `json:"additional_field1,omitempty"`
			GatewayName      string `json:"gateway_name,omitempty"`
		} `json:"payment_gateways,omitempty"`
	} `json:"payment_options,omitempty"`
	IsEmailed            bool    `json:"is_emailed,omitempty"`
	RemindersSent        int     `json:"reminders_sent,omitempty"`
	LastReminderSentDate string  `json:"last_reminder_sent_date,omitempty"`
	BillingAddress       Address `json:"billing_address,omitempty"`
	ShippingAddress      Address `json:"shipping_address,omitempty"`
	Notes                string  `json:"notes,omitempty"`
	Terms                string  `json:"terms,omitempty"`
	CustomFields         []struct {
		CustomfieldID string `json:"customfield_id,omitempty"`
		Value         string `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
	TemplateID       string `json:"template_id,omitempty"`
	TemplateName     string `json:"template_name,omitempty"`
	CreatedTime      string `json:"created_time,omitempty"`
	LastModifiedTime string `json:"last_modified_time,omitempty"`
	AttachmentName   string `json:"attachment_name,omitempty"`
	CanSendInMail    bool   `json:"can_send_in_mail,omitempty"`
	SalespersonID    string `json:"salesperson_id,omitempty"`
	SalespersonName  string `json:"salesperson_name,omitempty"`
	InvoiceURL       string `json:"invoice_url,omitempty"`
}

type Address struct {
	Attention string `json:"attention,omitempty"`
	Address   string `json:"address,omitempty"`
	Street2   string `json:"street2,omitempty"`
	StateCode string `json:"state_code,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Country   string `json:"country,omitempty"`
	Fax       string `json:"fax,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type InvoiceResponseLineItem struct {
	LineItemID         string        `json:"line_item_id,omitempty"`
	ItemID             string        `json:"item_id,omitempty"`
	ProjectID          string        `json:"project_id,omitempty"`
	SatItemKeyCode     int           `json:"sat_item_key_code,omitempty"`
	UnitKeyCode        string        `json:"unitkey_code,omitempty"`
	ProjectName        string        `json:"project_name,omitempty"`
	TimeEntryIDs       []interface{} `json:"time_entry_ids,omitempty"`
	Warehouses         []Warehouse   `json:"warehouses,omitempty"`
	ItemType           string        `json:"item_type,omitempty"`
	ProductType        string        `json:"product_type,omitempty"`
	ExpenseID          string        `json:"expense_id,omitempty"`
	ExpenseReceiptName string        `json:"expense_receipt_name,omitempty"`
	Name               string        `json:"name,omitempty"`
	Description        string        `json:"description,omitempty"`
	ItemOrder          int           `json:"item_order,omitempty"`
	BCYRate            float64       `json:"bcy_rate,omitempty"`
	Rate               float64       `json:"rate,omitempty"`
	Quantity           float64       `json:"quantity,omitempty"`
	Unit               string        `json:"unit,omitempty"`
	DiscountAmount     float64       `json:"discount_amount,omitempty"`
	Discount           float64       `json:"discount,omitempty"`
	Tags               []Tag         `json:"tags,omitempty"`
	TaxID              string        `json:"tax_id,omitempty"`
	TDSTaxID           string        `json:"tds_tax_id,omitempty"`
	TaxName            string        `json:"tax_name,omitempty"`
	TaxType            string        `json:"tax_type,omitempty"`
	TaxPercentage      float64       `json:"tax_percentage,omitempty"`
	TaxTreatmentCode   string        `json:"tax_treatment_code,omitempty"`
	ItemTotal          float64       `json:"item_total,omitempty"`
	HeaderName         string        `json:"header_name,omitempty"`
	HeaderID           string        `json:"header_id,omitempty"`
}
