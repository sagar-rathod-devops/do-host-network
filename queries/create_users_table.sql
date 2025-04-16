-- Enable the uuid-ossp extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   username VARCHAR(50) UNIQUE NOT NULL, 
    email VARCHAR(100) UNIQUE NOT NULL, 
    mobile_number VARCHAR(15) UNIQUE NOT NULL, 
    password_hash TEXT NOT NULL, 
    first_name VARCHAR(100), 
    last_name VARCHAR(100), 
    is_verified BOOLEAN DEFAULT FALSE, 
    mobile_verified BOOLEAN DEFAULT FALSE, 
    created_at TIMESTAMPTZ DEFAULT NOW(), 
    updated_at TIMESTAMPTZ DEFAULT NOW(), 
    last_login TIMESTAMPTZ, 
    reset_token TEXT, 
    reset_token_expires TIMESTAMPTZ, 
    jwt_token TEXT, 
    jwt_token_expires TIMESTAMPTZ
);
