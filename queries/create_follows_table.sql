CREATE TABLE IF NOT EXISTS follows ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique follow relationship ID 
  follower_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user who follows) 
  followed_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user being followed) 
  created_at TIMESTAMPTZ DEFAULT NOW(),                  -- Relationship creation timestamp 
  UNIQUE (follower_id, followed_id)                       -- Ensure a user can follow another user only once
);
