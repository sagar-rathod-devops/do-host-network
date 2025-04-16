package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

// Register handles user registration and sends OTP.
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode the request body.
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Register the user and send the OTP.
	if err := c.AuthService.RegisterUserWithOTP(r.Context(), payload.Email, payload.Username, payload.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header and write the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
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
		"access_token": token,
		"message":      "Login successful",
	})
}

// LogoutUser handles user logout by invalidating the JWT token
func (c *AuthController) LogoutUser(ctx *gin.Context) {
	// Call the service to log out the user
	c.AuthService.LogoutUser(ctx) // No need to handle error if there's no return value

	// Log the message to the console
	log.Printf(`{
		"message": "Logged out successfully",
		"status": "success"
	}`)

	// Return the response in the expected format
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
		"status":  "success",
	})
}

// VerifyOTP handles OTP verification.
func (c *AuthController) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	// Decode the request body.
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verify the OTP.
	if err := c.AuthService.VerifyOTP(r.Context(), payload.Email, payload.OTP); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the response header and write the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP verified successfully",
	})
}

// ForgotPassword handles OTP generation for password reset.
func (c *AuthController) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
	}

	// Decode the request body.
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Generate OTP for password reset.
	if err := c.AuthService.ForgotPassword(r.Context(), payload.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header and write the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP sent successfully",
	})
}

// ResetPassword handles password reset.
func (c *AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email       string `json:"email"`
		OTP         string `json:"otp"`
		NewPassword string `json:"new_password"`
	}

	// Decode the request body.
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Reset the password.
	if err := c.AuthService.ResetPassword(r.Context(), payload.Email, payload.OTP, payload.NewPassword); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header and write the response body.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset successfully",
	})
}
