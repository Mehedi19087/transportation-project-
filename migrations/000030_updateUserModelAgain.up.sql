ALTER TABLE users ADD COLUMN IF NOT EXISTS status TEXT NOT NULL DEFAULT 'pending';
ALTER TABLE users ADD CONSTRAINT check_status_valid CHECK (status IN ('pending', 'active', 'blocked'));