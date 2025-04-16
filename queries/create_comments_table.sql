CREATE TABLE IF NOT EXISTS comments ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique comment ID 
  post_id UUID REFERENCES posts(id) ON DELETE CASCADE,   -- Foreign key to 'posts' (post being commented on) 
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,   -- Foreign key to 'users' (user who commented) 
  content TEXT NOT NULL,                                  -- Comment text/content 
  created_at TIMESTAMPTZ DEFAULT NOW(),                   -- Comment creation timestamp 
  updated_at TIMESTAMPTZ DEFAULT NOW()                    -- Comment update timestamp 
);
