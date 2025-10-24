package handlers

import (
	"log"
	"mumuni_backend/database"
	"mumuni_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	db *database.Database
}

func NewHandlers(db *database.Database) *Handlers {
	return &Handlers{db: db}
}

// BookAppointment handles POST /api/appointments
func (h *Handlers) BookAppointment(c *gin.Context) {
	var req models.AppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Validate service type
	validServices := map[string]bool{
		"Bridal Makeup":     true,
		"Event Makeup":      true,
		"Photoshoot Makeup": true,
		"Everyday Glam":     true,
	}

	if !validServices[req.Service] {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid service type. Must be one of: Bridal Makeup, Event Makeup, Photoshoot Makeup, Everyday Glam",
		})
		return
	}

	appointment, err := h.db.CreateAppointment(c.Request.Context(), &req)
	if err != nil {
		log.Printf("Error creating appointment: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to book appointment",
		})
		return
	}

	c.JSON(http.StatusCreated, models.AppointmentResponse{
		Success:     true,
		Appointment: *appointment,
		Message:     "Appointment booked successfully",
	})
}

// EnrollInClass handles POST /api/classes
func (h *Handlers) EnrollInClass(c *gin.Context) {
	var req models.ClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid request data: " + err.Error(),
		})
		return
	}

	// Validate class type
	validClassTypes := map[string]bool{
		"Beginner Basics":     true,
		"Advanced Techniques": true,
		"Bridal Specialist":   true,
		"Business Training":   true,
	}

	if !validClassTypes[req.ClassType] {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid class type. Must be one of: Beginner Basics, Advanced Techniques, Bridal Specialist, Business Training",
		})
		return
	}

	// Validate experience level
	validExperienceLevels := map[string]bool{
		"Complete Beginner": true,
		"Some Experience":   true,
		"Intermediate":      true,
		"Advanced":          true,
	}

	if !validExperienceLevels[req.Experience] {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid experience level. Must be one of: Complete Beginner, Some Experience, Intermediate, Advanced",
		})
		return
	}

	// Validate schedule
	validSchedules := map[string]bool{
		"Weekdays":        true,
		"Weekends":        true,
		"Evening Classes": true,
		"Flexible":        true,
	}

	if !validSchedules[req.Schedule] {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Invalid schedule preference. Must be one of: Weekdays, Weekends, Evening Classes, Flexible",
		})
		return
	}

	class, err := h.db.CreateClass(c.Request.Context(), &req)
	if err != nil {
		log.Printf("Error creating class enrollment: %v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: "Failed to enroll in class",
		})
		return
	}

	c.JSON(http.StatusCreated, models.ClassResponse{
		Success:    true,
		Enrollment: *class,
		Message:    "Class enrollment successful",
	})
}
