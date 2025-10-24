#!/bin/bash

# Mumuni Backend API Test Script
# Run this after starting the server with: go run main.go

BASE_URL="http://localhost:8080"

echo "🧪 Testing Mumuni Backend API..."
echo "=============================="

# Test Health Check
echo -e "\n1. Testing Health Check..."
curl -s -o /dev/null -w "%{http_code}" "$BASE_URL/health" | grep -q "200" && echo "✅ Health check passed" || echo "❌ Health check failed"

# Test Admin Signup
echo -e "\n2. Testing Admin Signup..."
ADMIN_SIGNUP_RESPONSE=$(curl -s -X POST "$BASE_URL/api/admin/signup" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@mumuni.com","password":"admin123456","name":"Admin User"}')

if echo "$ADMIN_SIGNUP_RESPONSE" | grep -q "success"; then
  echo "✅ Admin signup successful"
else
  echo "❌ Admin signup failed: $ADMIN_SIGNUP_RESPONSE"
fi

# Test Admin Login
echo -e "\n3. Testing Admin Login..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/api/admin/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@mumuni.com","password":"admin123456"}')

TOKEN=$(echo "$LOGIN_RESPONSE" | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

if [ -n "$TOKEN" ]; then
  echo "✅ Admin login successful"
else
  echo "❌ Admin login failed: $LOGIN_RESPONSE"
fi

# Test Appointment Booking
echo -e "\n4. Testing Appointment Booking..."
APPOINTMENT_RESPONSE=$(curl -s -X POST "$BASE_URL/api/appointments" \
  -H "Content-Type: application/json" \
  -d '{"name":"Sarah Johnson","email":"sarah@email.com","phone":"+234-123-456-7890","date":"2024-02-15","time":"2:00 PM","service":"Bridal Makeup","message":"Wedding on March 1st"}')

if echo "$APPOINTMENT_RESPONSE" | grep -q "success"; then
  echo "✅ Appointment booking successful"
else
  echo "❌ Appointment booking failed: $APPOINTMENT_RESPONSE"
fi

# Test Class Enrollment
echo -e "\n5. Testing Class Enrollment..."
CLASS_RESPONSE=$(curl -s -X POST "$BASE_URL/api/classes" \
  -H "Content-Type: application/json" \
  -d '{"name":"Maria Garcia","email":"maria@email.com","phone":"+234-987-654-3210","classType":"Beginner Basics","experience":"Complete Beginner","goals":"Learn basic makeup","schedule":"Weekends"}')

if echo "$CLASS_RESPONSE" | grep -q "success"; then
  echo "✅ Class enrollment successful"
else
  echo "❌ Class enrollment failed: $CLASS_RESPONSE"
fi

# Test Admin Endpoints (if token available)
if [ -n "$TOKEN" ]; then
  echo -e "\n6. Testing Admin Endpoints..."
  
  # Test Get Appointments
  APPOINTMENTS_RESPONSE=$(curl -s -X GET "$BASE_URL/api/admin/appointments" \
    -H "Authorization: Bearer $TOKEN")
  
  if echo "$APPOINTMENTS_RESPONSE" | grep -q "success"; then
    echo "✅ Get appointments successful"
  else
    echo "❌ Get appointments failed: $APPOINTMENTS_RESPONSE"
  fi
  
  # Test Get Classes
  CLASSES_RESPONSE=$(curl -s -X GET "$BASE_URL/api/admin/classes" \
    -H "Authorization: Bearer $TOKEN")
  
  if echo "$CLASSES_RESPONSE" | grep -q "success"; then
    echo "✅ Get classes successful"
  else
    echo "❌ Get classes failed: $CLASSES_RESPONSE"
  fi
fi

echo -e "\n=============================="
echo "✅ API testing completed!"
