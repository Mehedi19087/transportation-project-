 -- Remove the role constraint
   ALTER TABLE users DROP CONSTRAINT IF EXISTS check_role_valid;

   -- Drop the new columns and restore google_id
   ALTER TABLE users 
       DROP COLUMN IF EXISTS password,
       DROP COLUMN IF EXISTS role,
       ADD COLUMN IF NOT EXISTS google_id TEXT;

   -- Recreate the original index
   CREATE INDEX IF NOT EXISTS ux_users_google_id ON users (google_id);