package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/UjjwalMahar/llamakraft/auth/initializers"
	"github.com/UjjwalMahar/llamakraft/auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	//Get the cookies
	tokenString , err := c.Cookie("Authorization")

	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
	
	
		return []byte(os.Getenv("SECRET")), nil
	})
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
	//Check expriration
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	//Find the user with the token "sub"
	var user models.User
	initializers.DB.First(&user, claims["sub"])

	if user.ID ==0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//Attach req

	c.Set("user", user)

	//Continue
		
	c.Next()

	fmt.Println(claims["foo"], claims["nbf"])

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	



}