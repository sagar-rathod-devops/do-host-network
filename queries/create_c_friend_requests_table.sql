CREATE TABLE IF NOT EXISTS friend_requests ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique friend request ID 
  sender_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user who sends the friend request) 
  receiver_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user receiving the friend request) 
  status VARCHAR(20) DEFAULT 'pending',                  -- Status of the friend request (e.g., 'pending', 'accepted', 'declined') 
  created_at TIMESTAMPTZ DEFAULT NOW(),                  -- Friend request creation timestamp 
  updated_at TIMESTAMPTZ DEFAULT NOW(),                  -- Timestamp for when the status was last updated 
  UNIQUE (sender_id, receiver_id)                        -- Ensure a user can send a request to another user only once
);
