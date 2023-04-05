package books

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

const moduleContacts = "contacts"

// GetCurrentUser will return the currently authenticated users
// https://www.zoho.com/books/api/v3/users/#get-current-user
func (c *API) GetContacts() (data *ListContactsResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         moduleContacts,
		URL:          fmt.Sprintf("https://books.zoho.%s/api/v3/%s", c.ZohoTLD, moduleContacts),
		Method:       zoho.HTTPGet,
		ResponseData: &ListContactsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current user: %s", err)
	}

	if v, ok := endpoint.ResponseData.(*ListContactsResponse); ok {
		return v, nil
	}

	return nil, fmt.Errorf("data retrieved was not 'ListContactsResponse'")
}

type Contact struct {
	ContactID                        string               `json:"contact_id,omitempty"`
	ContactName                      string               `json:"contact_name,omitempty"`
	CompanyName                      string               `json:"company_name,omitempty"`
	HasTransaction                   bool                 `json:"has_transaction,omitempty"`
	ContactType                      string               `json:"contact_type,omitempty"`
	CustomerSubType                  string               `json:"customer_sub_type,omitempty"`
	CreditLimit                      float64              `json:"credit_limit,omitempty"`
	IsPortalEnabled                  bool                 `json:"is_portal_enabled,omitempty"`
	LanguageCode                     string               `json:"language_code,omitempty"`
	IsTaxable                        bool                 `json:"is_taxable,omitempty"`
	TaxID                            string               `json:"tax_id,omitempty"`
	TDSTaxID                         string               `json:"tds_tax_id,omitempty"`
	TaxName                          string               `json:"tax_name,omitempty"`
	TaxPercentage                    float64              `json:"tax_percentage,omitempty"`
	TaxAuthorityID                   string               `json:"tax_authority_id,omitempty"`
	TaxExemptionID                   string               `json:"tax_exemption_id,omitempty"`
	TaxAuthorityName                 string               `json:"tax_authority_name,omitempty"`
	TaxExemptionCode                 string               `json:"tax_exemption_code,omitempty"`
	PlaceOfContact                   string               `json:"place_of_contact,omitempty"`
	GSTNo                            string               `json:"gst_no,omitempty"`
	VATTreatment                     string               `json:"vat_treatment,omitempty"`
	TaxTreatment                     string               `json:"tax_treatment,omitempty"`
	TaxRegime                        string               `json:"tax_regime,omitempty"`
	IsTDSRegistered                  bool                 `json:"is_tds_registered,omitempty"`
	GSTTreatment                     string               `json:"gst_treatment,omitempty"`
	IsLinkedWithZohoCRM              bool                 `json:"is_linked_with_zohocrm,omitempty"`
	Website                          string               `json:"website,omitempty"`
	OwnerID                          string               `json:"owner_id,omitempty"`
	PrimaryContactID                 string               `json:"primary_contact_id,omitempty"`
	PaymentTerms                     int                  `json:"payment_terms,omitempty"`
	PaymentTermsLabel                string               `json:"payment_terms_label,omitempty"`
	CurrencyID                       string               `json:"currency_id,omitempty"`
	CurrencyCode                     string               `json:"currency_code,omitempty"`
	CurrencySymbol                   string               `json:"currency_symbol,omitempty"`
	OpeningBalanceAmount             float64              `json:"opening_balance_amount,omitempty"`
	ExchangeRate                     float64              `json:"exchange_rate,omitempty"`
	OutstandingReceivableAmount      float64              `json:"outstanding_receivable_amount,omitempty"`
	OutstandingReceivableAmountBCY   float64              `json:"outstanding_receivable_amount_bcy,omitempty"`
	UnusedCreditsReceivableAmount    float64              `json:"unused_credits_receivable_amount,omitempty"`
	UnusedCreditsReceivableAmountBCY float64              `json:"unused_credits_receivable_amount_bcy,omitempty"`
	Status                           string               `json:"status,omitempty"`
	PaymentReminderEnabled           bool                 `json:"payment_reminder_enabled,omitempty"`
	CustomFields                     []ContactCustomField `json:"custom_fields,omitempty"`
	BillingAddress                   Address              `json:"billing_address,omitempty"`
	ShippingAddress                  Address              `json:"shipping_address,omitempty"`
	Facebook                         string               `json:"facebook,omitempty"`
	Twitter                          string               `json:"twitter,omitempty"`
	ContactPersons                   []ContactPerson      `json:"contact_persons,omitempty"`
	DefaultTemplates                 DefaultTemplates     `json:"default_templates,omitempty"`
	Notes                            string               `json:"notes,omitempty"`
	CreatedTime                      string               `json:"created_time,omitempty"`
	LastModifiedTime                 string               `json:"last_modified_time,omitempty"`
}

type ContactCustomField struct {
	Index int    `json:"index,omitempty"`
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

type ContactPerson struct {
	ContactPersonID  string `json:"contact_person_id,omitempty"`
	Salutation       string `json:"salutation,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
	Designation      string `json:"designation,omitempty"`
	Department       string `json:"department,omitempty"`
	Skype            string `json:"skype,omitempty"`
	IsPrimaryContact bool   `json:"is_primary_contact,omitempty"`
	EnablePortal     bool   `json:"enable_portal,omitempty"`
}

type DefaultTemplates struct {
	InvoiceTemplateID                             string `json:"invoice_template_id,omitempty"`
	EstimateTemplateID                            string `json:"estimate_template_id,omitempty"`
	CreditNoteTemplateID                          string `json:"creditnote_template_id,omitempty"`
	PurchaseOrderTemplateID                       string `json:"purchaseorder_template_id,omitempty"`
	SalesOrderTemplateID                          string `json:"salesorder_template_id,omitempty"`
	RetainerInvoiceTemplateID                     string `json:"retainerinvoice_template_id,omitempty"`
	PaymentThankyouTemplateID                     string `json:"paymentthankyou_template_id,omitempty"`
	RetainerInvoicePaymentThankyouTemplateID      string `json:"retainerinvoice_paymentthankyou_template_id,omitempty"`
	InvoiceEmailTemplateID                        string `json:"invoice_email_template_id,omitempty"`
	EstimateEmailTemplateID                       string `json:"estimate_email_template_id,omitempty"`
	CreditNoteEmailTemplateID                     string `json:"creditnote_email_template_id,omitempty"`
	PurchaseOrderEmailTemplateID                  string `json:"purchaseorder_email_template_id,omitempty"`
	SalesOrderEmailTemplateID                     string `json:"salesorder_email_template_id,omitempty"`
	RetainerInvoiceEmailTemplateID                string `json:"retainerinvoice_email_template_id,omitempty"`
	PaymentThankyouEmailTemplateID                string `json:"paymentthankyou_email_template_id,omitempty"`
	RetainerInvoicePaymentThankyouEmailTemplateID string `json:"retainerinvoice_paymentthankyou_email_template_id,omitempty"`
}

type ListContactsResponse struct {
	Contacts []Contact `json:"contacts,omitempty"`
	ApiResponse
}
