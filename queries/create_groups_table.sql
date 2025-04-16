CREATE TABLE IF NOT EXISTS groups ( 
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),         -- Unique group ID 
  name VARCHAR(100) NOT NULL UNIQUE,                       -- Group name (e.g., 'Sports Club') 
  description TEXT,                                       -- Description of the group 
  created_at TIMESTAMPTZ DEFAULT NOW()                     -- Timestamp when the group was created 
);
