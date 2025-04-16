CREATE TABLE IF NOT EXISTS conversations ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique conversation ID 
  user1_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to 'users' (first user in the conversation) 
  user2_id UUID REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key to 'users' (second user in the conversation) 
  created_at TIMESTAMPTZ DEFAULT NOW(),                  -- Timestamp when the conversation was created 
  UNIQUE (user1_id, user2_id)                            -- Ensure a unique conversation between the two users
);
