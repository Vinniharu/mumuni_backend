package models

import (
	"time"
)

// Appointment represents a makeup appointment booking
type Appointment struct {
	ID              int       `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Email           string    `json:"email" db:"email"`
	Phone           string    `json:"phone" db:"phone"`
	AppointmentDate string    `json:"appointment_date" db:"appointment_date"`
	AppointmentTime string    `json:"appointment_time" db:"appointment_time"`
	Service         string    `json:"service" db:"service"`
	Message         *string   `json:"message" db:"message"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// AppointmentRequest represents the request payload for booking an appointment
type AppointmentRequest struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required,email"`
	Phone   string  `json:"phone" binding:"required"`
	Date    string  `json:"date" binding:"required"`
	Time    string  `json:"time" binding:"required"`
	Service string  `json:"service" binding:"required"`
	Message *string `json:"message"`
}

// AppointmentResponse represents the response for appointment booking
type AppointmentResponse struct {
	Success     bool        `json:"success"`
	Appointment Appointment `json:"appointment"`
	Message     string      `json:"message"`
}

// Class represents a makeup class enrollment
type Class struct {
	ID                int       `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	Email             string    `json:"email" db:"email"`
	Phone             string    `json:"phone" db:"phone"`
	ClassType         string    `json:"class_type" db:"class_type"`
	ExperienceLevel   string    `json:"experience_level" db:"experience_level"`
	Goals             *string   `json:"goals" db:"goals"`
	PreferredSchedule string    `json:"preferred_schedule" db:"preferred_schedule"`
	Status            string    `json:"status" db:"status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// ClassRequest represents the request payload for class enrollment
type ClassRequest struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Phone      string  `json:"phone" binding:"required"`
	ClassType  string  `json:"classType" binding:"required"`
	Experience string  `json:"experience" binding:"required"`
	Goals      *string `json:"goals"`
	Schedule   string  `json:"schedule" binding:"required"`
}

// ClassResponse represents the response for class enrollment
type ClassResponse struct {
	Success    bool   `json:"success"`
	Enrollment Class  `json:"enrollment"`
	Message    string `json:"message"`
}

// AdminUser represents an admin user
type AdminUser struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// AdminSignupRequest represents the request payload for admin signup
type AdminSignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}

// AdminLoginRequest represents the request payload for admin login
type AdminLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AdminLoginResponse represents the response for admin login
type AdminLoginResponse struct {
	Success bool      `json:"success"`
	Token   string    `json:"token"`
	Admin   AdminUser `json:"admin"`
	Message string    `json:"message"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
