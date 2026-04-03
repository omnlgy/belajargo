package controller

import (
	"net/http"
	"time"

	"example.com/rest/helpers"
	"example.com/rest/models"
	"github.com/gin-gonic/gin"
)

func SignupUser(ctx *gin.Context) {
	data, exist := ctx.Get("data")

	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get data from validation middleware",
		})
		return
	}

	signupReq := data.(models.SignupRequest)
	user := signupReq.ToUser()
	user.Password, _ = helpers.HashPassword(user.Password)

	err := user.CreateUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func LoginUser(ctx *gin.Context) {
	data, exist := ctx.Get("data")

	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get data from validation middleware",
		})
		return
	}

	loginReq := data.(models.LoginRequest)
	user := loginReq.ToUser()

	userFromDB, err := user.GetUser()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "email or password is incorrect",
		})
		return
	}

	isPasswordValid := helpers.CheckPasswordHash(user.Password, userFromDB.Password)

	if !isPasswordValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "email or password is incorrect",
		})
		return
	}

	token, err := helpers.GenerateToken(*userFromDB)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
		return
	}

	ctx.SetCookie("tes", "tescookie", int((time.Minute * 5).Seconds()), "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
