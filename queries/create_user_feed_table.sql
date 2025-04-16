CREATE TABLE IF NOT EXISTS user_feed ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),      -- Unique feed entry ID 
  user_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' 
  post_id UUID REFERENCES posts(id) ON DELETE CASCADE, -- Foreign key to 'posts' 
  created_at TIMESTAMPTZ DEFAULT NOW()                 -- Entry was created 
);
