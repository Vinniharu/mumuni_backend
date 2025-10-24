# Mumuni Backend API Documentation

## Overview
This is a Go backend service for a makeup business that handles appointment bookings, class enrollments, and admin management with Supabase integration.

## Base URL
```
http://localhost:8080
```

## Authentication
Admin endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

---

## 📅 Appointment Booking

### Book an Appointment
**POST** `/api/appointments`

Creates a new makeup appointment booking.

#### Request Body
```json
{
  "name": "string (required)",
  "email": "string (required, valid email)",
  "phone": "string (required)",
  "date": "string (required, YYYY-MM-DD format)",
  "time": "string (required)",
  "service": "string (required, see valid services below)",
  "message": "string (optional)"
}
```

#### Valid Services
- `Bridal Makeup` (₦50,000 - ₦80,000)
- `Event Makeup` (₦25,000 - ₦40,000)
- `Photoshoot Makeup` (₦30,000 - ₦45,000)
- `Everyday Glam` (₦15,000 - ₦25,000)

#### Example Request
```json
{
  "name": "Sarah Johnson",
  "email": "sarah@email.com",
  "phone": "+234-123-456-7890",
  "date": "2024-02-15",
  "time": "2:00 PM",
  "service": "Bridal Makeup",
  "message": "Wedding on March 1st, need trial session"
}
```

#### Response
```json
{
  "success": true,
  "appointment": {
    "id": 123,
    "name": "Sarah Johnson",
    "email": "sarah@email.com",
    "phone": "+234-123-456-7890",
    "appointment_date": "2024-02-15",
    "appointment_time": "2:00 PM",
    "service": "Bridal Makeup",
    "message": "Wedding on March 1st, need trial session",
    "status": "pending",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "message": "Appointment booked successfully"
}
```

---

## 🎓 Class Enrollment

### Enroll in a Class
**POST** `/api/classes`

Creates a new makeup class enrollment.

#### Request Body
```json
{
  "name": "string (required)",
  "email": "string (required, valid email)",
  "phone": "string (required)",
  "classType": "string (required, see valid types below)",
  "experience": "string (required, see valid levels below)",
  "goals": "string (optional)",
  "schedule": "string (required, see valid schedules below)"
}
```

#### Valid Class Types
- `Beginner Basics` (₦35,000)
- `Advanced Techniques` (₦75,000)
- `Bridal Specialist` (₦60,000)
- `Business Training` (₦45,000)

#### Valid Experience Levels
- `Complete Beginner`
- `Some Experience`
- `Intermediate`
- `Advanced`

#### Valid Schedules
- `Weekdays`
- `Weekends`
- `Evening Classes`
- `Flexible`

#### Example Request
```json
{
  "name": "Maria Garcia",
  "email": "maria@email.com",
  "phone": "+234-987-654-3210",
  "classType": "Beginner Basics",
  "experience": "Complete Beginner",
  "goals": "Want to learn basic makeup for personal use and maybe start a side business",
  "schedule": "Weekends"
}
```

#### Response
```json
{
  "success": true,
  "enrollment": {
    "id": 456,
    "name": "Maria Garcia",
    "email": "maria@email.com",
    "phone": "+234-987-654-3210",
    "class_type": "Beginner Basics",
    "experience_level": "Complete Beginner",
    "goals": "Want to learn basic makeup for personal use and maybe start a side business",
    "preferred_schedule": "Weekends",
    "status": "pending",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "message": "Class enrollment successful"
}
```

---

## 👨‍💼 Admin Management

### Admin Signup
**POST** `/api/admin/signup`

Creates a new admin account.

#### Request Body
```json
{
  "email": "string (required, valid email)",
  "password": "string (required, min 6 characters)",
  "name": "string (required)"
}
```

#### Example Request
```json
{
  "email": "admin@mumuni.com",
  "password": "admin123456",
  "name": "Admin User"
}
```

#### Response
```json
{
  "success": true,
  "message": "Admin account created successfully"
}
```

### Admin Login
**POST** `/api/admin/login`

Authenticates an admin user and returns a JWT token.

#### Request Body
```json
{
  "email": "string (required, valid email)",
  "password": "string (required)"
}
```

#### Example Request
```json
{
  "email": "admin@mumuni.com",
  "password": "admin123456"
}
```

#### Response
```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "admin": {
    "id": "uuid",
    "email": "admin@mumuni.com",
    "name": "Admin User",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "message": "Login successful"
}
```

### Get All Appointments (Admin Only)
**GET** `/api/admin/appointments`

Retrieves all appointment bookings. Requires admin authentication.

#### Headers
```
Authorization: Bearer <jwt_token>
```

#### Response
```json
{
  "success": true,
  "appointments": [
    {
      "id": 123,
      "name": "Sarah Johnson",
      "email": "sarah@email.com",
      "phone": "+234-123-456-7890",
      "appointment_date": "2024-02-15",
      "appointment_time": "2:00 PM",
      "service": "Bridal Makeup",
      "message": "Wedding on March 1st, need trial session",
      "status": "pending",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "count": 1
}
```

### Get All Class Enrollments (Admin Only)
**GET** `/api/admin/classes`

Retrieves all class enrollments. Requires admin authentication.

#### Headers
```
Authorization: Bearer <jwt_token>
```

#### Response
```json
{
  "success": true,
  "classes": [
    {
      "id": 456,
      "name": "Maria Garcia",
      "email": "maria@email.com",
      "phone": "+234-987-654-3210",
      "class_type": "Beginner Basics",
      "experience_level": "Complete Beginner",
      "goals": "Want to learn basic makeup for personal use",
      "preferred_schedule": "Weekends",
      "status": "pending",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "count": 1
}
```

---

## 🔧 Utility Endpoints

### Health Check
**GET** `/health`

Returns the health status of the service.

#### Response
```json
{
  "status": "healthy",
  "service": "mumuni_backend"
}
```

---

## ⚠️ Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request data: field validation error"
}
```

### 401 Unauthorized
```json
{
  "error": "Invalid email or password"
}
```

### 403 Forbidden
```json
{
  "error": "Invalid or expired token"
}
```

### 409 Conflict
```json
{
  "error": "Admin with this email already exists"
}
```

### 500 Internal Server Error
```json
{
  "error": "Failed to book appointment"
}
```

---

## 📝 Usage Examples

### JavaScript/TypeScript Examples

#### Book Appointment
```javascript
const bookAppointment = async (appointmentData) => {
  try {
    const response = await fetch('/api/appointments', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(appointmentData),
    });

    if (!response.ok) {
      throw new Error('Failed to book appointment');
    }

    const result = await response.json();
    console.log('Appointment booked:', result);
    return result;
  } catch (error) {
    console.error('Error booking appointment:', error);
    throw error;
  }
};
```

#### Enroll in Class
```javascript
const enrollInClass = async (enrollmentData) => {
  try {
    const response = await fetch('/api/classes', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(enrollmentData),
    });

    if (!response.ok) {
      throw new Error('Failed to enroll in class');
    }

    const result = await response.json();
    console.log('Enrollment successful:', result);
    return result;
  } catch (error) {
    console.error('Error enrolling in class:', error);
    throw error;
  }
};
```

#### Admin Login
```javascript
const adminLogin = async (credentials) => {
  try {
    const response = await fetch('/api/admin/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      throw new Error('Login failed');
    }

    const result = await response.json();
    localStorage.setItem('adminToken', result.token);
    return result;
  } catch (error) {
    console.error('Error logging in:', error);
    throw error;
  }
};
```

#### Get Appointments (Admin)
```javascript
const getAppointments = async () => {
  try {
    const token = localStorage.getItem('adminToken');
    const response = await fetch('/api/admin/appointments', {
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch appointments');
    }

    const result = await response.json();
    return result.appointments;
  } catch (error) {
    console.error('Error fetching appointments:', error);
    throw error;
  }
};
```

---

## 🚀 Getting Started

1. **Install Dependencies**
   ```bash
   go mod tidy
   ```

2. **Set up Environment Variables**
   Create a `.env` file with your Supabase credentials:
   ```
   SUPABASE_URL=your_supabase_url
   SUPABASE_ANON_KEY=your_supabase_anon_key
   SUPABASE_SERVICE_ROLE_KEY=your_supabase_service_role_key
   JWT_SECRET=your_jwt_secret
   PORT=8080
   ```

3. **Set up Database**
   Run the SQL schema in your Supabase database:
   ```bash
   # Execute the contents of database/schema.sql in your Supabase SQL editor
   ```

4. **Run the Server**
   ```bash
   go run main.go
   ```

5. **Test the API**
   ```bash
   go run test_api.go
   ```

---

## 📊 Database Schema

The application uses three main tables:

- **admin_users**: Stores admin account information
- **appointments**: Stores makeup appointment bookings
- **classes**: Stores makeup class enrollments

See `database/schema.sql` for the complete schema definition.

---

## 🔒 Security Features

- JWT-based authentication for admin endpoints
- Password hashing using bcrypt
- CORS middleware for cross-origin requests
- Input validation and sanitization
- SQL injection protection through Supabase client

---

## 🛠️ Development

- **Format Code**: `make fmt`
- **Run Tests**: `make test`
- **Build**: `make build`
- **Run**: `make run`
- **Clean**: `make clean`
