package apis

import (
	"neckbuster/lazypay/v1/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/refund"
)

func InitializeStripe(stripeKeySecret, stripeSecretPublish string) {
	stripe.Key = stripeKeySecret
	StripePublishableKey = stripeSecretPublish
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateCharge(c *gin.Context) {
	// bind request with createCharge object
	var req createChargeRequest
	err := c.BindJSON(&req)
	if err != nil {
		logger.LogMessage(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// create PaymentIntent using stripe
	params := &stripe.PaymentIntentParams{
		Amount:   &req.Amount,
		Currency: stripe.String(string(stripe.CurrencyINR)),
	}
	pm, err := paymentintent.New(params)
	if err != nil {
		logger.LogMessage(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// return success message
	c.JSON(200, gin.H{
		"message":       "Charge Created Successfully",
		"client_secret": pm.ClientSecret,
	})
}

func ListPaymentIntents(c *gin.Context) {
	// list payment intents
	var intents []pmIntent
	params := &stripe.PaymentIntentListParams{}
	// params.Filters.AddFilter("limit", "", "3")
	i := paymentintent.List(params)
	for i.Next() {
		pi := i.PaymentIntent()
		intents = append(intents, pmIntent{
			Id:      pi.ID,
			Status:  string(pi.Status),
			Amount:  pi.Amount,
			Created: time.Unix(pi.Created, 0).Format("2006-01-02 15:04:05"),
		})
	}

	c.HTML(200, "list.tmpl", gin.H{
		"Intents": intents,
	})

}

func CreateRefund(c *gin.Context) {
	chargeId, isPresent := c.Params.Get("chargeId")
	if !isPresent {
		logger.LogMessage("Missing: ChargeID params is missing")
		c.JSON(400, gin.H{
			"message": "Internal server error",
		})
		return
	}
	// create refund
	params := &stripe.RefundParams{
		PaymentIntent: stripe.String(chargeId),
	}
	result, err := refund.New(params)
	if err != nil {
		logger.LogMessage(err.Error())
		c.JSON(400, gin.H{
			"message":   "Failed to Refund",
			"refund_id": result.ID,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "Refund Successfull",
		"refund_id": result.ID,
	})

}
