CREATE TABLE IF NOT EXISTS messages ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique message ID 
  sender_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user sending the message) 
  receiver_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (user receiving the message) 
  conversation_id UUID NOT NULL,                         -- Identifier for the conversation (UUID)
  content TEXT NOT NULL,                                  -- The text content of the message 
  created_at TIMESTAMPTZ DEFAULT NOW(),                   -- Timestamp when the message was created 
  is_read BOOLEAN DEFAULT FALSE,                          -- Read status (true if the message has been read) 
  UNIQUE (sender_id, receiver_id, created_at)             -- Ensure unique messages in a conversation based on sender, receiver, and timestamp
);
