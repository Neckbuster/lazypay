package main

import (
	"neckbuster/lazypay/v1/pkg/apis"
	"neckbuster/lazypay/v1/pkg/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	stripeKeyPublish = os.Getenv("STRIPE_PUBLISHABLE_KEY")
	stripeKeySecret  = os.Getenv("STRIPE_SECRET_KEY")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// Add support for html files
	r.LoadHTMLGlob("pkg/templates/*.tmpl")
	// Allow all origin
	conf := cors.DefaultConfig()
	conf.AllowOrigins = []string{"*"}
	r.Use(cors.New(conf))
	// initialize all routes
	routes.InitRoutes(r)
	// Initialize StripeKey
	apis.InitializeStripe(stripeKeySecret, stripeKeyPublish)
	// start gin server
	r.Run()
}
