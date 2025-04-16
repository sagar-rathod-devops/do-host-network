CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),        -- Unique post ID
    user_id INT REFERENCES users(id) ON DELETE CASCADE,     -- Foreign key to 'users' table (author of the post)
    content TEXT,                                           -- Text content of the post
    media_urls TEXT[],                                      -- Array of URLs for media (images/videos) stored in AWS S3
    created_at TIMESTAMPTZ DEFAULT NOW(),                   -- Post creation timestamp
    updated_at TIMESTAMPTZ DEFAULT NOW(),                   -- Last updated timestamp
    is_deleted BOOLEAN DEFAULT FALSE                        -- Soft delete flag for posts
);
