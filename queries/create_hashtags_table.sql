CREATE TABLE IF NOT EXISTS hashtags ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),         -- Unique hashtag ID 
  name VARCHAR(50) NOT NULL UNIQUE,                        -- Hashtag name (e.g., '#example') 
  created_at TIMESTAMPTZ DEFAULT NOW()                     -- Timestamp when the hashtag was created 
);
