

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS user_profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),                              -- Unique profile ID
    user_id INT REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' table
    bio TEXT,                                            -- User's biography or description
    profile_picture_url TEXT,                           -- URL to the user's profile picture (stored in S3)
    work_experience JSONB,                              -- JSONB field to store work experience details
    created_at TIMESTAMPTZ DEFAULT NOW(),               -- Record creation timestamp
    updated_at TIMESTAMPTZ DEFAULT NOW()                -- Record update timestamp
);
