package discountparam

type ApplyDiscountRequest struct {
	User         string `json:"user"`
	DiscountCode string `json:"discount_code"`
}
type ApplyDiscountResponse struct {
}
