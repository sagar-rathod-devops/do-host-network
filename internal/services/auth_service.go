package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/config"
	"github.com/sagar-rathod-devops/do-host-network/internal/models"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
	"github.com/sagar-rathod-devops/do-host-network/utils"
	"golang.org/x/crypto/bcrypt"
)

var otpLength = 6 // Global variable for OTP length

type AuthService struct {
	DB                  *sql.DB
	UserRepository      repositories.UserRepository
	OTPRepository       repositories.OTPRepository
	TokenExpiration     time.Duration
	OTPLifespan         time.Duration
	BlacklistRepository repositories.TokenBlacklistRepository
}

// RegisterUserWithOTP handles the registration of a new user and sends an OTP.
func (s *AuthService) RegisterUserWithOTP(ctx context.Context, email, username, password string) error {
	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Create the user model
	user := models.User{
		Email:        email,
		Username:     username,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save the user in the database
	err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Generate OTP
	otp := utils.GenerateOTP(otpLength)

	// Create OTP record
	otpRecord := models.OTP{
		Email:      email,
		OTP:        otp,
		IsVerified: false,
		CreatedAt:  time.Now(),
	}

	// Save OTP to the database
	err = s.OTPRepository.SaveOTP(ctx, otpRecord)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}

	// Send OTP email to the user
	subject := "Verify Your Email Address with Do Host Network"
	body := fmt.Sprintf(`Hello,

Thank you for choosing Do Host Network. Please use the One-Time Password (OTP) below to verify your email address.

Your OTP is: %s

This OTP is valid for the next 10 minutes. Please keep this code confidential and do not share it with anyone.

If you did not request this email, please contact our support team or ignore this message.

Best regards,
The Do Host Network Team`, otp)

	err = utils.SendEmail(email, subject, body)
	if err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	fmt.Println("Registration and OTP email sent successfully to", email)
	return nil
}

// LoginUser handles user login and token generation (pure business logic)
func (s *AuthService) LoginUser(ctx context.Context, email, password string) (string, error) {
	var hashedPassword, userID string

	// Query the user by email
	err := s.DB.QueryRowContext(ctx, `SELECT id, password_hash FROM users WHERE email = $1`, email).Scan(&userID, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Load the JWT secret key from config
	cfg, err := config.LoadConfig(".")
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(24*time.Hour, userID, cfg.TokenSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// LogoutUser handles user logout by invalidating the JWT token
func (s *AuthService) LogoutUser(ctx *gin.Context) {
	// Invalidate the token (e.g., clear cookie or blacklist the token)
	// Clear the token cookie
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)

	// Optionally, perform any other cleanup related to the token or session.

	// No return needed if the action is successful
}

// VerifyOTP handles OTP verification
func (s *AuthService) VerifyOTP(ctx context.Context, email, otp string) error {
	storedOTP, err := s.OTPRepository.GetOTPByEmail(ctx, email)
	if err != nil {
		return errors.New("OTP not found or expired")
	}

	if storedOTP.OTP != otp {
		return errors.New("invalid OTP")
	}

	return s.OTPRepository.MarkUserVerified(ctx, email)
}

// ForgotPassword generates an OTP for password reset and sends an email
func (s *AuthService) ForgotPassword(ctx context.Context, email string) error {
	// Generate OTP
	otp := utils.GenerateOTP(6) // Generate a 6-digit OTP

	// Create OTP record
	otpRecord := models.OTP{
		Email:      email,
		OTP:        otp,
		IsVerified: false,
		CreatedAt:  time.Now(),
	}

	// Save OTP to the database
	err := s.OTPRepository.SaveOTP(ctx, otpRecord)
	if err != nil {
		return fmt.Errorf("failed to save OTP: %w", err)
	}

	// Send email to the user
	subject := "Reset Your Password - Do Host Network"
	body := fmt.Sprintf(`Hello,

We received a request to reset the password for your Do Host Network account. Please use the One-Time Password (OTP) below to reset your password.

Your OTP is: %s

This OTP is valid for the next 10 minutes. If you did not request this password reset, please contact our support team or ignore this email.

Best regards,
The Do Host Network Team`, otp)

	err = utils.SendEmail(email, subject, body)
	if err != nil {
		return fmt.Errorf("failed to send OTP email: %w", err)
	}

	fmt.Println("Password reset email sent successfully to", email)
	return nil
}

// ResetPassword resets the user's password
func (s *AuthService) ResetPassword(ctx context.Context, email, otp, newPassword string) error {
	if err := s.VerifyOTP(ctx, email, otp); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.UserRepository.UpdatePassword(ctx, email, hashedPassword)
}
