# Project Overview and Assistant Context

Project Type: Multi-tenant SaaS Application
Backend: FastAPI + Supabase
Authentication: Supabase Auth
Database: PostgreSQL (via Supabase)

## Technical Stack Details

- FastAPI for REST API
- Supabase for:
  - Authentication and Authorization
  - Database (PostgreSQL)
  - Row Level Security (RLS)
- Pydantic for data validation
- Python 3.12+ compatible
- Environment-based configuration

## Current Implementation Status

- Authentication system implemented with email/password
- Multi-tenant architecture with user-tenant relationships
- Email verification flow
- Token-based authentication
- Admin operations using Supabase service role

## Core Features

1. User Management:
   - Signup
   - Login/Logout
   - Email verification
   - Token-based authentication

2. Tenant Management:
   - Multi-tenant support
   - Tenant creation
   - User-tenant relationships

3. Security:
   - Row Level Security (RLS)
   - Role-based access control
   - Service role for admin operations
