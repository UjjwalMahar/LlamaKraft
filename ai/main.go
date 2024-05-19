package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var client *genai.Client
var model *genai.GenerativeModel

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env key:", err)
	}

	apiKey := os.Getenv("API_KEY")

	ctx := context.Background()

	client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Println("Error creating genai client:", err)
		return 
	}
	
	defer client.Close()

	
	model = client.GenerativeModel("gemini-pro")


	r := gin.Default()


	r.POST("/chat", chatHandler)

	r.GET("/health", Health)


	r.Run(":8080")
}

func Health(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message": "UP",
	})

}
func chatHandler(c *gin.Context){

	var req struct {
		Message string `json:"message"`
	}


	if c.Bind(&req) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the message", 
		})
	}

	ctx := context.Background()

	
	resp, err := model.GenerateContent(ctx, genai.Text(req.Message))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error generating text: %v", err)})
		return 
	}

	
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
