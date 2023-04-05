package books

import (
	"fmt"

	zoho "github.com/schmorrison/Zoho"
)

type Item struct {
	ItemID             string        `json:"item_id,omitempty"`
	Name               string        `json:"name,omitempty"`
	Status             string        `json:"status,omitempty"`
	Description        string        `json:"description,omitempty"`
	Rate               float64       `json:"rate,omitempty"`
	Unit               string        `json:"unit,omitempty"`
	TaxID              string        `json:"tax_id,omitempty"`
	PurchaseTaxRuleID  string        `json:"purchase_tax_rule_id,omitempty"`
	SalesTaxRuleID     string        `json:"sales_tax_rule_id,omitempty"`
	TaxName            string        `json:"tax_name,omitempty"`
	HSNOrSAC           string        `json:"hsn_or_sac,omitempty"`
	SatItemKeyCode     string        `json:"sat_item_key_code,omitempty"`
	UnitKeyCode        string        `json:"unitkey_code,omitempty"`
	TaxPercentage      float64       `json:"tax_percentage,omitempty"`
	TaxType            string        `json:"tax_type,omitempty"`
	SKU                string        `json:"sku,omitempty"`
	ProductType        string        `json:"product_type,omitempty"`
	ItemTaxPreferences []ItemTaxPref `json:"item_tax_preferences,omitempty"`
	Warehouses         []Warehouse   `json:"warehouses,omitempty"`
}

type ItemTaxPref struct {
	TaxID            string `json:"tax_id,omitempty"`
	TaxSpecification string `json:"tax_specification,omitempty"`
}

type Warehouse struct {
	WarehouseID                   string `json:"warehouse_id,omitempty"`
	WarehouseName                 string `json:"warehouse_name,omitempty"`
	Status                        string `json:"status,omitempty"`
	IsPrimary                     bool   `json:"is_primary,omitempty"`
	WarehouseStockOnHand          string `json:"warehouse_stock_on_hand,omitempty"`
	WarehouseAvailableStock       string `json:"warehouse_available_stock,omitempty"`
	WarehouseActualAvailableStock string `json:"warehouse_actual_available_stock,omitempty"`
}

type ListItemsResponse struct {
	Items []Item `json:"items,omitempty"`
	ApiResponse
}

type ListItemResponse struct {
	Item Item `json:"items,omitempty"`
	ApiResponse
}

const moduleItems = "items"

// GetCurrentUser will return the currently authenticated users
// https://www.zoho.com/books/api/v3/users/#get-current-user
func (c *API) GetItems() (data *ListItemsResponse, err error) {
	url := fmt.Sprintf("https://books.zoho.%s/api/v3/%s", c.ZohoTLD, moduleItems)
	endpoint := zoho.Endpoint{
		Name:         moduleItems,
		URL:          url,
		Method:       zoho.HTTPGet,
		ResponseData: &ListItemsResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve item[%s]: %s", url, err)
	}

	if v, ok := endpoint.ResponseData.(*ListItemsResponse); ok {
		return v, nil
	}

	return nil, fmt.Errorf("data retrieved was not 'ListItemsResponse'")
}

func (c *API) GetItem(itemID string) (data *ListItemResponse, err error) {
	url := fmt.Sprintf("https://books.zoho.%s/api/v3/%s/%s", c.ZohoTLD, moduleItems, itemID)
	endpoint := zoho.Endpoint{
		Name:         moduleItems,
		URL:          url,
		Method:       zoho.HTTPGet,
		ResponseData: &ListItemResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve item[%s]: %s", url, err)
	}

	if v, ok := endpoint.ResponseData.(*ListItemResponse); ok {
		return v, nil
	}

	return nil, fmt.Errorf("data retrieved was not 'ListItemResponse'")
}

func (c *API) UpdateItem(itemID string, item *Item) (data *ListItemResponse, err error) {
	url := fmt.Sprintf("https://books.zoho.%s/api/v3/%s/%s", c.ZohoTLD, moduleItems, itemID)
	endpoint := zoho.Endpoint{
		Name:         moduleItems,
		URL:          url,
		Method:       zoho.HTTPPut,
		RequestBody:  &item,
		ResponseData: &ListItemResponse{},
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve item[%s]: %s", url, err)
	}

	if v, ok := endpoint.ResponseData.(*ListItemResponse); ok {
		return v, nil
	}

	return nil, fmt.Errorf("data retrieved was not 'ListItemResponse'")
}
