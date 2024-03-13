package handlers

import (
	"bookmyshow/auth"
	"bookmyshow/models"
	"bookmyshow/postgresql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var body models.AuthLoginReqModel

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
	}

	//check in db if the user exists with the username
	id, password, err := postgresql.GetIdPasswordQuery(body.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	// verify the password
	isVerified := comparePassword(password, body.Password)
	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password did not match",
		})
		return
	}
	// generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//save tokens in db
	err = postgresql.UpdateUserTokensQuery(refreshToken, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "updating tokens",
		})
		return
	}
	//send the tokens to frontend
	var res models.AuthResModel
	res.Username = body.Username
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func Signup(c *gin.Context) {
	var body models.AuthSignupReqModel

	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	//hash the password
	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//create user
	var user models.User
	user.Username = body.Username
	user.Password = hashedPassword
	user.Name = body.Name

	//save user in db
	id, err := postgresql.RegisterNewUserQuery(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//generate tokens
	accessToken, refreshToken, err := auth.GenerateTokens(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//update user refresh token
	err = postgresql.UpdateUserTokensQuery(refreshToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResModel
	res.Username = body.Username
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func RefreshToken(c *gin.Context) {
	var body models.RefreshreqModel
	err := c.ShouldBind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	refreshToken := body.RefreshToken

	id, err := auth.ExtractIdFromToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok := postgresql.CompareRefreshToken(refreshToken, id)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Refresh token has changed or is invalid",
		})
		return
	}
	accessToken, err := auth.RefreshAccessToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var res models.AuthResModel
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	c.JSON(http.StatusOK, res)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return string(hash), fmt.Errorf("error hashing the password : %v", err)
	}

	return string(hash), nil
}

func comparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
