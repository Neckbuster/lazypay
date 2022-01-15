package main

import (
	"neckbuster/lazypay/v1/pkg/apis"
	"neckbuster/lazypay/v1/pkg/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	stripeKey = os.Getenv("STRIPE_SECRET_KEY")
)

func main() {
	r := gin.Default()
	// Allow all origin
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"*"}
	r.Use(cors.New(conf))
	// initialize all routes
	routes.InitRoutes(r)
	// Initialize StripeKey
	apis.InitializeStripe(stripeKey)
	// start gin server
	r.Run()
}
