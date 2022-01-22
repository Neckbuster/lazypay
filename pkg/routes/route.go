package routes

import (
	"neckbuster/lazypay/v1/pkg/apis"
	"neckbuster/lazypay/v1/pkg/web"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Web apis
	r.GET("/", web.GetIndexPage)
	// list all charges
	r.GET("/api/v1/get_charge", apis.ListPaymentIntents)
	// PaymentIntent apis
	// create charge
	r.POST("/api/v1/create_charge", apis.CreateCharge)
	// capture charge with Id
	r.POST("/api/v1/capture_charge/:chargeId", apis.Ping)
	// create refund for chargeid
	r.POST("/api/v1/create_refund/:chargeId", apis.CreateRefund)
}
