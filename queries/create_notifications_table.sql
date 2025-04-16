CREATE TABLE IF NOT EXISTS notifications ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique notification ID 
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,   -- Foreign key to 'users' (the user receiving the notification) 
  sender_id UUID REFERENCES users(id) ON DELETE CASCADE, -- Foreign key to 'users' (the user who triggered the notification) 
  notification_type VARCHAR(50) NOT NULL,                 -- Type of notification (e.g., 'like', 'comment', 'friend_request') 
  post_id UUID REFERENCES posts(id) ON DELETE CASCADE,   -- Foreign key to 'posts' (if related to a post) 
  message TEXT,                                          -- Notification message (e.g., "User A liked your post") 
  is_read BOOLEAN DEFAULT FALSE,                          -- Read status (true if the notification is read) 
  created_at TIMESTAMPTZ DEFAULT NOW()                    -- Notification creation timestamp 
);
