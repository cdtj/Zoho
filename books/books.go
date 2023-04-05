package books

import (
	"math/rand"
	"time"

	zoho "github.com/schmorrison/Zoho"
)

// API is used for interacting with the Zoho Books API
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *books.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			source := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(source)
			id = append(id, keyspace[rnd.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		Zoho: z,
		id:   id,
	}
}

type ApiResponse struct {
	Code        int    `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	PageContext struct {
		Page          int    `json:"page,omitempty"`
		PerPage       int    `json:"per_page,omitempty"`
		HasMorePage   bool   `json:"has_more_page,omitempty"`
		AppliedFilter string `json:"applied_filter,omitempty"`
		ReportName    string `json:"report_name,omitempty"`
		SortColumn    string `json:"sort_column,omitempty"`
		SortOrder     string `json:"sort_order,omitempty"`
	} `json:"page_context,omitempty"`
}
