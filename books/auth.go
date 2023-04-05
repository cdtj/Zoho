package books

import zoho "github.com/schmorrison/Zoho"

const (
	ServiceBooks zoho.Service = "ZohoBooks"

	ScopeContacts         zoho.Scope = "contacts"
	ScopeSettings         zoho.Scope = "settings"
	ScopeEstimates        zoho.Scope = "estimates"
	ScopeInvoices         zoho.Scope = "invoices"
	ScopeCustomerpayments zoho.Scope = "customerpayments"
	ScopeCreditnotes      zoho.Scope = "creditnotes"
	ScopeProjects         zoho.Scope = "projects"
	ScopeExpenses         zoho.Scope = "expenses"
	ScopeSalesorder       zoho.Scope = "salesorder"
	ScopePurchaseorder    zoho.Scope = "purchaseorder"
	ScopeBills            zoho.Scope = "bills"
	ScopeDebitnotes       zoho.Scope = "debitnotes"
	ScopeVendorpayments   zoho.Scope = "vendorpayments"
	ScopeBanking          zoho.Scope = "banking"
	ScopeAccountants      zoho.Scope = "accountants"
	ScopeUsers            zoho.Scope = "users"

	MethodNoOp zoho.Method = ""

	OpRead   zoho.Operation = "READ"
	OpDelete zoho.Operation = "DELETE"
	OpUpdate zoho.Operation = "UPDATE"
	OpCreate zoho.Operation = "CREATE"
)
