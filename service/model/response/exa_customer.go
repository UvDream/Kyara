package response

import "blog-api/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
