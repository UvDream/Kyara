package response

import "server/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
