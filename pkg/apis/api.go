package apis

import (
	"fmt"
	"neckbuster/lazypay/v1/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func InitializeStripe(stripeKey string) {
	stripe.Key = stripeKey
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
	// create charge using stripe
	params := &stripe.ChargeParams{
		Amount:   &req.Amount,
		Currency: &req.Currency,
	}

	params.SetSource(req.Token)

	nc, err := charge.New(params)
	if err != nil {
		logger.LogMessage(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	fmt.Println(nc.Status)
	// return success message
	c.JSON(200, gin.H{
		"message": "Charge Created Successfully",
	})
}
