package books

type BankAccount struct {
	AccountID                    string  `json:"account_id,omitempty"`
	AccountName                  string  `json:"account_name,omitempty"`
	AccountCode                  string  `json:"account_code,omitempty"`
	CurrencyID                   string  `json:"currency_id,omitempty"`
	CurrencyCode                 string  `json:"currency_code,omitempty"`
	CurrencySymbol               string  `json:"currency_symbol,omitempty"`
	PricePrecision               int     `json:"price_precision,omitempty"`
	AccountType                  string  `json:"account_type,omitempty"`
	AccountNumber                string  `json:"account_number,omitempty"`
	UncategorizedTransactions    int     `json:"uncategorized_transactions,omitempty"`
	TotalUnprintedChecks         int     `json:"total_unprinted_checks,omitempty"`
	IsActive                     bool    `json:"is_active,omitempty"`
	IsFeedsSubscribed            bool    `json:"is_feeds_subscribed,omitempty"`
	IsFeedsActive                bool    `json:"is_feeds_active,omitempty"`
	Balance                      float64 `json:"balance,omitempty"`
	BankBalance                  float64 `json:"bank_balance,omitempty"`
	BCYBalance                   float64 `json:"bcy_balance,omitempty"`
	BankName                     string  `json:"bank_name,omitempty"`
	RoutingNumber                string  `json:"routing_number,omitempty"`
	IsPrimaryAccount             bool    `json:"is_primary_account,omitempty"`
	IsPaypalAccount              bool    `json:"is_paypal_account,omitempty"`
	Description                  string  `json:"description,omitempty"`
	RefreshStatusCode            string  `json:"refresh_status_code,omitempty"`
	FeedsLastRefreshDate         string  `json:"feeds_last_refresh_date,omitempty"`
	ServiceID                    string  `json:"service_id,omitempty"`
	IsSystemAccount              bool    `json:"is_system_account,omitempty"`
	IsShowWarningForFeedsRefresh bool    `json:"is_show_warning_for_feeds_refresh,omitempty"`
}

type ListBankAccount struct {
	BankAccount BankAccount `json:"bankaccount"`
	ApiResponse
}

type ListBankAccounts struct {
	BankAccount []BankAccount `json:"bankaccounts"`
	ApiResponse
}
