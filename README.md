# Mumuni Backend

A Go backend service for appointment booking and class enrollment with Supabase integration.

## Features

- Appointment booking system
- Class enrollment system
- Admin authentication (login/signup)
- Admin dashboard for viewing appointments and classes
- Supabase database integration

## Quick Setup

### 1. Configure Supabase

1. Create a Supabase project at [supabase.com](https://supabase.com)
2. Go to **Settings** → **API** in your Supabase dashboard
3. Copy your project URL and API keys

### 2. Set Up Environment Variables

Create a `.env` file with your Supabase credentials:

```env
# Supabase Configuration (get from Settings > API)
SUPABASE_URL=https://your-project-ref.supabase.co
SUPABASE_ANON_KEY=your_supabase_anon_key
SUPABASE_SERVICE_ROLE_KEY=your_supabase_service_role_key

# JWT Secret for admin authentication
JWT_SECRET=your_jwt_secret_key_here

# Server Configuration
PORT=8080
```

**Important**: Use the REST API URL (https://), NOT the database URL (postgresql://)

### 3. Set Up Database

1. Go to your Supabase project dashboard
2. Navigate to **SQL Editor**
3. Copy and paste the contents of `database/schema.sql`
4. Click **Run** to create the tables

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Server

```bash
go run main.go
```

You should see:
```
Initializing Supabase client with URL: https://your-project-ref.supabase.co
Anon key length: XX characters
Server starting on port 8080
```

### 6. Test the API

```bash
go run test_api.go
```

## API Endpoints

### Public Endpoints
- `POST /api/appointments` - Book an appointment
- `POST /api/classes` - Enroll in a class
- `GET /health` - Health check

### Admin Endpoints
- `POST /api/admin/signup` - Admin signup
- `POST /api/admin/login` - Admin login
- `GET /api/admin/appointments` - Get all appointments (requires auth)
- `GET /api/admin/classes` - Get all class enrollments (requires auth)

## Troubleshooting

### Common Issues

1. **"unsupported protocol scheme 'postgresql'"**
   - Solution: Use the REST API URL (https://) instead of database URL (postgresql://)

2. **"SUPABASE_URL is required"**
   - Solution: Make sure your `.env` file exists and contains the correct URL

3. **Database connection errors**
   - Solution: Ensure you've run the SQL schema in Supabase SQL Editor

### Need Help?

Check the `SETUP_GUIDE.md` file for detailed troubleshooting steps.
