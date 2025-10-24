# 🚀 Quick Setup Guide

## Supabase Configuration Issue Fix

The error you're seeing indicates that the Supabase URL is incorrect. Here's how to fix it:

### 1. Get Your Supabase Credentials

1. Go to your Supabase project dashboard
2. Navigate to **Settings** → **API**
3. Copy the following values:

### 2. Correct URL Format

**❌ Wrong (Database URL):**
```
postgresql://postgres:***@db.jikbpfanqzhbhuglbtqo.supabase.co:5432/postgres
```

**✅ Correct (REST API URL):**
```
https://jikbpfanqzhbhuglbtqo.supabase.co
```

### 3. Update Your .env File

Create a `.env` file in your project root with:

```env
# Supabase Configuration
SUPABASE_URL=https://jikbpfanqzhbhuglbtqo.supabase.co
SUPABASE_ANON_KEY=your_anon_key_from_supabase_dashboard
SUPABASE_SERVICE_ROLE_KEY=your_service_role_key_from_supabase_dashboard

# JWT Secret for admin authentication
JWT_SECRET=your_jwt_secret_key_here

# Server Configuration
PORT=8080
```

### 4. Set Up Database Schema

1. Go to your Supabase project dashboard
2. Navigate to **SQL Editor**
3. Copy and paste the contents of `database/schema.sql`
4. Click **Run** to create the tables

### 5. Test the Setup

1. Run the server:
   ```bash
   go run main.go
   ```

2. You should see:
   ```
   Initializing Supabase client with URL: https://jikbpfanqzhbhuglbtqo.supabase.co
   Anon key length: XX characters
   Server starting on port 8080
   ```

3. Test the API:
   ```bash
   curl -X POST http://localhost:8080/api/admin/signup \
     -H "Content-Type: application/json" \
     -d '{"email":"admin@test.com","password":"password123","name":"Test Admin"}'
   ```

### Common Issues

- **Wrong URL format**: Make sure you're using the REST API URL (https://), not the database URL (postgresql://)
- **Missing keys**: Ensure both SUPABASE_ANON_KEY and SUPABASE_SERVICE_ROLE_KEY are set
- **Database not set up**: Run the SQL schema in Supabase SQL Editor first

### Need Help?

If you're still having issues, check:
1. Your Supabase project is active
2. The API keys are correct
3. The database schema has been created
4. Your .env file is in the project root directory
