package controllers

import (
	"net/http"

	"github.com/abe16s/Online-Marketplace-API/models"
	"github.com/abe16s/Online-Marketplace-API/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AuthController struct {
    AuthUseCase *usecases.AuthUseCase
}

func (ctrl *AuthController) Register(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
		  validationErrors = errors
		}
	
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
	
		  field := e.Field()
		  switch field {
			case "Email":
				errorMessages["email"] = "email is required."
			case "Password":
				errorMessages["password"] = "password is required."
		  }
		}
	
		if len(errorMessages) == 0 {
			errorMessages["json"] = "Invalid JSON"
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
    }
    user.ID = uuid.New()
	newUser, err := ctrl.AuthUseCase.Register(&user)
    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.IndentedJSON(http.StatusCreated, newUser)
}

func (ctrl *AuthController) Login(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        var validationErrors validator.ValidationErrors
		if errors, ok := err.(validator.ValidationErrors); ok {
		  validationErrors = errors
		}
	
		errorMessages := make(map[string]string)
		for _, e := range validationErrors {
	
		  field := e.Field()
		  switch field {
			case "Email":
				errorMessages["email"] = "email is required."
			case "Password":
				errorMessages["password"] = "password is required."
		  }
		}
	
		if len(errorMessages) == 0 {
			errorMessages["json"] = "Invalid JSON"
		}
		
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
    }

    err := ctrl.AuthUseCase.Login(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.Status(http.StatusNoContent)
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {

    refreshToken, exists := c.Get("refresh_token")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
        return
    }
    
    tokenData, ok := refreshToken.(map[string]interface{})
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }
    newToken, newRefreshToken, err := ctrl.AuthUseCase.RefreshToken(tokenData["email"].(string), tokenData["isSeller"].(bool))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.SetCookie("access_token", newToken, 300, "/", "localhost", false, true)
    c.SetCookie("refresh_token", newRefreshToken, 7776000, "/refreshtoken", "localhost", false, true)
    c.Status(http.StatusNoContent)
}


func (ctrl *AuthController) ActivateAccount(c *gin.Context) {
    token := c.Query("token")
    if token == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activation token"})
        return
    }

    accessToken, refreshToken, err := ctrl.AuthUseCase.ActivateAccount(token)
    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activation token"})
        return
    }

    c.SetCookie("access_token", accessToken, 300, "/", "localhost", false, true)
    c.SetCookie("refresh_token", refreshToken, 7776000, "/refreshtoken", "localhost", false, true)
}

func (ctrl *AuthController) Logout(c *gin.Context) {
    // Clear the cookies by setting their expiration time in the past
    c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
    c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
