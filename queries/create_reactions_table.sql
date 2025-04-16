CREATE TABLE IF NOT EXISTS reactions ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique reaction ID 
  post_id UUID REFERENCES posts(id) ON DELETE CASCADE,   -- Foreign key to 'posts' (post being reacted to) 
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,   -- Foreign key to 'users' (user who reacted) 
  reaction_type VARCHAR(50) NOT NULL,                     -- Reaction type (e.g., 'laugh') 
  created_at TIMESTAMPTZ DEFAULT NOW(),                   -- Reaction creation timestamp 
  UNIQUE (post_id, user_id, reaction_type)                -- Ensure unique reactions of a certain type per post by a user
);
