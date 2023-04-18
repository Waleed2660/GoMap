package Gateway

import (
	"github.com/Waleed2660/GoMap/src/main/Helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type InboundGateway struct {
	// You can define any variables or dependencies needed by the controller here
}

// NewInboundGateway creates a new instance of InboundGatewayController
func NewInboundGateway() *InboundGateway {
	return &InboundGateway{}
}

// healthCheck endpoint
func (c *InboundGateway) healthCheck(ctx *gin.Context) {
	// Logic for handling GET request for /healthcheck
	log.Printf("Received request: Method=%s, URL=%s\n", ctx.Request.Method, ctx.Request.URL.Path)
	ctx.String(http.StatusOK, "OK, GoMap is up & running !")
}

// handleGetRequest2 dummy endpoint
func (c *InboundGateway) getApiKey(ctx *gin.Context) {
	// Gets the API key from handler
	key, err := Helper.GetApiKey()
	if err == nil {
		ctx.Header("Content-Type", "text/plain")
		ctx.String(http.StatusOK, "Google API Key: %s", key)
	}
}

// Run starts the HTTP server and registers the controller to handle incoming requests
func (c *InboundGateway) Run() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes and associate them with handler functions
	router.GET("/healthcheck", c.healthCheck)
	router.GET("/key", c.getApiKey)

	// Start the HTTP server and listen for incoming requests on port 8080
	log.Println("Starting HTTP server on :8080...")
	log.Fatal(router.Run(":8080"))
}
