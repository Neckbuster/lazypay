package web

import (
	"neckbuster/lazypay/v1/pkg/apis"

	"github.com/gin-gonic/gin"
)

func GetIndexPage(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"message":            "Page Loaded",
		"amount":             1000,
		"stripe_publish_key": apis.StripePublishableKey,
	})
}
