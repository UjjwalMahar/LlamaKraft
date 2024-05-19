package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/UjjwalMahar/llamakraft/initializers"
	"github.com/UjjwalMahar/llamakraft/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//for sign-up

func SignUp(c *gin.Context){

	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})
		return
	}

	//Hash Password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
	}
	
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
	})

}

func Login(c *gin.Context){

	//Need to user username and password for login

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body", 
		})
	}

	var user models.User

	initializers.DB.First(&user , "Username = ?", body.Username)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 *30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "","", false,true)

	c.JSON(http.StatusOK, gin.H{
	})
}

func Validate(c *gin.Context){

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message":user,
	})
}

func Health(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "UP",
	})

}