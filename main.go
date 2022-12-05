package main

import (
	"Restapi/config"
	"Restapi/routes"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		ginLambda = ginadapter.New(ginEngine())
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func ginEngine() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	// app.SetTrustedProxies([]string{"127.0.0.1"})
	app.Use(corsMiddleware())
	config.ConnectMailer(
		os.Getenv("MAILER_USERNAME"),
		os.Getenv("MAILER_PASSWORD"),
	)
	app.Use(gin.Logger())
	app.GET("/", test)
	app.NoRoute(routes.Stoproute)
	routes.MapRoutes(app)
	return app
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With,userToken")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Start Project"})
	c.Abort()
	return
}

func main() {
	if gin.Mode() == "release" {
		lambda.Start(lambdaHandler)
	} else {
		config.Envload()
		port := os.Getenv("PORT")
		if port == "" {
			port = "80"
		}
		app := ginEngine()
		log.Fatal(app.Run(":" + port))
	}
}
