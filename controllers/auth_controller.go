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

    accessToken, refreshToken, err := ctrl.AuthUseCase.Login(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
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

    c.JSON(http.StatusOK, gin.H{"token": newToken, "refresh_token": newRefreshToken})
}


func (ctrl *AuthController) Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
