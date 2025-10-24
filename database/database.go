package database

import (
	"context"
	"fmt"
	"mumuni_backend/config"
	"mumuni_backend/models"

	"github.com/supabase-community/supabase-go"
	"github.com/supabase/postgrest-go"
)

type Database struct {
	client *supabase.Client
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	// Validate configuration
	if cfg.SupabaseURL == "" {
		return nil, fmt.Errorf("SUPABASE_URL is required")
	}
	if cfg.SupabaseAnonKey == "" {
		return nil, fmt.Errorf("SUPABASE_ANON_KEY is required")
	}

	// Log configuration (without sensitive data)
	fmt.Printf("Initializing Supabase client with URL: %s\n", cfg.SupabaseURL)
	fmt.Printf("Anon key length: %d characters\n", len(cfg.SupabaseAnonKey))

	client, err := supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseAnonKey, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create supabase client: %w", err)
	}

	return &Database{client: client}, nil
}

// Appointment methods
func (db *Database) CreateAppointment(ctx context.Context, req *models.AppointmentRequest) (*models.Appointment, error) {
	appointment := map[string]interface{}{
		"name":             req.Name,
		"email":            req.Email,
		"phone":            req.Phone,
		"appointment_date": req.Date,
		"appointment_time": req.Time,
		"service":          req.Service,
		"message":          req.Message,
		"status":           "pending",
	}

	var result []models.Appointment
	_, err := db.client.From("appointments").Insert(appointment, false, "", "", "").ExecuteTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to create appointment: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no appointment created")
	}

	return &result[0], nil
}

func (db *Database) GetAppointments(ctx context.Context) ([]models.Appointment, error) {
	var appointments []models.Appointment
	_, err := db.client.From("appointments").Select("*", "", false).Order("created_at", &postgrest.OrderOpts{Ascending: false}).ExecuteTo(&appointments)
	if err != nil {
		return nil, fmt.Errorf("failed to get appointments: %w", err)
	}

	return appointments, nil
}

// Class methods
func (db *Database) CreateClass(ctx context.Context, req *models.ClassRequest) (*models.Class, error) {
	class := map[string]interface{}{
		"name":               req.Name,
		"email":              req.Email,
		"phone":              req.Phone,
		"class_type":         req.ClassType,
		"experience_level":   req.Experience,
		"goals":              req.Goals,
		"preferred_schedule": req.Schedule,
		"status":             "pending",
	}

	var result []models.Class
	_, err := db.client.From("classes").Insert(class, false, "", "", "").ExecuteTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to create class enrollment: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no class enrollment created")
	}

	return &result[0], nil
}

func (db *Database) GetClasses(ctx context.Context) ([]models.Class, error) {
	var classes []models.Class
	_, err := db.client.From("classes").Select("*", "", false).Order("created_at", &postgrest.OrderOpts{Ascending: false}).ExecuteTo(&classes)
	if err != nil {
		return nil, fmt.Errorf("failed to get classes: %w", err)
	}

	return classes, nil
}

// Admin methods
func (db *Database) CreateAdmin(ctx context.Context, email, passwordHash, name string) (*models.AdminUser, error) {
	admin := map[string]interface{}{
		"email":         email,
		"password_hash": passwordHash,
		"name":          name,
	}

	var result []models.AdminUser
	_, err := db.client.From("admin_users").Insert(admin, false, "", "", "").ExecuteTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to create admin: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no admin created")
	}

	return &result[0], nil
}

func (db *Database) GetAdminByEmail(ctx context.Context, email string) (*models.AdminUser, error) {
	var admins []models.AdminUser
	_, err := db.client.From("admin_users").Select("*", "", false).Eq("email", email).ExecuteTo(&admins)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin by email: %w", err)
	}

	if len(admins) == 0 {
		return nil, fmt.Errorf("admin not found")
	}

	return &admins[0], nil
}

func (db *Database) GetAdminPasswordHash(ctx context.Context, email string) (string, error) {
	var result []struct {
		PasswordHash string `json:"password_hash"`
	}

	_, err := db.client.From("admin_users").Select("password_hash", "", false).Eq("email", email).ExecuteTo(&result)
	if err != nil {
		return "", fmt.Errorf("failed to get admin password hash: %w", err)
	}

	if len(result) == 0 {
		return "", fmt.Errorf("admin not found")
	}

	return result[0].PasswordHash, nil
}
