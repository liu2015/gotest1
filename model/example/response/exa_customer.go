package response

import "ginserver/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"costomer"`
}
