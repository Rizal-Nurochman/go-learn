package controllers

import (
	"net/http"
	"os"
	"time"

	models "github.com/NurochmanR/GO-JWT/MODELS"
	"github.com/NurochmanR/GO-JWT/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//Get the email/pass off req body
	var body struct{
		Email string
		Password string
	}

	err:=c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	//hash the password
	hash, err:= bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}

	//create the user
	user:=models.User{
		Email: body.Email,
		Password: string(hash),
	}
	createUser:=initializers.DB.Create(&user).Error
	if createUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":createUser.Error(),
		})
		return
	}

	//respond
		c.JSON(http.StatusOK, gin.H{
		"message":"Create account success",
	})
}

func Login(c *gin.Context) {
	//Get Email and Pass off Req
	var body struct{
		Email string
		Password string
	}

	err:=c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}
	//Look up requested user
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)
	if user.ID==0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid email or password",
		})
		return
	}
	//compare sent in pass with saved user pass hash
	err=bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"invalid email or password", 
		})
		return
	}
	//generate a jwt token
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":user.ID,
		"exp":time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//send it back 
	tokenString, err:=token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "","", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validator(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message":"I'm logged in",
	})
}