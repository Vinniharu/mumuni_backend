package main

import (
	"log"
	"mumuni_backend/auth"
	"mumuni_backend/config"
	"mumuni_backend/database"
	"mumuni_backend/handlers"
	"mumuni_backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Validate required configuration
	if cfg.SupabaseURL == "" || cfg.SupabaseAnonKey == "" {
		log.Fatal("Missing required Supabase configuration. Please set SUPABASE_URL and SUPABASE_ANON_KEY")
	}

	// Set JWT secret
	auth.SetJWTSecret(cfg.JWTSecret)

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize handlers
	h := handlers.NewHandlers(db)

	// Set Gin mode
	if gin.Mode() == gin.DebugMode {
		log.Println("Running in debug mode")
	}

	// Create Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "mumuni_backend",
		})
	})

	// Public API routes
	api := r.Group("/api")
	{
		// Appointment booking
		api.POST("/appointments", h.BookAppointment)

		// Class enrollment
		api.POST("/classes", h.EnrollInClass)
	}

	// Admin routes
	admin := r.Group("/api/admin")
	{
		// Admin authentication (no auth required)
		admin.POST("/signup", h.AdminSignup)
		admin.POST("/login", h.AdminLogin)

		// Protected admin routes
		adminProtected := admin.Group("/")
		adminProtected.Use(middleware.AuthMiddleware())
		{
			adminProtected.GET("/appointments", h.GetAppointments)
			adminProtected.GET("/classes", h.GetClasses)
		}
	}

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Health check: http://localhost:%s/health", port)
	log.Printf("API endpoints:")
	log.Printf("  POST /api/appointments - Book appointment")
	log.Printf("  POST /api/classes - Enroll in class")
	log.Printf("  POST /api/admin/signup - Admin signup")
	log.Printf("  POST /api/admin/login - Admin login")
	log.Printf("  GET /api/admin/appointments - Get appointments (requires auth)")
	log.Printf("  GET /api/admin/classes - Get classes (requires auth)")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
