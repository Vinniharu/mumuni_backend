package handlers

import (
	"log"
	"mumuni_backend/auth"
	"mumuni_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminSignup handles POST /api/admin/signup
func (h *Handlers) AdminSignup(c *gin.Context) {
	var req models.AdminSignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Check if admin already exists
	_, err := h.db.GetAdminByEmail(c.Request.Context(), req.Email)
	if err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{
			Error: "Admin with this email already exists",
		})
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to process password",
		})
		return
	}

	// Create admin
	_, err = h.db.CreateAdmin(c.Request.Context(), req.Email, hashedPassword, req.Name)
	if err != nil {
		log.Printf("Error creating admin: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to create admin account",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SuccessResponse{
		Success: true,
		Message: "Admin account created successfully",
	})
}

// AdminLogin handles POST /api/admin/login
func (h *Handlers) AdminLogin(c *gin.Context) {
	var req models.AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Get admin by email
	admin, err := h.db.GetAdminByEmail(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid email or password",
		})
		return
	}

	// Get password hash
	passwordHash, err := h.db.GetAdminPasswordHash(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid email or password",
		})
		return
	}

	// Check password
	if !auth.CheckPasswordHash(req.Password, passwordHash) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(admin.ID, admin.Email)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to generate authentication token",
		})
		return
	}

	c.JSON(http.StatusOK, models.AdminLoginResponse{
		Success: true,
		Token:   token,
		Admin:   *admin,
		Message: "Login successful",
	})
}

// GetAppointments handles GET /api/admin/appointments
func (h *Handlers) GetAppointments(c *gin.Context) {
	appointments, err := h.db.GetAppointments(c.Request.Context())
	if err != nil {
		log.Printf("Error getting appointments: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to fetch appointments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"appointments": appointments,
		"count":        len(appointments),
	})
}

// GetClasses handles GET /api/admin/classes
func (h *Handlers) GetClasses(c *gin.Context) {
	classes, err := h.db.GetClasses(c.Request.Context())
	if err != nil {
		log.Printf("Error getting classes: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to fetch class enrollments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"classes": classes,
		"count":   len(classes),
	})
}
