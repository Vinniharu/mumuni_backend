package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8080"

func main() {
	fmt.Println("🧪 Testing Mumuni Backend Status Update API...")
	fmt.Println("=============================================")

	// Test admin login to get token
	fmt.Println("\n1. Testing Admin Login...")
	token := testAdminLogin()
	if token == "" {
		fmt.Println("❌ Cannot proceed without admin token")
		return
	}

	// Test appointment booking
	fmt.Println("\n2. Testing Appointment Booking...")
	appointmentID := testAppointmentBooking()
	if appointmentID == 0 {
		fmt.Println("❌ Cannot test status update without appointment")
		return
	}

	// Test class enrollment
	fmt.Println("\n3. Testing Class Enrollment...")
	classID := testClassEnrollment()
	if classID == 0 {
		fmt.Println("❌ Cannot test status update without class")
		return
	}

	// Test appointment status updates
	fmt.Println("\n4. Testing Appointment Status Updates...")
	testAppointmentStatusUpdates(token, appointmentID)

	// Test class status updates
	fmt.Println("\n5. Testing Class Status Updates...")
	testClassStatusUpdates(token, classID)

	fmt.Println("\n=============================================")
	fmt.Println("✅ Status update API testing completed!")
}

func testAdminLogin() string {
	loginData := map[string]interface{}{
		"email":    "admin@mumuni.com",
		"password": "admin123456",
	}

	jsonData, _ := json.Marshal(loginData)
	resp, err := http.Post(baseURL+"/api/admin/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Admin login failed: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if token, ok := result["token"].(string); ok {
			fmt.Println("✅ Admin login successful")
			return token
		}
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("❌ Admin login failed: %s\n", string(body))
	return ""
}

func testAppointmentBooking() int {
	appointmentData := map[string]interface{}{
		"name":    "Sarah Johnson",
		"email":   "sarah@email.com",
		"phone":   "+234-123-456-7890",
		"date":    "2024-02-15",
		"time":    "2:00 PM",
		"service": "Bridal Makeup",
		"message": "Wedding on March 1st, need trial session",
	}

	jsonData, _ := json.Marshal(appointmentData)
	resp, err := http.Post(baseURL+"/api/appointments", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Appointment booking failed: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if appointment, ok := result["appointment"].(map[string]interface{}); ok {
			if id, ok := appointment["id"].(float64); ok {
				fmt.Println("✅ Appointment booking successful")
				return int(id)
			}
		}
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("❌ Appointment booking failed: %s\n", string(body))
	return 0
}

func testClassEnrollment() int {
	classData := map[string]interface{}{
		"name":       "Maria Garcia",
		"email":      "maria@email.com",
		"phone":      "+234-987-654-3210",
		"classType":  "Beginner Basics",
		"experience": "Complete Beginner",
		"goals":      "Want to learn basic makeup for personal use",
		"schedule":   "Weekends",
	}

	jsonData, _ := json.Marshal(classData)
	resp, err := http.Post(baseURL+"/api/classes", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Class enrollment failed: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if enrollment, ok := result["enrollment"].(map[string]interface{}); ok {
			if id, ok := enrollment["id"].(float64); ok {
				fmt.Println("✅ Class enrollment successful")
				return int(id)
			}
		}
	}

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("❌ Class enrollment failed: %s\n", string(body))
	return 0
}

func testAppointmentStatusUpdates(token string, appointmentID int) {
	client := &http.Client{Timeout: 10 * time.Second}

	statuses := []string{"confirmed", "cancelled", "completed"}

	for _, status := range statuses {
		statusData := map[string]interface{}{
			"status": status,
		}

		jsonData, _ := json.Marshal(statusData)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/api/admin/appointments/%d/status", baseURL, appointmentID), bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Update appointment to %s failed: %v\n", status, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("✅ Appointment status updated to %s\n", status)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("❌ Update appointment to %s failed: %s\n", status, string(body))
		}
	}
}

func testClassStatusUpdates(token string, classID int) {
	client := &http.Client{Timeout: 10 * time.Second}

	statuses := []string{"confirmed", "cancelled", "completed"}

	for _, status := range statuses {
		statusData := map[string]interface{}{
			"status": status,
		}

		jsonData, _ := json.Marshal(statusData)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/api/admin/classes/%d/status", baseURL, classID), bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Update class to %s failed: %v\n", status, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			fmt.Printf("✅ Class status updated to %s\n", status)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("❌ Update class to %s failed: %s\n", status, string(body))
		}
	}
}
