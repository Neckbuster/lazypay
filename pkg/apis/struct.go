package apis

type createChargeRequest struct {
	Amount   int64  `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type CreateRefundRequest struct {
	PaymentIntentID string `json:"pm_intent_id"`
}

type pmIntent struct {
	Id      string `json:"id"`
	Status  string `json:"status"`
	Amount  int64  `json:"amount"`
	Created string `json:"created"`
}
