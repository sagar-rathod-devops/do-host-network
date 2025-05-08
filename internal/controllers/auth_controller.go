package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

// Register handles user registration and sends OTP.
func (c *AuthController) Register(ctx *gin.Context) {
	var payload struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind the JSON body to the payload struct.
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Register the user and send the OTP.
	if err := c.AuthService.RegisterUserWithOTP(ctx, payload.Email, payload.Username, payload.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send success response.
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully. OTP sent to email.",
	})
}

// Login handles user login and returns a token.
func (c *AuthController) Login(ctx *gin.Context) {
	var payload struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON input
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service to log in the user and get a token
	token, err := c.AuthService.LoginUser(ctx, payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set token as a secure cookie (1 day expiry)
	ctx.SetCookie("token", token, 3600*24, "/", "localhost", false, true)

	// Return the token in the response
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// LogoutUser handles user logout by invalidating the JWT token
func (c *AuthController) LogoutUser(ctx *gin.Context) {
	// Call the service to log out the user
	c.AuthService.LogoutUser(ctx) // No need to handle error if there's no return value

	// Log the message to the console
	log.Printf(`{
		"message": "Logged out successfully",
	}`)

	// Return the response in the expected format
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// VerifyOTP handles OTP verification.
func (c *AuthController) VerifyOTP(ctx *gin.Context) {
	var payload struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	// Bind the request body.
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Verify the OTP.
	if err := c.AuthService.VerifyOTP(ctx, payload.Email, payload.OTP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond success.
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}

// ForgotPassword handles OTP generation for password reset.
func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var payload struct {
		Email string `json:"email"`
	}

	// Bind the request body.
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Generate OTP.
	if err := c.AuthService.ForgotPassword(ctx, payload.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond success.
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// ResetPassword handles password reset.
func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var payload struct {
		Email       string `json:"email"`
		OTP         string `json:"otp"`
		NewPassword string `json:"new_password"`
	}

	// Bind the request body.
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Reset the password.
	if err := c.AuthService.ResetPassword(ctx, payload.Email, payload.OTP, payload.NewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond success.
	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
