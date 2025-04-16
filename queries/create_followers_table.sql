CREATE TABLE IF NOT EXISTS followers ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique follower relationship ID 
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,   -- Foreign key to 'users' (the user being followed) 
  follower_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (the follower) 
  created_at TIMESTAMPTZ DEFAULT NOW()                    -- Relationship was created 
);
