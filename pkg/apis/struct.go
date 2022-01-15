package apis

// type card struct {
// 	Number   string `json:"card_number"`
// 	ExpMonth string `json:"expiry_month"`
// 	ExpYear  string `json:"expiry_year"`
// 	CVC      string `json:"cvc"`
// }
type createChargeRequest struct {
	Amount   int64  `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Token    string `json:"stripe_token"`
}
